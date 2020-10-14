// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package app

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GnosisSafeMinimalABI is the input ABI used to generate the binding from.
const GnosisSafeMinimalABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GnosisSafeMinimalBin is the compiled bytecode used for deploying new contracts.
var GnosisSafeMinimalBin = "0x608060405234801561001057600080fd5b506102ee806100206000396000f3fe6080604052600436106100345760003560e01c80636a76120214610039578063affed0e014610159578063d8d11f7814610180575b600080fd5b610145600480360361014081101561005057600080fd5b6001600160a01b038235169160208101359181019060608101604082013564010000000081111561008057600080fd5b82018360208201111561009257600080fd5b803590602001918460018302840111640100000000831117156100b457600080fd5b9193909260ff833516926020810135926040820135926060830135926001600160a01b03608082013581169360a083013590911692909160e081019060c0013564010000000081111561010657600080fd5b82018360208201111561011857600080fd5b8035906020019184600183028401116401000000008311171561013a57600080fd5b50909250905061027c565b604080519115158252519081900360200190f35b34801561016557600080fd5b5061016e61028e565b60408051918252519081900360200190f35b34801561018c57600080fd5b5061016e60048036036101408110156101a457600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156101d457600080fd5b8201836020820111156101e657600080fd5b8035906020019184600183028401116401000000008311171561020857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505060ff833516935050506020810135906040810135906060810135906001600160a01b03608082013581169160a08101359091169060c00135610294565b5060019b9a5050505050505050505050565b60005481565b60408051607360f81b815290519081900360010190209a995050505050505050505056fea26469706673582212200f948493aba9761a02536cee5429351753abc4e144c9e08620b8cd8c7935c13864736f6c634300060a0033"

