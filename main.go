package main

import (
	"errors"
	"fmt"

	"github.com/caarlos0/env"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/linkpoolio/bridges"
)

// Rockside is a Chainlink external Adapter.
// It takes the same parameters as the EthTx core adapter
// But wraps everything in a Rockside meta transaction
type Rockside struct {
	client *ethclient.Client
}

func NewRocksisdeFromEnv() (*Rockside, error) {
	type config struct {
		RpcURL string `env:"RPC_URL,required"`
	}

	c := &config{}
	if err := env.Parse(c); err != nil {
		return nil, fmt.Errorf("could not parse environment config: %w", err)
	}
	client, err := ethclient.Dial(c.RpcURL)

	// TODO:
	// - get chain ID
	// - parse private key from env (to sign meta tx)
	// - handle each request by computing the data, creating a meta tx (through a Safe
	//   given in parameters) and sending it

	return &Rockside{client}
}

type Request struct {
	FunctionSelector hexutil.Bytes `json:"functionSelector"`
	DataPrefix       hexutil.Bytes `json:"dataPrefix"`
	DataFormat       string        `json:"format"`
	GasLimit         uint64        `json:"gasLimit,omitempty"`
}

func fromHelper(h *bridges.Helper) (*Request, error) {
	fmt.Printf("%s\n", h.String())
	return nil, errors.New("not implemented")
}

// Run implements Bridge Run for querying the Wolfram short answers API
func (cc *Rockside) Run(h *bridges.Helper) (interface{}, error) {
	return nil, nil
}

// Opts is the bridge.Bridge implementation
func (cc *Rockside) Opts() *bridges.Opts {
	return &bridges.Opts{
		Name:   "WolframAlpha",
		Lambda: true,
	}
}

func main() {
	rockside := NewRocksideFromEnv()
	bridges.NewServer(rockside).Start(9090)
}
