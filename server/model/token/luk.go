// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LukTokenMetaData contains all meta data concerning the LukToken contract.
var LukTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"synthesize\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"orePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_whiteList\",\"type\":\"address\"}],\"name\":\"getWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"officeReceiveAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"possessor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"officeAddress\",\"type\":\"address\"}],\"name\":\"setOfficeReceiveAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isPool\",\"type\":\"bool\"}],\"name\":\"setPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_possessor\",\"type\":\"address\"}],\"name\":\"setPossessor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_whiteList\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWhite\",\"type\":\"bool\"}],\"name\":\"setWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LukTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use LukTokenMetaData.ABI instead.
var LukTokenABI = LukTokenMetaData.ABI

// LukToken is an auto generated Go binding around an Ethereum contract.
type LukToken struct {
	LukTokenCaller     // Read-only binding to the contract
	LukTokenTransactor // Write-only binding to the contract
	LukTokenFilterer   // Log filterer for contract events
}

// LukTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type LukTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LukTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LukTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LukTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LukTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LukTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LukTokenSession struct {
	Contract     *LukToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LukTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LukTokenCallerSession struct {
	Contract *LukTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LukTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LukTokenTransactorSession struct {
	Contract     *LukTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LukTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type LukTokenRaw struct {
	Contract *LukToken // Generic contract binding to access the raw methods on
}

// LukTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LukTokenCallerRaw struct {
	Contract *LukTokenCaller // Generic read-only contract binding to access the raw methods on
}

// LukTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LukTokenTransactorRaw struct {
	Contract *LukTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLukToken creates a new instance of LukToken, bound to a specific deployed contract.
func NewLukToken(address common.Address, backend bind.ContractBackend) (*LukToken, error) {
	contract, err := bindLukToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LukToken{LukTokenCaller: LukTokenCaller{contract: contract}, LukTokenTransactor: LukTokenTransactor{contract: contract}, LukTokenFilterer: LukTokenFilterer{contract: contract}}, nil
}

// NewLukTokenCaller creates a new read-only instance of LukToken, bound to a specific deployed contract.
func NewLukTokenCaller(address common.Address, caller bind.ContractCaller) (*LukTokenCaller, error) {
	contract, err := bindLukToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LukTokenCaller{contract: contract}, nil
}

// NewLukTokenTransactor creates a new write-only instance of LukToken, bound to a specific deployed contract.
func NewLukTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LukTokenTransactor, error) {
	contract, err := bindLukToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LukTokenTransactor{contract: contract}, nil
}

// NewLukTokenFilterer creates a new log filterer instance of LukToken, bound to a specific deployed contract.
func NewLukTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*LukTokenFilterer, error) {
	contract, err := bindLukToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LukTokenFilterer{contract: contract}, nil
}

