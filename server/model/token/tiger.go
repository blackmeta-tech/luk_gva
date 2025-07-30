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

// TigerTokenMetaData contains all meta data concerning the TigerToken contract.
var TigerTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"billingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRandom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBillingAddress\",\"type\":\"address\"}],\"name\":\"setBillingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_random\",\"type\":\"string\"}],\"name\":\"setRandom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdtAddress\",\"type\":\"address\"}],\"name\":\"setUsdtAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"verify\",\"type\":\"bytes32\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TigerTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TigerTokenMetaData.ABI instead.
var TigerTokenABI = TigerTokenMetaData.ABI

// TigerToken is an auto generated Go binding around an Ethereum contract.
type TigerToken struct {
	TigerTokenCaller     // Read-only binding to the contract
	TigerTokenTransactor // Write-only binding to the contract
	TigerTokenFilterer   // Log filterer for contract events
}

// TigerTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TigerTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TigerTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TigerTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TigerTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TigerTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TigerTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TigerTokenSession struct {
	Contract     *TigerToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TigerTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TigerTokenCallerSession struct {
	Contract *TigerTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TigerTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TigerTokenTransactorSession struct {
	Contract     *TigerTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TigerTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TigerTokenRaw struct {
	Contract *TigerToken // Generic contract binding to access the raw methods on
}

// TigerTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TigerTokenCallerRaw struct {
	Contract *TigerTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TigerTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TigerTokenTransactorRaw struct {
	Contract *TigerTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTigerToken creates a new instance of TigerToken, bound to a specific deployed contract.
func NewTigerToken(address common.Address, backend bind.ContractBackend) (*TigerToken, error) {
	contract, err := bindTigerToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TigerToken{TigerTokenCaller: TigerTokenCaller{contract: contract}, TigerTokenTransactor: TigerTokenTransactor{contract: contract}, TigerTokenFilterer: TigerTokenFilterer{contract: contract}}, nil
}

// NewTigerTokenCaller creates a new read-only instance of TigerToken, bound to a specific deployed contract.
func NewTigerTokenCaller(address common.Address, caller bind.ContractCaller) (*TigerTokenCaller, error) {
	contract, err := bindTigerToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TigerTokenCaller{contract: contract}, nil
}

// NewTigerTokenTransactor creates a new write-only instance of TigerToken, bound to a specific deployed contract.
func NewTigerTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TigerTokenTransactor, error) {
	contract, err := bindTigerToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TigerTokenTransactor{contract: contract}, nil
}

// NewTigerTokenFilterer creates a new log filterer instance of TigerToken, bound to a specific deployed contract.
func NewTigerTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TigerTokenFilterer, error) {
	contract, err := bindTigerToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TigerTokenFilterer{contract: contract}, nil
}

// bindTigerToken binds a generic wrapper to an already deployed contract.
func bindTigerToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TigerTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TigerToken *TigerTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TigerToken.Contract.TigerTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TigerToken *TigerTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TigerToken.Contract.TigerTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TigerToken *TigerTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TigerToken.Contract.TigerTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TigerToken *TigerTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TigerToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TigerToken *TigerTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TigerToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TigerToken *TigerTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TigerToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TigerToken *TigerTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TigerToken *TigerTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TigerToken.Contract.BalanceOf(&_TigerToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TigerToken *TigerTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TigerToken.Contract.BalanceOf(&_TigerToken.CallOpts, owner)
}

// BillingAddress is a free data retrieval call binding the contract method 0x14e18a5a.
//
// Solidity: function billingAddress() view returns(address)
func (_TigerToken *TigerTokenCaller) BillingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "billingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BillingAddress is a free data retrieval call binding the contract method 0x14e18a5a.
//
// Solidity: function billingAddress() view returns(address)
func (_TigerToken *TigerTokenSession) BillingAddress() (common.Address, error) {
	return _TigerToken.Contract.BillingAddress(&_TigerToken.CallOpts)
}

// BillingAddress is a free data retrieval call binding the contract method 0x14e18a5a.
//
// Solidity: function billingAddress() view returns(address)
func (_TigerToken *TigerTokenCallerSession) BillingAddress() (common.Address, error) {
	return _TigerToken.Contract.BillingAddress(&_TigerToken.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TigerToken *TigerTokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TigerToken *TigerTokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TigerToken.Contract.GetApproved(&_TigerToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TigerToken *TigerTokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TigerToken.Contract.GetApproved(&_TigerToken.CallOpts, tokenId)
}

// GetRandom is a free data retrieval call binding the contract method 0xaacc5a17.
//
// Solidity: function getRandom() view returns(string)
func (_TigerToken *TigerTokenCaller) GetRandom(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "getRandom")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetRandom is a free data retrieval call binding the contract method 0xaacc5a17.
//
// Solidity: function getRandom() view returns(string)
func (_TigerToken *TigerTokenSession) GetRandom() (string, error) {
	return _TigerToken.Contract.GetRandom(&_TigerToken.CallOpts)
}

// GetRandom is a free data retrieval call binding the contract method 0xaacc5a17.
//
// Solidity: function getRandom() view returns(string)
func (_TigerToken *TigerTokenCallerSession) GetRandom() (string, error) {
	return _TigerToken.Contract.GetRandom(&_TigerToken.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TigerToken *TigerTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TigerToken *TigerTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TigerToken.Contract.IsApprovedForAll(&_TigerToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TigerToken *TigerTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TigerToken.Contract.IsApprovedForAll(&_TigerToken.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TigerToken *TigerTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TigerToken *TigerTokenSession) Name() (string, error) {
	return _TigerToken.Contract.Name(&_TigerToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TigerToken *TigerTokenCallerSession) Name() (string, error) {
	return _TigerToken.Contract.Name(&_TigerToken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TigerToken *TigerTokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TigerToken *TigerTokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TigerToken.Contract.OwnerOf(&_TigerToken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TigerToken *TigerTokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TigerToken.Contract.OwnerOf(&_TigerToken.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TigerToken *TigerTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TigerToken *TigerTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TigerToken.Contract.SupportsInterface(&_TigerToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TigerToken *TigerTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TigerToken.Contract.SupportsInterface(&_TigerToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TigerToken *TigerTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TigerToken *TigerTokenSession) Symbol() (string, error) {
	return _TigerToken.Contract.Symbol(&_TigerToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TigerToken *TigerTokenCallerSession) Symbol() (string, error) {
	return _TigerToken.Contract.Symbol(&_TigerToken.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TigerToken *TigerTokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TigerToken *TigerTokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TigerToken.Contract.TokenURI(&_TigerToken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TigerToken *TigerTokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TigerToken.Contract.TokenURI(&_TigerToken.CallOpts, tokenId)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_TigerToken *TigerTokenCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TigerToken.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_TigerToken *TigerTokenSession) Usdt() (common.Address, error) {
	return _TigerToken.Contract.Usdt(&_TigerToken.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_TigerToken *TigerTokenCallerSession) Usdt() (common.Address, error) {
	return _TigerToken.Contract.Usdt(&_TigerToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TigerToken *TigerTokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TigerToken *TigerTokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TigerToken.Contract.Approve(&_TigerToken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TigerToken *TigerTokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TigerToken.Contract.Approve(&_TigerToken.TransactOpts, to, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _usdt) returns()
func (_TigerToken *TigerTokenTransactor) Initialize(opts *bind.TransactOpts, _usdt common.Address) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "initialize", _usdt)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _usdt) returns()
func (_TigerToken *TigerTokenSession) Initialize(_usdt common.Address) (*types.Transaction, error) {
	return _TigerToken.Contract.Initialize(&_TigerToken.TransactOpts, _usdt)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _usdt) returns()
func (_TigerToken *TigerTokenTransactorSession) Initialize(_usdt common.Address) (*types.Transaction, error) {
	return _TigerToken.Contract.Initialize(&_TigerToken.TransactOpts, _usdt)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to, uint256 tokenId, string _tokenURI) returns()
func (_TigerToken *TigerTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "mint", to, tokenId, _tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to, uint256 tokenId, string _tokenURI) returns()
func (_TigerToken *TigerTokenSession) Mint(to common.Address, tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _TigerToken.Contract.Mint(&_TigerToken.TransactOpts, to, tokenId, _tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to, uint256 tokenId, string _tokenURI) returns()
func (_TigerToken *TigerTokenTransactorSession) Mint(to common.Address, tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _TigerToken.Contract.Mint(&_TigerToken.TransactOpts, to, tokenId, _tokenURI)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TigerToken *TigerTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TigerToken *TigerTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TigerToken.Contract.SafeTransferFrom(&_TigerToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TigerToken *TigerTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TigerToken.Contract.SafeTransferFrom(&_TigerToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_TigerToken *TigerTokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_TigerToken *TigerTokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TigerToken.Contract.SafeTransferFrom0(&_TigerToken.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_TigerToken *TigerTokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TigerToken.Contract.SafeTransferFrom0(&_TigerToken.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TigerToken *TigerTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TigerToken *TigerTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TigerToken.Contract.SetApprovalForAll(&_TigerToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TigerToken *TigerTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TigerToken.Contract.SetApprovalForAll(&_TigerToken.TransactOpts, operator, approved)
}

// SetBillingAddress is a paid mutator transaction binding the contract method 0x8048895f.
//
// Solidity: function setBillingAddress(address _newBillingAddress) returns()
func (_TigerToken *TigerTokenTransactor) SetBillingAddress(opts *bind.TransactOpts, _newBillingAddress common.Address) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "setBillingAddress", _newBillingAddress)
}

// SetBillingAddress is a paid mutator transaction binding the contract method 0x8048895f.
//
// Solidity: function setBillingAddress(address _newBillingAddress) returns()
func (_TigerToken *TigerTokenSession) SetBillingAddress(_newBillingAddress common.Address) (*types.Transaction, error) {
	return _TigerToken.Contract.SetBillingAddress(&_TigerToken.TransactOpts, _newBillingAddress)
}

// SetBillingAddress is a paid mutator transaction binding the contract method 0x8048895f.
//
// Solidity: function setBillingAddress(address _newBillingAddress) returns()
func (_TigerToken *TigerTokenTransactorSession) SetBillingAddress(_newBillingAddress common.Address) (*types.Transaction, error) {
	return _TigerToken.Contract.SetBillingAddress(&_TigerToken.TransactOpts, _newBillingAddress)
}

// SetRandom is a paid mutator transaction binding the contract method 0x2043c667.
//
// Solidity: function setRandom(string _random) returns()
func (_TigerToken *TigerTokenTransactor) SetRandom(opts *bind.TransactOpts, _random string) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "setRandom", _random)
}

// SetRandom is a paid mutator transaction binding the contract method 0x2043c667.
//
// Solidity: function setRandom(string _random) returns()
func (_TigerToken *TigerTokenSession) SetRandom(_random string) (*types.Transaction, error) {
	return _TigerToken.Contract.SetRandom(&_TigerToken.TransactOpts, _random)
}

// SetRandom is a paid mutator transaction binding the contract method 0x2043c667.
//
// Solidity: function setRandom(string _random) returns()
func (_TigerToken *TigerTokenTransactorSession) SetRandom(_random string) (*types.Transaction, error) {
	return _TigerToken.Contract.SetRandom(&_TigerToken.TransactOpts, _random)
}

// SetUsdtAddress is a paid mutator transaction binding the contract method 0x0cb46b75.
//
// Solidity: function setUsdtAddress(address _usdtAddress) returns()
func (_TigerToken *TigerTokenTransactor) SetUsdtAddress(opts *bind.TransactOpts, _usdtAddress common.Address) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "setUsdtAddress", _usdtAddress)
}

// SetUsdtAddress is a paid mutator transaction binding the contract method 0x0cb46b75.
//
// Solidity: function setUsdtAddress(address _usdtAddress) returns()
func (_TigerToken *TigerTokenSession) SetUsdtAddress(_usdtAddress common.Address) (*types.Transaction, error) {
	return _TigerToken.Contract.SetUsdtAddress(&_TigerToken.TransactOpts, _usdtAddress)
}

// SetUsdtAddress is a paid mutator transaction binding the contract method 0x0cb46b75.
//
// Solidity: function setUsdtAddress(address _usdtAddress) returns()
func (_TigerToken *TigerTokenTransactorSession) SetUsdtAddress(_usdtAddress common.Address) (*types.Transaction, error) {
	return _TigerToken.Contract.SetUsdtAddress(&_TigerToken.TransactOpts, _usdtAddress)
}

// Subscribe is a paid mutator transaction binding the contract method 0x756ab07f.
//
// Solidity: function subscribe(uint256 tokenId, string _uri, uint256 price, bytes32 verify) returns()
func (_TigerToken *TigerTokenTransactor) Subscribe(opts *bind.TransactOpts, tokenId *big.Int, _uri string, price *big.Int, verify [32]byte) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "subscribe", tokenId, _uri, price, verify)
}

// Subscribe is a paid mutator transaction binding the contract method 0x756ab07f.
//
// Solidity: function subscribe(uint256 tokenId, string _uri, uint256 price, bytes32 verify) returns()
func (_TigerToken *TigerTokenSession) Subscribe(tokenId *big.Int, _uri string, price *big.Int, verify [32]byte) (*types.Transaction, error) {
	return _TigerToken.Contract.Subscribe(&_TigerToken.TransactOpts, tokenId, _uri, price, verify)
}

// Subscribe is a paid mutator transaction binding the contract method 0x756ab07f.
//
// Solidity: function subscribe(uint256 tokenId, string _uri, uint256 price, bytes32 verify) returns()
func (_TigerToken *TigerTokenTransactorSession) Subscribe(tokenId *big.Int, _uri string, price *big.Int, verify [32]byte) (*types.Transaction, error) {
	return _TigerToken.Contract.Subscribe(&_TigerToken.TransactOpts, tokenId, _uri, price, verify)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TigerToken *TigerTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TigerToken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TigerToken *TigerTokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TigerToken.Contract.TransferFrom(&_TigerToken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TigerToken *TigerTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TigerToken.Contract.TransferFrom(&_TigerToken.TransactOpts, from, to, tokenId)
}

// TigerTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TigerToken contract.
type TigerTokenApprovalIterator struct {
	Event *TigerTokenApproval // Event containing the contract specifics and raw log

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
func (it *TigerTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TigerTokenApproval)
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
		it.Event = new(TigerTokenApproval)
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
func (it *TigerTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TigerTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TigerTokenApproval represents a Approval event raised by the TigerToken contract.
type TigerTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TigerToken *TigerTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TigerTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TigerToken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TigerTokenApprovalIterator{contract: _TigerToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TigerToken *TigerTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TigerTokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TigerToken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TigerTokenApproval)
				if err := _TigerToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TigerToken *TigerTokenFilterer) ParseApproval(log types.Log) (*TigerTokenApproval, error) {
	event := new(TigerTokenApproval)
	if err := _TigerToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TigerTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TigerToken contract.
type TigerTokenApprovalForAllIterator struct {
	Event *TigerTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TigerTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TigerTokenApprovalForAll)
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
		it.Event = new(TigerTokenApprovalForAll)
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
func (it *TigerTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TigerTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TigerTokenApprovalForAll represents a ApprovalForAll event raised by the TigerToken contract.
type TigerTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TigerToken *TigerTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TigerTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TigerToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TigerTokenApprovalForAllIterator{contract: _TigerToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TigerToken *TigerTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TigerTokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TigerToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TigerTokenApprovalForAll)
				if err := _TigerToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TigerToken *TigerTokenFilterer) ParseApprovalForAll(log types.Log) (*TigerTokenApprovalForAll, error) {
	event := new(TigerTokenApprovalForAll)
	if err := _TigerToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TigerTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TigerToken contract.
type TigerTokenTransferIterator struct {
	Event *TigerTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TigerTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TigerTokenTransfer)
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
		it.Event = new(TigerTokenTransfer)
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
func (it *TigerTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TigerTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TigerTokenTransfer represents a Transfer event raised by the TigerToken contract.
type TigerTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TigerToken *TigerTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TigerTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TigerToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TigerTokenTransferIterator{contract: _TigerToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TigerToken *TigerTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TigerTokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TigerToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TigerTokenTransfer)
				if err := _TigerToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TigerToken *TigerTokenFilterer) ParseTransfer(log types.Log) (*TigerTokenTransfer, error) {
	event := new(TigerTokenTransfer)
	if err := _TigerToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
