package app

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/caarlos0/env"
	"github.com/linkpoolio/bridges"
	"github.com/sirupsen/logrus"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type GnosisSafeTx struct {
	To             common.Address
	Value          *big.Int
	Data           []byte
	Operation      uint8
	SafeTxGas      *big.Int
	BaseGas        *big.Int
	GasPrice       *big.Int
	GasToken       common.Address
	RefundReceiver common.Address
	Nonce          *big.Int
}

// RelayerAdapter is a Chainlink external Adapter.
// It takes the same parameters as the EthTx core adapter
// But wraps everything in a meta transaction for a Gnosis Safe
type RelayerAdapter struct {
	client      *ethclient.Client
	gnosisOwner *ecdsa.PrivateKey
	chainID     *big.Int
	gnosis      common.Address

	relayerPrivateKey *ecdsa.PrivateKey
	relayer           common.Address
}

func NewRelayerAdapterFromEnv() (*RelayerAdapter, error) {
	type config struct {
		RPCURL                string `env:"RPC_URL,required"`
		GnosisOwnerPrivateKey string `env:"GNOSIS_OWNER_PRIVATE_KEY,required"`
		Gnosis                string `env:"GNOSIS_ADDRESS,required"`

		RelayerPrivateKey string `env:"RELAYER_PRIVATE_KEY,required"`
	}

	c := &config{}
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("could not parse environment config: %w", err)
	}
	client, err := ethclient.Dial(c.RPCURL)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	gnosisOwner, err := crypto.HexToECDSA(c.GnosisOwnerPrivateKey)
	if err != nil {
		return nil, err
	}

	relayerPrivateKey, err := crypto.HexToECDSA(c.RelayerPrivateKey)
	if err != nil {
		return nil, err
	}
	relayer := crypto.PubkeyToAddress(relayerPrivateKey.Public().(ecdsa.PublicKey))

	if !common.IsHexAddress(c.Gnosis) {
		return nil, fmt.Errorf("invalid gnosis address")
	}

	return &RelayerAdapter{
		client:            client,
		gnosisOwner:       gnosisOwner,
		chainID:           chainID,
		gnosis:            common.HexToAddress(c.Gnosis),
		relayerPrivateKey: relayerPrivateKey,
		relayer:           relayer,
	}, nil
}

func (adapter *RelayerAdapter) Run(h *bridges.Helper) (interface{}, error) {
	parsed, err := abi.JSON(strings.NewReader(GnosisSafeMinimalABI))
	if err != nil {
		logrus.WithError(err).Error("could not parse Safe ABI")
		return nil, err
	}

	nonceData, err := parsed.Pack("nonce")
	if err != nil {
		logrus.WithError(err).Error("could not pack nonce call")
		return nil, err
	}

	safeNonce, err := adapter.client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &adapter.gnosis,
		Data: nonceData,
	}, nil)
	if err != nil {
		logrus.WithError(err).Error("could not retrieve safe nonce")
		return nil, err
	}

	functionSelector := "0x"
	if value := h.GetParam("functionSelector"); len(value) != 0 {
		functionSelector = value
	}

	dataPrefix := ""
	if value := h.GetParam("dataPrefix"); len(value) != 0 {
		dataPrefix = value[2:]
	}

	result := ""
	if value := h.GetParam("result"); len(value) != 0 {
		result = value[2:]
	}

	address := "0x0000000000000000000000000000000000000000"
	if value := h.GetParam("address"); len(value) != 0 {
		address = value
	}

	safeData := functionSelector + dataPrefix + result

	safeTx := &GnosisSafeTx{
		To:        common.HexToAddress(address),
		Value:     big.NewInt(0),
		Data:      common.FromHex(safeData),
		SafeTxGas: big.NewInt(0),
		BaseGas:   big.NewInt(0),
		GasPrice:  big.NewInt(0),
		Nonce:     new(big.Int).SetBytes(safeNonce),
	}

	data, err := adapter.makeGnosisData(safeTx)
	if err != nil {
		logrus.WithError(err).Error("could not make safe transaction data")
		return nil, err
	}

	txHash, err := adapter.relay(data)
	if err != nil {
		logrus.WithError(err).Error("could not relay transaction")
		return nil, err
	}

	res := make(map[string]interface{})
	res["tx_hash"] = txHash
	return res, nil
}

func (adapter *RelayerAdapter) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "relayer",
		Lambda: true,
	}
}

func (adapter *RelayerAdapter) makeGnosisData(safeTx *GnosisSafeTx) ([]byte, error) {
	gnosisABI, _ := abi.JSON(strings.NewReader(GnosisSafeMinimalABI))
	hashData, err := gnosisABI.Pack(
		"getTransactionHash",
		safeTx.To,
		safeTx.Value,
		safeTx.Data,
		safeTx.Operation,
		safeTx.SafeTxGas,
		safeTx.BaseGas,
		safeTx.GasPrice,
		safeTx.GasToken,
		safeTx.RefundReceiver,
		safeTx.Nonce,
	)
	if err != nil {
		return nil, fmt.Errorf("gnosis ABI getTransactionHash Pack failed: %w", err)
	}

	message, err := adapter.client.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   &adapter.gnosis,
		Data: hashData,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("gnosis Call Contract failed: %w", err)
	}

	signature, err := crypto.Sign(message, adapter.gnosisOwner)
	if err != nil {
		return nil, fmt.Errorf("signature failed: %w", err)
	}

	signature[64] = signature[64] + uint8(27)
	data, err := gnosisABI.Pack(
		"execTransaction",
		safeTx.To,
		safeTx.Value,
		safeTx.Data,
		safeTx.Operation,
		safeTx.SafeTxGas,
		safeTx.BaseGas,
		safeTx.GasPrice,
		safeTx.GasToken,
		safeTx.RefundReceiver,
		signature,
	)
	if err != nil {
		return nil, fmt.Errorf("gnosis ABI execTransaction failed: %w", err)
	}

	return data, nil
}

func (adapter *RelayerAdapter) relay(data []byte) (string, error) {
	nonce, err := adapter.client.NonceAt(context.Background(), adapter.relayer, nil)
	if err != nil {
		return "", fmt.Errorf("could not get relayer nonce: %w", err)
	}

	tx := types.NewTransaction(nonce, adapter.gnosis, big.NewInt(0), 800000, big.NewInt(1000000000), data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(adapter.chainID), adapter.relayerPrivateKey)
	if err != nil {
		return "", fmt.Errorf("could not sign transaction: %w", err)
	}

	if err := adapter.client.SendTransaction(context.Background(), signedTx); err != nil {
		return "", fmt.Errorf("could not send transaction: %w", err)
	}

	return signedTx.Hash().Hex(), nil
}
