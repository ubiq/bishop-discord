// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package n

import (
	"math/big"
	"strings"

	ethereum "github.com/ubiq/go-ubiq/v7"
	"github.com/ubiq/go-ubiq/v7/accounts/abi"
	"github.com/ubiq/go-ubiq/v7/accounts/abi/bind"
	"github.com/ubiq/go-ubiq/v7/common"
	"github.com/ubiq/go-ubiq/v7/core/types"
	"github.com/ubiq/go-ubiq/v7/event"
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

// NABI is the input ABI used to generate the binding from.
const NABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_burnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_claimPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnerAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getEight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFifth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFirst\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getFourth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSeventh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getSixth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getThird\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_burnerAddress\",\"type\":\"address\"}],\"name\":\"setBurnerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_claimPrice\",\"type\":\"uint256\"}],\"name\":\"setClaimPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// N is an auto generated Go binding around an Ethereum contract.
type N struct {
	NCaller     // Read-only binding to the contract
	NTransactor // Write-only binding to the contract
	NFilterer   // Log filterer for contract events
}

// NCaller is an auto generated read-only Go binding around an Ethereum contract.
type NCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NSession struct {
	Contract     *N                // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NCallerSession struct {
	Contract *NCaller      // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NTransactorSession struct {
	Contract     *NTransactor      // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NRaw is an auto generated low-level Go binding around an Ethereum contract.
type NRaw struct {
	Contract *N // Generic contract binding to access the raw methods on
}

// NCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NCallerRaw struct {
	Contract *NCaller // Generic read-only contract binding to access the raw methods on
}

// NTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NTransactorRaw struct {
	Contract *NTransactor // Generic write-only contract binding to access the raw methods on
}

// NewN creates a new instance of N, bound to a specific deployed contract.
func NewN(address common.Address, backend bind.ContractBackend) (*N, error) {
	contract, err := bindN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &N{NCaller: NCaller{contract: contract}, NTransactor: NTransactor{contract: contract}, NFilterer: NFilterer{contract: contract}}, nil
}

// NewNCaller creates a new read-only instance of N, bound to a specific deployed contract.
func NewNCaller(address common.Address, caller bind.ContractCaller) (*NCaller, error) {
	contract, err := bindN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NCaller{contract: contract}, nil
}

// NewNTransactor creates a new write-only instance of N, bound to a specific deployed contract.
func NewNTransactor(address common.Address, transactor bind.ContractTransactor) (*NTransactor, error) {
	contract, err := bindN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NTransactor{contract: contract}, nil
}

// NewNFilterer creates a new log filterer instance of N, bound to a specific deployed contract.
func NewNFilterer(address common.Address, filterer bind.ContractFilterer) (*NFilterer, error) {
	contract, err := bindN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFilterer{contract: contract}, nil
}

// bindN binds a generic wrapper to an already deployed contract.
func bindN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_N *NRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _N.Contract.NCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_N *NRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _N.Contract.NTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_N *NRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _N.Contract.NTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_N *NCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _N.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_N *NTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _N.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_N *NTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _N.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_N *NCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_N *NSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _N.Contract.BalanceOf(&_N.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_N *NCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _N.Contract.BalanceOf(&_N.CallOpts, owner)
}

// BurnerAddress is a free data retrieval call binding the contract method 0xe6293e23.
//
// Solidity: function burnerAddress() view returns(address)
func (_N *NCaller) BurnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "burnerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnerAddress is a free data retrieval call binding the contract method 0xe6293e23.
//
// Solidity: function burnerAddress() view returns(address)
func (_N *NSession) BurnerAddress() (common.Address, error) {
	return _N.Contract.BurnerAddress(&_N.CallOpts)
}

// BurnerAddress is a free data retrieval call binding the contract method 0xe6293e23.
//
// Solidity: function burnerAddress() view returns(address)
func (_N *NCallerSession) BurnerAddress() (common.Address, error) {
	return _N.Contract.BurnerAddress(&_N.CallOpts)
}

// ClaimPrice is a free data retrieval call binding the contract method 0x15d655c9.
//
// Solidity: function claimPrice() view returns(uint256)
func (_N *NCaller) ClaimPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "claimPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimPrice is a free data retrieval call binding the contract method 0x15d655c9.
//
// Solidity: function claimPrice() view returns(uint256)
func (_N *NSession) ClaimPrice() (*big.Int, error) {
	return _N.Contract.ClaimPrice(&_N.CallOpts)
}