// bindLukToken binds a generic wrapper to an already deployed contract.
func bindLukToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LukTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LukToken *LukTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LukToken.Contract.LukTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LukToken *LukTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LukToken.Contract.LukTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LukToken *LukTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LukToken.Contract.LukTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LukToken *LukTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LukToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LukToken *LukTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LukToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LukToken *LukTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LukToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LukToken *LukTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LukToken *LukTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LukToken.Contract.Allowance(&_LukToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LukToken *LukTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LukToken.Contract.Allowance(&_LukToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LukToken *LukTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LukToken *LukTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LukToken.Contract.BalanceOf(&_LukToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LukToken *LukTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LukToken.Contract.BalanceOf(&_LukToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LukToken *LukTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LukToken *LukTokenSession) Decimals() (uint8, error) {
	return _LukToken.Contract.Decimals(&_LukToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LukToken *LukTokenCallerSession) Decimals() (uint8, error) {
	return _LukToken.Contract.Decimals(&_LukToken.CallOpts)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0x1228cbee.
//
// Solidity: function getPoolAddress(address pool) view returns(bool)
func (_LukToken *LukTokenCaller) GetPoolAddress(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "getPoolAddress", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPoolAddress is a free data retrieval call binding the contract method 0x1228cbee.
//
// Solidity: function getPoolAddress(address pool) view returns(bool)
func (_LukToken *LukTokenSession) GetPoolAddress(pool common.Address) (bool, error) {
	return _LukToken.Contract.GetPoolAddress(&_LukToken.CallOpts, pool)
}

// GetPoolAddress is a free data retrieval call binding the contract method 0x1228cbee.
//
// Solidity: function getPoolAddress(address pool) view returns(bool)
func (_LukToken *LukTokenCallerSession) GetPoolAddress(pool common.Address) (bool, error) {
	return _LukToken.Contract.GetPoolAddress(&_LukToken.CallOpts, pool)
}

// GetWhiteList is a free data retrieval call binding the contract method 0x9dfe9d68.
//
// Solidity: function getWhiteList(address _whiteList) view returns(bool)
func (_LukToken *LukTokenCaller) GetWhiteList(opts *bind.CallOpts, _whiteList common.Address) (bool, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "getWhiteList", _whiteList)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetWhiteList is a free data retrieval call binding the contract method 0x9dfe9d68.
//
// Solidity: function getWhiteList(address _whiteList) view returns(bool)
func (_LukToken *LukTokenSession) GetWhiteList(_whiteList common.Address) (bool, error) {
	return _LukToken.Contract.GetWhiteList(&_LukToken.CallOpts, _whiteList)
}

// GetWhiteList is a free data retrieval call binding the contract method 0x9dfe9d68.
//
// Solidity: function getWhiteList(address _whiteList) view returns(bool)
func (_LukToken *LukTokenCallerSession) GetWhiteList(_whiteList common.Address) (bool, error) {
	return _LukToken.Contract.GetWhiteList(&_LukToken.CallOpts, _whiteList)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LukToken *LukTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LukToken *LukTokenSession) Name() (string, error) {
	return _LukToken.Contract.Name(&_LukToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LukToken *LukTokenCallerSession) Name() (string, error) {
	return _LukToken.Contract.Name(&_LukToken.CallOpts)
}

// OfficeReceiveAddress is a free data retrieval call binding the contract method 0x0a4a78d3.
//
// Solidity: function officeReceiveAddress() view returns(address)
func (_LukToken *LukTokenCaller) OfficeReceiveAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "officeReceiveAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OfficeReceiveAddress is a free data retrieval call binding the contract method 0x0a4a78d3.
//
// Solidity: function officeReceiveAddress() view returns(address)
func (_LukToken *LukTokenSession) OfficeReceiveAddress() (common.Address, error) {
	return _LukToken.Contract.OfficeReceiveAddress(&_LukToken.CallOpts)
}

// OfficeReceiveAddress is a free data retrieval call binding the contract method 0x0a4a78d3.
//
// Solidity: function officeReceiveAddress() view returns(address)
func (_LukToken *LukTokenCallerSession) OfficeReceiveAddress() (common.Address, error) {
	return _LukToken.Contract.OfficeReceiveAddress(&_LukToken.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_LukToken *LukTokenCaller) Pair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_LukToken *LukTokenSession) Pair() (common.Address, error) {
	return _LukToken.Contract.Pair(&_LukToken.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0xa8aa1b31.
//
// Solidity: function pair() view returns(address)
func (_LukToken *LukTokenCallerSession) Pair() (common.Address, error) {
	return _LukToken.Contract.Pair(&_LukToken.CallOpts)
}

// Possessor is a free data retrieval call binding the contract method 0xdc358d8e.
//
// Solidity: function possessor() view returns(address)
func (_LukToken *LukTokenCaller) Possessor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "possessor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Possessor is a free data retrieval call binding the contract method 0xdc358d8e.
//
// Solidity: function possessor() view returns(address)
func (_LukToken *LukTokenSession) Possessor() (common.Address, error) {
	return _LukToken.Contract.Possessor(&_LukToken.CallOpts)
}

// Possessor is a free data retrieval call binding the contract method 0xdc358d8e.
//
// Solidity: function possessor() view returns(address)
func (_LukToken *LukTokenCallerSession) Possessor() (common.Address, error) {
	return _LukToken.Contract.Possessor(&_LukToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LukToken *LukTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LukToken *LukTokenSession) Symbol() (string, error) {
	return _LukToken.Contract.Symbol(&_LukToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LukToken *LukTokenCallerSession) Symbol() (string, error) {
	return _LukToken.Contract.Symbol(&_LukToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LukToken *LukTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LukToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LukToken *LukTokenSession) TotalSupply() (*big.Int, error) {
	return _LukToken.Contract.TotalSupply(&_LukToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LukToken *LukTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _LukToken.Contract.TotalSupply(&_LukToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LukToken *LukTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LukToken *LukTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.Approve(&_LukToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LukToken *LukTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.Approve(&_LukToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_LukToken *LukTokenTransactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_LukToken *LukTokenSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.Burn(&_LukToken.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_LukToken *LukTokenTransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.Burn(&_LukToken.TransactOpts, from, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LukToken *LukTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LukToken *LukTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.DecreaseAllowance(&_LukToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LukToken *LukTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.DecreaseAllowance(&_LukToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LukToken *LukTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LukToken *LukTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.IncreaseAllowance(&_LukToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LukToken *LukTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.IncreaseAllowance(&_LukToken.TransactOpts, spender, addedValue)
}

// SetOfficeReceiveAddress is a paid mutator transaction binding the contract method 0x06f2e5ab.
//
// Solidity: function setOfficeReceiveAddress(address officeAddress) returns()
func (_LukToken *LukTokenTransactor) SetOfficeReceiveAddress(opts *bind.TransactOpts, officeAddress common.Address) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "setOfficeReceiveAddress", officeAddress)
}

// SetOfficeReceiveAddress is a paid mutator transaction binding the contract method 0x06f2e5ab.
//
// Solidity: function setOfficeReceiveAddress(address officeAddress) returns()
func (_LukToken *LukTokenSession) SetOfficeReceiveAddress(officeAddress common.Address) (*types.Transaction, error) {
	return _LukToken.Contract.SetOfficeReceiveAddress(&_LukToken.TransactOpts, officeAddress)
}

// SetOfficeReceiveAddress is a paid mutator transaction binding the contract method 0x06f2e5ab.
//
// Solidity: function setOfficeReceiveAddress(address officeAddress) returns()
func (_LukToken *LukTokenTransactorSession) SetOfficeReceiveAddress(officeAddress common.Address) (*types.Transaction, error) {
	return _LukToken.Contract.SetOfficeReceiveAddress(&_LukToken.TransactOpts, officeAddress)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xcbfa11a2.
//
// Solidity: function setPoolAddress(address pool, bool isPool) returns()
func (_LukToken *LukTokenTransactor) SetPoolAddress(opts *bind.TransactOpts, pool common.Address, isPool bool) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "setPoolAddress", pool, isPool)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xcbfa11a2.
//
// Solidity: function setPoolAddress(address pool, bool isPool) returns()
func (_LukToken *LukTokenSession) SetPoolAddress(pool common.Address, isPool bool) (*types.Transaction, error) {
	return _LukToken.Contract.SetPoolAddress(&_LukToken.TransactOpts, pool, isPool)
}

// SetPoolAddress is a paid mutator transaction binding the contract method 0xcbfa11a2.
//
// Solidity: function setPoolAddress(address pool, bool isPool) returns()
func (_LukToken *LukTokenTransactorSession) SetPoolAddress(pool common.Address, isPool bool) (*types.Transaction, error) {
	return _LukToken.Contract.SetPoolAddress(&_LukToken.TransactOpts, pool, isPool)
}

// SetPossessor is a paid mutator transaction binding the contract method 0x0064645e.
//
// Solidity: function setPossessor(address _possessor) returns()
func (_LukToken *LukTokenTransactor) SetPossessor(opts *bind.TransactOpts, _possessor common.Address) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "setPossessor", _possessor)
}

// SetPossessor is a paid mutator transaction binding the contract method 0x0064645e.
//
// Solidity: function setPossessor(address _possessor) returns()
func (_LukToken *LukTokenSession) SetPossessor(_possessor common.Address) (*types.Transaction, error) {
	return _LukToken.Contract.SetPossessor(&_LukToken.TransactOpts, _possessor)
}

// SetPossessor is a paid mutator transaction binding the contract method 0x0064645e.
//
// Solidity: function setPossessor(address _possessor) returns()
func (_LukToken *LukTokenTransactorSession) SetPossessor(_possessor common.Address) (*types.Transaction, error) {
	return _LukToken.Contract.SetPossessor(&_LukToken.TransactOpts, _possessor)
}

// SetWhiteList is a paid mutator transaction binding the contract method 0x8d14e127.
//
// Solidity: function setWhiteList(address _whiteList, bool isWhite) returns()
func (_LukToken *LukTokenTransactor) SetWhiteList(opts *bind.TransactOpts, _whiteList common.Address, isWhite bool) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "setWhiteList", _whiteList, isWhite)
}

// SetWhiteList is a paid mutator transaction binding the contract method 0x8d14e127.
//
// Solidity: function setWhiteList(address _whiteList, bool isWhite) returns()
func (_LukToken *LukTokenSession) SetWhiteList(_whiteList common.Address, isWhite bool) (*types.Transaction, error) {
	return _LukToken.Contract.SetWhiteList(&_LukToken.TransactOpts, _whiteList, isWhite)
}

// SetWhiteList is a paid mutator transaction binding the contract method 0x8d14e127.
//
// Solidity: function setWhiteList(address _whiteList, bool isWhite) returns()
func (_LukToken *LukTokenTransactorSession) SetWhiteList(_whiteList common.Address, isWhite bool) (*types.Transaction, error) {
	return _LukToken.Contract.SetWhiteList(&_LukToken.TransactOpts, _whiteList, isWhite)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LukToken *LukTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LukToken *LukTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.Transfer(&_LukToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LukToken *LukTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.Transfer(&_LukToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LukToken *LukTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LukToken *LukTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.TransferFrom(&_LukToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LukToken *LukTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LukToken.Contract.TransferFrom(&_LukToken.TransactOpts, from, to, amount)
}

// LukTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LukToken contract.
type LukTokenApprovalIterator struct {
	Event *LukTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LukTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LukTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LukTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LukTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LukTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LukTokenApproval represents a Approval event raised by the LukToken contract.
type LukTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LukToken *LukTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LukTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LukToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LukTokenApprovalIterator{contract: _LukToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LukToken *LukTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LukTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LukToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LukTokenApproval)
				if err := _LukToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LukToken *LukTokenFilterer) ParseApproval(log types.Log) (*LukTokenApproval, error) {
	event := new(LukTokenApproval)
	if err := _LukToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LukTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LukToken contract.
type LukTokenTransferIterator struct {
	Event *LukTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LukTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LukTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LukTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LukTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LukTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LukTokenTransfer represents a Transfer event raised by the LukToken contract.
type LukTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LukToken *LukTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LukTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LukToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LukTokenTransferIterator{contract: _LukToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LukToken *LukTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LukTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LukToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LukTokenTransfer)
				if err := _LukToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LukToken *LukTokenFilterer) ParseTransfer(log types.Log) (*LukTokenTransfer, error) {
	event := new(LukTokenTransfer)
	if err := _LukToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