// DeployGnosisSafeMinimal deploys a new Ethereum contract, binding an instance of GnosisSafeMinimal to it.
func DeployGnosisSafeMinimal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GnosisSafeMinimal, error) {
	parsed, err := abi.JSON(strings.NewReader(GnosisSafeMinimalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GnosisSafeMinimalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GnosisSafeMinimal{GnosisSafeMinimalCaller: GnosisSafeMinimalCaller{contract: contract}, GnosisSafeMinimalTransactor: GnosisSafeMinimalTransactor{contract: contract}, GnosisSafeMinimalFilterer: GnosisSafeMinimalFilterer{contract: contract}}, nil
}

// GnosisSafeMinimal is an auto generated Go binding around an Ethereum contract.
type GnosisSafeMinimal struct {
	GnosisSafeMinimalCaller     // Read-only binding to the contract
	GnosisSafeMinimalTransactor // Write-only binding to the contract
	GnosisSafeMinimalFilterer   // Log filterer for contract events
}

// GnosisSafeMinimalCaller is an auto generated read-only Go binding around an Ethereum contract.
type GnosisSafeMinimalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeMinimalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GnosisSafeMinimalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeMinimalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnosisSafeMinimalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeMinimalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnosisSafeMinimalSession struct {
	Contract     *GnosisSafeMinimal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GnosisSafeMinimalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnosisSafeMinimalCallerSession struct {
	Contract *GnosisSafeMinimalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GnosisSafeMinimalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnosisSafeMinimalTransactorSession struct {
	Contract     *GnosisSafeMinimalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GnosisSafeMinimalRaw is an auto generated low-level Go binding around an Ethereum contract.
type GnosisSafeMinimalRaw struct {
	Contract *GnosisSafeMinimal // Generic contract binding to access the raw methods on
}

// GnosisSafeMinimalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnosisSafeMinimalCallerRaw struct {
	Contract *GnosisSafeMinimalCaller // Generic read-only contract binding to access the raw methods on
}

// GnosisSafeMinimalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnosisSafeMinimalTransactorRaw struct {
	Contract *GnosisSafeMinimalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGnosisSafeMinimal creates a new instance of GnosisSafeMinimal, bound to a specific deployed contract.
func NewGnosisSafeMinimal(address common.Address, backend bind.ContractBackend) (*GnosisSafeMinimal, error) {
	contract, err := bindGnosisSafeMinimal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeMinimal{GnosisSafeMinimalCaller: GnosisSafeMinimalCaller{contract: contract}, GnosisSafeMinimalTransactor: GnosisSafeMinimalTransactor{contract: contract}, GnosisSafeMinimalFilterer: GnosisSafeMinimalFilterer{contract: contract}}, nil
}

// NewGnosisSafeMinimalCaller creates a new read-only instance of GnosisSafeMinimal, bound to a specific deployed contract.
func NewGnosisSafeMinimalCaller(address common.Address, caller bind.ContractCaller) (*GnosisSafeMinimalCaller, error) {
	contract, err := bindGnosisSafeMinimal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeMinimalCaller{contract: contract}, nil
}

// NewGnosisSafeMinimalTransactor creates a new write-only instance of GnosisSafeMinimal, bound to a specific deployed contract.
func NewGnosisSafeMinimalTransactor(address common.Address, transactor bind.ContractTransactor) (*GnosisSafeMinimalTransactor, error) {
	contract, err := bindGnosisSafeMinimal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeMinimalTransactor{contract: contract}, nil
}

// NewGnosisSafeMinimalFilterer creates a new log filterer instance of GnosisSafeMinimal, bound to a specific deployed contract.
func NewGnosisSafeMinimalFilterer(address common.Address, filterer bind.ContractFilterer) (*GnosisSafeMinimalFilterer, error) {
	contract, err := bindGnosisSafeMinimal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeMinimalFilterer{contract: contract}, nil
}

// bindGnosisSafeMinimal binds a generic wrapper to an already deployed contract.
func bindGnosisSafeMinimal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GnosisSafeMinimalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeMinimal *GnosisSafeMinimalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GnosisSafeMinimal.Contract.GnosisSafeMinimalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeMinimal *GnosisSafeMinimalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeMinimal.Contract.GnosisSafeMinimalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeMinimal *GnosisSafeMinimalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeMinimal.Contract.GnosisSafeMinimalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeMinimal *GnosisSafeMinimalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GnosisSafeMinimal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeMinimal *GnosisSafeMinimalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeMinimal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeMinimal *GnosisSafeMinimalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeMinimal.Contract.contract.Transact(opts, method, params...)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafeMinimal *GnosisSafeMinimalCaller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _GnosisSafeMinimal.contract.Call(opts, out, "getTransactionHash", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
	return *ret0, err
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafeMinimal *GnosisSafeMinimalSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _GnosisSafeMinimal.Contract.GetTransactionHash(&_GnosisSafeMinimal.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_GnosisSafeMinimal *GnosisSafeMinimalCallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _GnosisSafeMinimal.Contract.GetTransactionHash(&_GnosisSafeMinimal.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafeMinimal *GnosisSafeMinimalCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GnosisSafeMinimal.contract.Call(opts, out, "nonce")
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafeMinimal *GnosisSafeMinimalSession) Nonce() (*big.Int, error) {
	return _GnosisSafeMinimal.Contract.Nonce(&_GnosisSafeMinimal.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_GnosisSafeMinimal *GnosisSafeMinimalCallerSession) Nonce() (*big.Int, error) {
	return _GnosisSafeMinimal.Contract.Nonce(&_GnosisSafeMinimal.CallOpts)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_GnosisSafeMinimal *GnosisSafeMinimalTransactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafeMinimal.contract.Transact(opts, "execTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_GnosisSafeMinimal *GnosisSafeMinimalSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafeMinimal.Contract.ExecTransaction(&_GnosisSafeMinimal.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_GnosisSafeMinimal *GnosisSafeMinimalTransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _GnosisSafeMinimal.Contract.ExecTransaction(&_GnosisSafeMinimal.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}