// ClaimPrice is a free data retrieval call binding the contract method 0x15d655c9.
//
// Solidity: function claimPrice() view returns(uint256)
func (_N *NCallerSession) ClaimPrice() (*big.Int, error) {
	return _N.Contract.ClaimPrice(&_N.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_N *NCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_N *NSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _N.Contract.GetApproved(&_N.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_N *NCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _N.Contract.GetApproved(&_N.CallOpts, tokenId)
}

// GetEight is a free data retrieval call binding the contract method 0x9347e43f.
//
// Solidity: function getEight(uint256 tokenId) view returns(uint256)
func (_N *NCaller) GetEight(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "getEight", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEight is a free data retrieval call binding the contract method 0x9347e43f.
//
// Solidity: function getEight(uint256 tokenId) view returns(uint256)
func (_N *NSession) GetEight(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetEight(&_N.CallOpts, tokenId)
}

// GetEight is a free data retrieval call binding the contract method 0x9347e43f.
//
// Solidity: function getEight(uint256 tokenId) view returns(uint256)
func (_N *NCallerSession) GetEight(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetEight(&_N.CallOpts, tokenId)
}

// GetFifth is a free data retrieval call binding the contract method 0x0b2503a6.
//
// Solidity: function getFifth(uint256 tokenId) view returns(uint256)
func (_N *NCaller) GetFifth(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "getFifth", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFifth is a free data retrieval call binding the contract method 0x0b2503a6.
//
// Solidity: function getFifth(uint256 tokenId) view returns(uint256)
func (_N *NSession) GetFifth(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetFifth(&_N.CallOpts, tokenId)
}

// GetFifth is a free data retrieval call binding the contract method 0x0b2503a6.
//
// Solidity: function getFifth(uint256 tokenId) view returns(uint256)
func (_N *NCallerSession) GetFifth(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetFifth(&_N.CallOpts, tokenId)
}

// GetFirst is a free data retrieval call binding the contract method 0x667386f7.
//
// Solidity: function getFirst(uint256 tokenId) view returns(uint256)
func (_N *NCaller) GetFirst(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "getFirst", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFirst is a free data retrieval call binding the contract method 0x667386f7.
//
// Solidity: function getFirst(uint256 tokenId) view returns(uint256)
func (_N *NSession) GetFirst(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetFirst(&_N.CallOpts, tokenId)
}

// GetFirst is a free data retrieval call binding the contract method 0x667386f7.
//
// Solidity: function getFirst(uint256 tokenId) view returns(uint256)
func (_N *NCallerSession) GetFirst(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetFirst(&_N.CallOpts, tokenId)
}

// GetFourth is a free data retrieval call binding the contract method 0x9ec29b82.
//
// Solidity: function getFourth(uint256 tokenId) view returns(uint256)
func (_N *NCaller) GetFourth(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "getFourth", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFourth is a free data retrieval call binding the contract method 0x9ec29b82.
//
// Solidity: function getFourth(uint256 tokenId) view returns(uint256)
func (_N *NSession) GetFourth(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetFourth(&_N.CallOpts, tokenId)
}

// GetFourth is a free data retrieval call binding the contract method 0x9ec29b82.
//
// Solidity: function getFourth(uint256 tokenId) view returns(uint256)
func (_N *NCallerSession) GetFourth(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetFourth(&_N.CallOpts, tokenId)
}

// GetSecond is a free data retrieval call binding the contract method 0x8aa001fc.
//
// Solidity: function getSecond(uint256 tokenId) view returns(uint256)
func (_N *NCaller) GetSecond(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "getSecond", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSecond is a free data retrieval call binding the contract method 0x8aa001fc.
//
// Solidity: function getSecond(uint256 tokenId) view returns(uint256)
func (_N *NSession) GetSecond(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetSecond(&_N.CallOpts, tokenId)
}

// GetSecond is a free data retrieval call binding the contract method 0x8aa001fc.
//
// Solidity: function getSecond(uint256 tokenId) view returns(uint256)
func (_N *NCallerSession) GetSecond(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetSecond(&_N.CallOpts, tokenId)
}

// GetSeventh is a free data retrieval call binding the contract method 0x8c921d06.
//
// Solidity: function getSeventh(uint256 tokenId) view returns(uint256)
func (_N *NCaller) GetSeventh(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "getSeventh", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSeventh is a free data retrieval call binding the contract method 0x8c921d06.
//
// Solidity: function getSeventh(uint256 tokenId) view returns(uint256)
func (_N *NSession) GetSeventh(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetSeventh(&_N.CallOpts, tokenId)
}

// GetSeventh is a free data retrieval call binding the contract method 0x8c921d06.
//
// Solidity: function getSeventh(uint256 tokenId) view returns(uint256)
func (_N *NCallerSession) GetSeventh(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetSeventh(&_N.CallOpts, tokenId)
}

// GetSixth is a free data retrieval call binding the contract method 0x42d9d876.
//
// Solidity: function getSixth(uint256 tokenId) view returns(uint256)
func (_N *NCaller) GetSixth(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "getSixth", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSixth is a free data retrieval call binding the contract method 0x42d9d876.
//
// Solidity: function getSixth(uint256 tokenId) view returns(uint256)
func (_N *NSession) GetSixth(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetSixth(&_N.CallOpts, tokenId)
}

// GetSixth is a free data retrieval call binding the contract method 0x42d9d876.
//
// Solidity: function getSixth(uint256 tokenId) view returns(uint256)
func (_N *NCallerSession) GetSixth(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetSixth(&_N.CallOpts, tokenId)
}

// GetThird is a free data retrieval call binding the contract method 0xfa7f71b1.
//
// Solidity: function getThird(uint256 tokenId) view returns(uint256)
func (_N *NCaller) GetThird(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "getThird", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThird is a free data retrieval call binding the contract method 0xfa7f71b1.
//
// Solidity: function getThird(uint256 tokenId) view returns(uint256)
func (_N *NSession) GetThird(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetThird(&_N.CallOpts, tokenId)
}

// GetThird is a free data retrieval call binding the contract method 0xfa7f71b1.
//
// Solidity: function getThird(uint256 tokenId) view returns(uint256)
func (_N *NCallerSession) GetThird(tokenId *big.Int) (*big.Int, error) {
	return _N.Contract.GetThird(&_N.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_N *NCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_N *NSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _N.Contract.IsApprovedForAll(&_N.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_N *NCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _N.Contract.IsApprovedForAll(&_N.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_N *NCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_N *NSession) Name() (string, error) {
	return _N.Contract.Name(&_N.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_N *NCallerSession) Name() (string, error) {
	return _N.Contract.Name(&_N.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_N *NCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_N *NSession) Owner() (common.Address, error) {
	return _N.Contract.Owner(&_N.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_N *NCallerSession) Owner() (common.Address, error) {
	return _N.Contract.Owner(&_N.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_N *NCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_N *NSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _N.Contract.OwnerOf(&_N.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_N *NCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _N.Contract.OwnerOf(&_N.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_N *NCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_N *NSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _N.Contract.SupportsInterface(&_N.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_N *NCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _N.Contract.SupportsInterface(&_N.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_N *NCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_N *NSession) Symbol() (string, error) {
	return _N.Contract.Symbol(&_N.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_N *NCallerSession) Symbol() (string, error) {
	return _N.Contract.Symbol(&_N.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_N *NCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_N *NSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _N.Contract.TokenByIndex(&_N.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_N *NCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _N.Contract.TokenByIndex(&_N.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_N *NCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_N *NSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _N.Contract.TokenOfOwnerByIndex(&_N.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_N *NCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _N.Contract.TokenOfOwnerByIndex(&_N.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_N *NCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_N *NSession) TokenURI(tokenId *big.Int) (string, error) {
	return _N.Contract.TokenURI(&_N.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_N *NCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _N.Contract.TokenURI(&_N.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_N *NCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _N.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_N *NSession) TotalSupply() (*big.Int, error) {
	return _N.Contract.TotalSupply(&_N.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_N *NCallerSession) TotalSupply() (*big.Int, error) {
	return _N.Contract.TotalSupply(&_N.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_N *NTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_N *NSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _N.Contract.Approve(&_N.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_N *NTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _N.Contract.Approve(&_N.TransactOpts, to, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 tokenId) payable returns()
func (_N *NTransactor) Claim(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "claim", tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 tokenId) payable returns()
func (_N *NSession) Claim(tokenId *big.Int) (*types.Transaction, error) {
	return _N.Contract.Claim(&_N.TransactOpts, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 tokenId) payable returns()
func (_N *NTransactorSession) Claim(tokenId *big.Int) (*types.Transaction, error) {
	return _N.Contract.Claim(&_N.TransactOpts, tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_N *NTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_N *NSession) RenounceOwnership() (*types.Transaction, error) {
	return _N.Contract.RenounceOwnership(&_N.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_N *NTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _N.Contract.RenounceOwnership(&_N.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_N *NTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_N *NSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _N.Contract.SafeTransferFrom(&_N.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_N *NTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _N.Contract.SafeTransferFrom(&_N.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_N *NTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_N *NSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _N.Contract.SafeTransferFrom0(&_N.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_N *NTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _N.Contract.SafeTransferFrom0(&_N.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_N *NTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_N *NSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _N.Contract.SetApprovalForAll(&_N.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_N *NTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _N.Contract.SetApprovalForAll(&_N.TransactOpts, operator, approved)
}

// SetBurnerAddress is a paid mutator transaction binding the contract method 0x9df806d6.
//
// Solidity: function setBurnerAddress(address _burnerAddress) returns()
func (_N *NTransactor) SetBurnerAddress(opts *bind.TransactOpts, _burnerAddress common.Address) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "setBurnerAddress", _burnerAddress)
}

// SetBurnerAddress is a paid mutator transaction binding the contract method 0x9df806d6.
//
// Solidity: function setBurnerAddress(address _burnerAddress) returns()
func (_N *NSession) SetBurnerAddress(_burnerAddress common.Address) (*types.Transaction, error) {
	return _N.Contract.SetBurnerAddress(&_N.TransactOpts, _burnerAddress)
}

// SetBurnerAddress is a paid mutator transaction binding the contract method 0x9df806d6.
//
// Solidity: function setBurnerAddress(address _burnerAddress) returns()
func (_N *NTransactorSession) SetBurnerAddress(_burnerAddress common.Address) (*types.Transaction, error) {
	return _N.Contract.SetBurnerAddress(&_N.TransactOpts, _burnerAddress)
}

// SetClaimPrice is a paid mutator transaction binding the contract method 0x51f468c0.
//
// Solidity: function setClaimPrice(uint256 _claimPrice) returns()
func (_N *NTransactor) SetClaimPrice(opts *bind.TransactOpts, _claimPrice *big.Int) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "setClaimPrice", _claimPrice)
}

// SetClaimPrice is a paid mutator transaction binding the contract method 0x51f468c0.
//
// Solidity: function setClaimPrice(uint256 _claimPrice) returns()
func (_N *NSession) SetClaimPrice(_claimPrice *big.Int) (*types.Transaction, error) {
	return _N.Contract.SetClaimPrice(&_N.TransactOpts, _claimPrice)
}

// SetClaimPrice is a paid mutator transaction binding the contract method 0x51f468c0.
//
// Solidity: function setClaimPrice(uint256 _claimPrice) returns()
func (_N *NTransactorSession) SetClaimPrice(_claimPrice *big.Int) (*types.Transaction, error) {
	return _N.Contract.SetClaimPrice(&_N.TransactOpts, _claimPrice)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_N *NTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_N *NSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _N.Contract.TransferFrom(&_N.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_N *NTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _N.Contract.TransferFrom(&_N.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_N *NTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _N.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_N *NSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _N.Contract.TransferOwnership(&_N.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_N *NTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _N.Contract.TransferOwnership(&_N.TransactOpts, newOwner)
}

// NApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the N contract.
type NApprovalIterator struct {
	Event *NApproval // Event containing the contract specifics and raw log

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
func (it *NApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NApproval)
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
		it.Event = new(NApproval)
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
func (it *NApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NApproval represents a Approval event raised by the N contract.
type NApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_N *NFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NApprovalIterator, error) {

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

	logs, sub, err := _N.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NApprovalIterator{contract: _N.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_N *NFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _N.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NApproval)
				if err := _N.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_N *NFilterer) ParseApproval(log types.Log) (*NApproval, error) {
	event := new(NApproval)
	if err := _N.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the N contract.
type NApprovalForAllIterator struct {
	Event *NApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NApprovalForAll)
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
		it.Event = new(NApprovalForAll)
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
func (it *NApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NApprovalForAll represents a ApprovalForAll event raised by the N contract.
type NApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_N *NFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _N.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NApprovalForAllIterator{contract: _N.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_N *NFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _N.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NApprovalForAll)
				if err := _N.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_N *NFilterer) ParseApprovalForAll(log types.Log) (*NApprovalForAll, error) {
	event := new(NApprovalForAll)
	if err := _N.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the N contract.
type NOwnershipTransferredIterator struct {
	Event *NOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NOwnershipTransferred)
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
		it.Event = new(NOwnershipTransferred)
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
func (it *NOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NOwnershipTransferred represents a OwnershipTransferred event raised by the N contract.
type NOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_N *NFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _N.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NOwnershipTransferredIterator{contract: _N.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_N *NFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _N.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NOwnershipTransferred)
				if err := _N.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_N *NFilterer) ParseOwnershipTransferred(log types.Log) (*NOwnershipTransferred, error) {
	event := new(NOwnershipTransferred)
	if err := _N.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the N contract.
type NTransferIterator struct {
	Event *NTransfer // Event containing the contract specifics and raw log

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
func (it *NTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NTransfer)
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
		it.Event = new(NTransfer)
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
func (it *NTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NTransfer represents a Transfer event raised by the N contract.
type NTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_N *NFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NTransferIterator, error) {

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

	logs, sub, err := _N.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NTransferIterator{contract: _N.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_N *NFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _N.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NTransfer)
				if err := _N.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_N *NFilterer) ParseTransfer(log types.Log) (*NTransfer, error) {
	event := new(NTransfer)
	if err := _N.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
