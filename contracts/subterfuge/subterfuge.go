// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subterfuge

import (
	"math/big"
	"strings"

	ethereum "github.com/ubiq/go-ubiq/v5"
	"github.com/ubiq/go-ubiq/v5/accounts/abi"
	"github.com/ubiq/go-ubiq/v5/accounts/abi/bind"
	"github.com/ubiq/go-ubiq/v5/common"
	"github.com/ubiq/go-ubiq/v5/core/types"
	"github.com/ubiq/go-ubiq/v5/event"
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

// SubterfugeABI is the input ABI used to generate the binding from.
const SubterfugeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceForNHoldersInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceForOpenMintInWei\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_burnerAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MULTI_MINT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_N_TOKEN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnerAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRibbonInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTierInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mintWithN\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"multiMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"multiMintWithN\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"n\",\"outputs\":[{\"internalType\":\"contractIN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nHoldersMintsAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onlyNHolders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openMintsAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceForNHoldersInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceForOpenMintInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicSale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveMinted\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservedAllowance\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_burnerAddress\",\"type\":\"address\"}],\"name\":\"setBurnerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_publicSale\",\"type\":\"bool\"}],\"name\":\"setPublicSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenSVG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Subterfuge is an auto generated Go binding around an Ethereum contract.
type Subterfuge struct {
	SubterfugeCaller     // Read-only binding to the contract
	SubterfugeTransactor // Write-only binding to the contract
	SubterfugeFilterer   // Log filterer for contract events
}

// SubterfugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubterfugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubterfugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubterfugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubterfugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubterfugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubterfugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubterfugeSession struct {
	Contract     *Subterfuge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubterfugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubterfugeCallerSession struct {
	Contract *SubterfugeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SubterfugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubterfugeTransactorSession struct {
	Contract     *SubterfugeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SubterfugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubterfugeRaw struct {
	Contract *Subterfuge // Generic contract binding to access the raw methods on
}

// SubterfugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubterfugeCallerRaw struct {
	Contract *SubterfugeCaller // Generic read-only contract binding to access the raw methods on
}

// SubterfugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubterfugeTransactorRaw struct {
	Contract *SubterfugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubterfuge creates a new instance of Subterfuge, bound to a specific deployed contract.
func NewSubterfuge(address common.Address, backend bind.ContractBackend) (*Subterfuge, error) {
	contract, err := bindSubterfuge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subterfuge{SubterfugeCaller: SubterfugeCaller{contract: contract}, SubterfugeTransactor: SubterfugeTransactor{contract: contract}, SubterfugeFilterer: SubterfugeFilterer{contract: contract}}, nil
}

// NewSubterfugeCaller creates a new read-only instance of Subterfuge, bound to a specific deployed contract.
func NewSubterfugeCaller(address common.Address, caller bind.ContractCaller) (*SubterfugeCaller, error) {
	contract, err := bindSubterfuge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubterfugeCaller{contract: contract}, nil
}

// NewSubterfugeTransactor creates a new write-only instance of Subterfuge, bound to a specific deployed contract.
func NewSubterfugeTransactor(address common.Address, transactor bind.ContractTransactor) (*SubterfugeTransactor, error) {
	contract, err := bindSubterfuge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubterfugeTransactor{contract: contract}, nil
}

// NewSubterfugeFilterer creates a new log filterer instance of Subterfuge, bound to a specific deployed contract.
func NewSubterfugeFilterer(address common.Address, filterer bind.ContractFilterer) (*SubterfugeFilterer, error) {
	contract, err := bindSubterfuge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubterfugeFilterer{contract: contract}, nil
}

// bindSubterfuge binds a generic wrapper to an already deployed contract.
func bindSubterfuge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SubterfugeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subterfuge *SubterfugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subterfuge.Contract.SubterfugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subterfuge *SubterfugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subterfuge.Contract.SubterfugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subterfuge *SubterfugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subterfuge.Contract.SubterfugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subterfuge *SubterfugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subterfuge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subterfuge *SubterfugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subterfuge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subterfuge *SubterfugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subterfuge.Contract.contract.Transact(opts, method, params...)
}

// MAXMULTIMINTAMOUNT is a free data retrieval call binding the contract method 0x5d929f70.
//
// Solidity: function MAX_MULTI_MINT_AMOUNT() view returns(uint256)
func (_Subterfuge *SubterfugeCaller) MAXMULTIMINTAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "MAX_MULTI_MINT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMULTIMINTAMOUNT is a free data retrieval call binding the contract method 0x5d929f70.
//
// Solidity: function MAX_MULTI_MINT_AMOUNT() view returns(uint256)
func (_Subterfuge *SubterfugeSession) MAXMULTIMINTAMOUNT() (*big.Int, error) {
	return _Subterfuge.Contract.MAXMULTIMINTAMOUNT(&_Subterfuge.CallOpts)
}

// MAXMULTIMINTAMOUNT is a free data retrieval call binding the contract method 0x5d929f70.
//
// Solidity: function MAX_MULTI_MINT_AMOUNT() view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) MAXMULTIMINTAMOUNT() (*big.Int, error) {
	return _Subterfuge.Contract.MAXMULTIMINTAMOUNT(&_Subterfuge.CallOpts)
}

// MAXNTOKENID is a free data retrieval call binding the contract method 0xae5a583f.
//
// Solidity: function MAX_N_TOKEN_ID() view returns(uint256)
func (_Subterfuge *SubterfugeCaller) MAXNTOKENID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "MAX_N_TOKEN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNTOKENID is a free data retrieval call binding the contract method 0xae5a583f.
//
// Solidity: function MAX_N_TOKEN_ID() view returns(uint256)
func (_Subterfuge *SubterfugeSession) MAXNTOKENID() (*big.Int, error) {
	return _Subterfuge.Contract.MAXNTOKENID(&_Subterfuge.CallOpts)
}

// MAXNTOKENID is a free data retrieval call binding the contract method 0xae5a583f.
//
// Solidity: function MAX_N_TOKEN_ID() view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) MAXNTOKENID() (*big.Int, error) {
	return _Subterfuge.Contract.MAXNTOKENID(&_Subterfuge.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Subterfuge *SubterfugeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Subterfuge *SubterfugeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Subterfuge.Contract.BalanceOf(&_Subterfuge.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Subterfuge.Contract.BalanceOf(&_Subterfuge.CallOpts, owner)
}

// BurnerAddress is a free data retrieval call binding the contract method 0xe6293e23.
//
// Solidity: function burnerAddress() view returns(address)
func (_Subterfuge *SubterfugeCaller) BurnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "burnerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnerAddress is a free data retrieval call binding the contract method 0xe6293e23.
//
// Solidity: function burnerAddress() view returns(address)
func (_Subterfuge *SubterfugeSession) BurnerAddress() (common.Address, error) {
	return _Subterfuge.Contract.BurnerAddress(&_Subterfuge.CallOpts)
}

// BurnerAddress is a free data retrieval call binding the contract method 0xe6293e23.
//
// Solidity: function burnerAddress() view returns(address)
func (_Subterfuge *SubterfugeCallerSession) BurnerAddress() (common.Address, error) {
	return _Subterfuge.Contract.BurnerAddress(&_Subterfuge.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Subterfuge *SubterfugeCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Subterfuge *SubterfugeSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Subterfuge.Contract.GetApproved(&_Subterfuge.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Subterfuge *SubterfugeCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Subterfuge.Contract.GetApproved(&_Subterfuge.CallOpts, tokenId)
}

// GetRibbonInformation is a free data retrieval call binding the contract method 0xe412c7f1.
//
// Solidity: function getRibbonInformation(uint256 tokenId) view returns(string, string)
func (_Subterfuge *SubterfugeCaller) GetRibbonInformation(opts *bind.CallOpts, tokenId *big.Int) (string, string, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "getRibbonInformation", tokenId)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetRibbonInformation is a free data retrieval call binding the contract method 0xe412c7f1.
//
// Solidity: function getRibbonInformation(uint256 tokenId) view returns(string, string)
func (_Subterfuge *SubterfugeSession) GetRibbonInformation(tokenId *big.Int) (string, string, error) {
	return _Subterfuge.Contract.GetRibbonInformation(&_Subterfuge.CallOpts, tokenId)
}

// GetRibbonInformation is a free data retrieval call binding the contract method 0xe412c7f1.
//
// Solidity: function getRibbonInformation(uint256 tokenId) view returns(string, string)
func (_Subterfuge *SubterfugeCallerSession) GetRibbonInformation(tokenId *big.Int) (string, string, error) {
	return _Subterfuge.Contract.GetRibbonInformation(&_Subterfuge.CallOpts, tokenId)
}

// GetTierInformation is a free data retrieval call binding the contract method 0x13f4d737.
//
// Solidity: function getTierInformation(uint256 tokenId) view returns(string, string)
func (_Subterfuge *SubterfugeCaller) GetTierInformation(opts *bind.CallOpts, tokenId *big.Int) (string, string, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "getTierInformation", tokenId)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetTierInformation is a free data retrieval call binding the contract method 0x13f4d737.
//
// Solidity: function getTierInformation(uint256 tokenId) view returns(string, string)
func (_Subterfuge *SubterfugeSession) GetTierInformation(tokenId *big.Int) (string, string, error) {
	return _Subterfuge.Contract.GetTierInformation(&_Subterfuge.CallOpts, tokenId)
}

// GetTierInformation is a free data retrieval call binding the contract method 0x13f4d737.
//
// Solidity: function getTierInformation(uint256 tokenId) view returns(string, string)
func (_Subterfuge *SubterfugeCallerSession) GetTierInformation(tokenId *big.Int) (string, string, error) {
	return _Subterfuge.Contract.GetTierInformation(&_Subterfuge.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Subterfuge *SubterfugeCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Subterfuge *SubterfugeSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Subterfuge.Contract.IsApprovedForAll(&_Subterfuge.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Subterfuge *SubterfugeCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Subterfuge.Contract.IsApprovedForAll(&_Subterfuge.CallOpts, owner, operator)
}

// MaxTokenId is a free data retrieval call binding the contract method 0x91ba317a.
//
// Solidity: function maxTokenId() view returns(uint256)
func (_Subterfuge *SubterfugeCaller) MaxTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "maxTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTokenId is a free data retrieval call binding the contract method 0x91ba317a.
//
// Solidity: function maxTokenId() view returns(uint256)
func (_Subterfuge *SubterfugeSession) MaxTokenId() (*big.Int, error) {
	return _Subterfuge.Contract.MaxTokenId(&_Subterfuge.CallOpts)
}

// MaxTokenId is a free data retrieval call binding the contract method 0x91ba317a.
//
// Solidity: function maxTokenId() view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) MaxTokenId() (*big.Int, error) {
	return _Subterfuge.Contract.MaxTokenId(&_Subterfuge.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Subterfuge *SubterfugeCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Subterfuge *SubterfugeSession) MaxTotalSupply() (*big.Int, error) {
	return _Subterfuge.Contract.MaxTotalSupply(&_Subterfuge.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _Subterfuge.Contract.MaxTotalSupply(&_Subterfuge.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(address)
func (_Subterfuge *SubterfugeCaller) N(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "n")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(address)
func (_Subterfuge *SubterfugeSession) N() (common.Address, error) {
	return _Subterfuge.Contract.N(&_Subterfuge.CallOpts)
}

// N is a free data retrieval call binding the contract method 0x2e52d606.
//
// Solidity: function n() view returns(address)
func (_Subterfuge *SubterfugeCallerSession) N() (common.Address, error) {
	return _Subterfuge.Contract.N(&_Subterfuge.CallOpts)
}

// NHoldersMintsAvailable is a free data retrieval call binding the contract method 0xc5de34a0.
//
// Solidity: function nHoldersMintsAvailable() view returns(uint256)
func (_Subterfuge *SubterfugeCaller) NHoldersMintsAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "nHoldersMintsAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NHoldersMintsAvailable is a free data retrieval call binding the contract method 0xc5de34a0.
//
// Solidity: function nHoldersMintsAvailable() view returns(uint256)
func (_Subterfuge *SubterfugeSession) NHoldersMintsAvailable() (*big.Int, error) {
	return _Subterfuge.Contract.NHoldersMintsAvailable(&_Subterfuge.CallOpts)
}

// NHoldersMintsAvailable is a free data retrieval call binding the contract method 0xc5de34a0.
//
// Solidity: function nHoldersMintsAvailable() view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) NHoldersMintsAvailable() (*big.Int, error) {
	return _Subterfuge.Contract.NHoldersMintsAvailable(&_Subterfuge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Subterfuge *SubterfugeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Subterfuge *SubterfugeSession) Name() (string, error) {
	return _Subterfuge.Contract.Name(&_Subterfuge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Subterfuge *SubterfugeCallerSession) Name() (string, error) {
	return _Subterfuge.Contract.Name(&_Subterfuge.CallOpts)
}

// OnlyNHolders is a free data retrieval call binding the contract method 0x6a4c19d9.
//
// Solidity: function onlyNHolders() view returns(bool)
func (_Subterfuge *SubterfugeCaller) OnlyNHolders(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "onlyNHolders")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyNHolders is a free data retrieval call binding the contract method 0x6a4c19d9.
//
// Solidity: function onlyNHolders() view returns(bool)
func (_Subterfuge *SubterfugeSession) OnlyNHolders() (bool, error) {
	return _Subterfuge.Contract.OnlyNHolders(&_Subterfuge.CallOpts)
}

// OnlyNHolders is a free data retrieval call binding the contract method 0x6a4c19d9.
//
// Solidity: function onlyNHolders() view returns(bool)
func (_Subterfuge *SubterfugeCallerSession) OnlyNHolders() (bool, error) {
	return _Subterfuge.Contract.OnlyNHolders(&_Subterfuge.CallOpts)
}

// OpenMintsAvailable is a free data retrieval call binding the contract method 0x8416b696.
//
// Solidity: function openMintsAvailable() view returns(uint256)
func (_Subterfuge *SubterfugeCaller) OpenMintsAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "openMintsAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenMintsAvailable is a free data retrieval call binding the contract method 0x8416b696.
//
// Solidity: function openMintsAvailable() view returns(uint256)
func (_Subterfuge *SubterfugeSession) OpenMintsAvailable() (*big.Int, error) {
	return _Subterfuge.Contract.OpenMintsAvailable(&_Subterfuge.CallOpts)
}

// OpenMintsAvailable is a free data retrieval call binding the contract method 0x8416b696.
//
// Solidity: function openMintsAvailable() view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) OpenMintsAvailable() (*big.Int, error) {
	return _Subterfuge.Contract.OpenMintsAvailable(&_Subterfuge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Subterfuge *SubterfugeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Subterfuge *SubterfugeSession) Owner() (common.Address, error) {
	return _Subterfuge.Contract.Owner(&_Subterfuge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Subterfuge *SubterfugeCallerSession) Owner() (common.Address, error) {
	return _Subterfuge.Contract.Owner(&_Subterfuge.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Subterfuge *SubterfugeCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Subterfuge *SubterfugeSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Subterfuge.Contract.OwnerOf(&_Subterfuge.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Subterfuge *SubterfugeCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Subterfuge.Contract.OwnerOf(&_Subterfuge.CallOpts, tokenId)
}

// PriceForNHoldersInWei is a free data retrieval call binding the contract method 0xdaf5d374.
//
// Solidity: function priceForNHoldersInWei() view returns(uint256)
func (_Subterfuge *SubterfugeCaller) PriceForNHoldersInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "priceForNHoldersInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceForNHoldersInWei is a free data retrieval call binding the contract method 0xdaf5d374.
//
// Solidity: function priceForNHoldersInWei() view returns(uint256)
func (_Subterfuge *SubterfugeSession) PriceForNHoldersInWei() (*big.Int, error) {
	return _Subterfuge.Contract.PriceForNHoldersInWei(&_Subterfuge.CallOpts)
}

// PriceForNHoldersInWei is a free data retrieval call binding the contract method 0xdaf5d374.
//
// Solidity: function priceForNHoldersInWei() view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) PriceForNHoldersInWei() (*big.Int, error) {
	return _Subterfuge.Contract.PriceForNHoldersInWei(&_Subterfuge.CallOpts)
}

// PriceForOpenMintInWei is a free data retrieval call binding the contract method 0xf82cbbe8.
//
// Solidity: function priceForOpenMintInWei() view returns(uint256)
func (_Subterfuge *SubterfugeCaller) PriceForOpenMintInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "priceForOpenMintInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceForOpenMintInWei is a free data retrieval call binding the contract method 0xf82cbbe8.
//
// Solidity: function priceForOpenMintInWei() view returns(uint256)
func (_Subterfuge *SubterfugeSession) PriceForOpenMintInWei() (*big.Int, error) {
	return _Subterfuge.Contract.PriceForOpenMintInWei(&_Subterfuge.CallOpts)
}

// PriceForOpenMintInWei is a free data retrieval call binding the contract method 0xf82cbbe8.
//
// Solidity: function priceForOpenMintInWei() view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) PriceForOpenMintInWei() (*big.Int, error) {
	return _Subterfuge.Contract.PriceForOpenMintInWei(&_Subterfuge.CallOpts)
}

// PublicSale is a free data retrieval call binding the contract method 0x33bc1c5c.
//
// Solidity: function publicSale() view returns(bool)
func (_Subterfuge *SubterfugeCaller) PublicSale(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "publicSale")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicSale is a free data retrieval call binding the contract method 0x33bc1c5c.
//
// Solidity: function publicSale() view returns(bool)
func (_Subterfuge *SubterfugeSession) PublicSale() (bool, error) {
	return _Subterfuge.Contract.PublicSale(&_Subterfuge.CallOpts)
}

// PublicSale is a free data retrieval call binding the contract method 0x33bc1c5c.
//
// Solidity: function publicSale() view returns(bool)
func (_Subterfuge *SubterfugeCallerSession) PublicSale() (bool, error) {
	return _Subterfuge.Contract.PublicSale(&_Subterfuge.CallOpts)
}

// ReserveMinted is a free data retrieval call binding the contract method 0x4c81433f.
//
// Solidity: function reserveMinted() view returns(uint16)
func (_Subterfuge *SubterfugeCaller) ReserveMinted(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "reserveMinted")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ReserveMinted is a free data retrieval call binding the contract method 0x4c81433f.
//
// Solidity: function reserveMinted() view returns(uint16)
func (_Subterfuge *SubterfugeSession) ReserveMinted() (uint16, error) {
	return _Subterfuge.Contract.ReserveMinted(&_Subterfuge.CallOpts)
}

// ReserveMinted is a free data retrieval call binding the contract method 0x4c81433f.
//
// Solidity: function reserveMinted() view returns(uint16)
func (_Subterfuge *SubterfugeCallerSession) ReserveMinted() (uint16, error) {
	return _Subterfuge.Contract.ReserveMinted(&_Subterfuge.CallOpts)
}

// ReservedAllowance is a free data retrieval call binding the contract method 0x20bc84ce.
//
// Solidity: function reservedAllowance() view returns(uint16)
func (_Subterfuge *SubterfugeCaller) ReservedAllowance(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "reservedAllowance")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ReservedAllowance is a free data retrieval call binding the contract method 0x20bc84ce.
//
// Solidity: function reservedAllowance() view returns(uint16)
func (_Subterfuge *SubterfugeSession) ReservedAllowance() (uint16, error) {
	return _Subterfuge.Contract.ReservedAllowance(&_Subterfuge.CallOpts)
}

// ReservedAllowance is a free data retrieval call binding the contract method 0x20bc84ce.
//
// Solidity: function reservedAllowance() view returns(uint16)
func (_Subterfuge *SubterfugeCallerSession) ReservedAllowance() (uint16, error) {
	return _Subterfuge.Contract.ReservedAllowance(&_Subterfuge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Subterfuge *SubterfugeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Subterfuge *SubterfugeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Subterfuge.Contract.SupportsInterface(&_Subterfuge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Subterfuge *SubterfugeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Subterfuge.Contract.SupportsInterface(&_Subterfuge.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Subterfuge *SubterfugeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Subterfuge *SubterfugeSession) Symbol() (string, error) {
	return _Subterfuge.Contract.Symbol(&_Subterfuge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Subterfuge *SubterfugeCallerSession) Symbol() (string, error) {
	return _Subterfuge.Contract.Symbol(&_Subterfuge.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Subterfuge *SubterfugeCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Subterfuge *SubterfugeSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Subterfuge.Contract.TokenByIndex(&_Subterfuge.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Subterfuge.Contract.TokenByIndex(&_Subterfuge.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Subterfuge *SubterfugeCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Subterfuge *SubterfugeSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Subterfuge.Contract.TokenOfOwnerByIndex(&_Subterfuge.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Subterfuge.Contract.TokenOfOwnerByIndex(&_Subterfuge.CallOpts, owner, index)
}

// TokenSVG is a free data retrieval call binding the contract method 0x9bac5f7a.
//
// Solidity: function tokenSVG(uint256 tokenId) view returns(string)
func (_Subterfuge *SubterfugeCaller) TokenSVG(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "tokenSVG", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenSVG is a free data retrieval call binding the contract method 0x9bac5f7a.
//
// Solidity: function tokenSVG(uint256 tokenId) view returns(string)
func (_Subterfuge *SubterfugeSession) TokenSVG(tokenId *big.Int) (string, error) {
	return _Subterfuge.Contract.TokenSVG(&_Subterfuge.CallOpts, tokenId)
}

// TokenSVG is a free data retrieval call binding the contract method 0x9bac5f7a.
//
// Solidity: function tokenSVG(uint256 tokenId) view returns(string)
func (_Subterfuge *SubterfugeCallerSession) TokenSVG(tokenId *big.Int) (string, error) {
	return _Subterfuge.Contract.TokenSVG(&_Subterfuge.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Subterfuge *SubterfugeCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Subterfuge *SubterfugeSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Subterfuge.Contract.TokenURI(&_Subterfuge.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Subterfuge *SubterfugeCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Subterfuge.Contract.TokenURI(&_Subterfuge.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Subterfuge *SubterfugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subterfuge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Subterfuge *SubterfugeSession) TotalSupply() (*big.Int, error) {
	return _Subterfuge.Contract.TotalSupply(&_Subterfuge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Subterfuge *SubterfugeCallerSession) TotalSupply() (*big.Int, error) {
	return _Subterfuge.Contract.TotalSupply(&_Subterfuge.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Subterfuge *SubterfugeTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Subterfuge *SubterfugeSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.Approve(&_Subterfuge.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Subterfuge *SubterfugeTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.Approve(&_Subterfuge.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) payable returns()
func (_Subterfuge *SubterfugeTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "mint", tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) payable returns()
func (_Subterfuge *SubterfugeSession) Mint(tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.Mint(&_Subterfuge.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) payable returns()
func (_Subterfuge *SubterfugeTransactorSession) Mint(tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.Mint(&_Subterfuge.TransactOpts, tokenId)
}

// MintWithN is a paid mutator transaction binding the contract method 0x0860b12c.
//
// Solidity: function mintWithN(uint256 tokenId) payable returns()
func (_Subterfuge *SubterfugeTransactor) MintWithN(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "mintWithN", tokenId)
}

// MintWithN is a paid mutator transaction binding the contract method 0x0860b12c.
//
// Solidity: function mintWithN(uint256 tokenId) payable returns()
func (_Subterfuge *SubterfugeSession) MintWithN(tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.MintWithN(&_Subterfuge.TransactOpts, tokenId)
}

// MintWithN is a paid mutator transaction binding the contract method 0x0860b12c.
//
// Solidity: function mintWithN(uint256 tokenId) payable returns()
func (_Subterfuge *SubterfugeTransactorSession) MintWithN(tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.MintWithN(&_Subterfuge.TransactOpts, tokenId)
}

// MultiMint is a paid mutator transaction binding the contract method 0x9a0e4ebb.
//
// Solidity: function multiMint(uint256[] tokenIds) payable returns()
func (_Subterfuge *SubterfugeTransactor) MultiMint(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "multiMint", tokenIds)
}

// MultiMint is a paid mutator transaction binding the contract method 0x9a0e4ebb.
//
// Solidity: function multiMint(uint256[] tokenIds) payable returns()
func (_Subterfuge *SubterfugeSession) MultiMint(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.MultiMint(&_Subterfuge.TransactOpts, tokenIds)
}

// MultiMint is a paid mutator transaction binding the contract method 0x9a0e4ebb.
//
// Solidity: function multiMint(uint256[] tokenIds) payable returns()
func (_Subterfuge *SubterfugeTransactorSession) MultiMint(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.MultiMint(&_Subterfuge.TransactOpts, tokenIds)
}

// MultiMintWithN is a paid mutator transaction binding the contract method 0x47febae8.
//
// Solidity: function multiMintWithN(uint256[] tokenIds) payable returns()
func (_Subterfuge *SubterfugeTransactor) MultiMintWithN(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "multiMintWithN", tokenIds)
}

// MultiMintWithN is a paid mutator transaction binding the contract method 0x47febae8.
//
// Solidity: function multiMintWithN(uint256[] tokenIds) payable returns()
func (_Subterfuge *SubterfugeSession) MultiMintWithN(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.MultiMintWithN(&_Subterfuge.TransactOpts, tokenIds)
}

// MultiMintWithN is a paid mutator transaction binding the contract method 0x47febae8.
//
// Solidity: function multiMintWithN(uint256[] tokenIds) payable returns()
func (_Subterfuge *SubterfugeTransactorSession) MultiMintWithN(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.MultiMintWithN(&_Subterfuge.TransactOpts, tokenIds)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subterfuge *SubterfugeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subterfuge *SubterfugeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subterfuge.Contract.RenounceOwnership(&_Subterfuge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Subterfuge *SubterfugeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Subterfuge.Contract.RenounceOwnership(&_Subterfuge.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Subterfuge *SubterfugeTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Subterfuge *SubterfugeSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.SafeTransferFrom(&_Subterfuge.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Subterfuge *SubterfugeTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.SafeTransferFrom(&_Subterfuge.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Subterfuge *SubterfugeTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Subterfuge *SubterfugeSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Subterfuge.Contract.SafeTransferFrom0(&_Subterfuge.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Subterfuge *SubterfugeTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Subterfuge.Contract.SafeTransferFrom0(&_Subterfuge.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Subterfuge *SubterfugeTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Subterfuge *SubterfugeSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Subterfuge.Contract.SetApprovalForAll(&_Subterfuge.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Subterfuge *SubterfugeTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Subterfuge.Contract.SetApprovalForAll(&_Subterfuge.TransactOpts, operator, approved)
}

// SetBurnerAddress is a paid mutator transaction binding the contract method 0x9df806d6.
//
// Solidity: function setBurnerAddress(address _burnerAddress) returns()
func (_Subterfuge *SubterfugeTransactor) SetBurnerAddress(opts *bind.TransactOpts, _burnerAddress common.Address) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "setBurnerAddress", _burnerAddress)
}

// SetBurnerAddress is a paid mutator transaction binding the contract method 0x9df806d6.
//
// Solidity: function setBurnerAddress(address _burnerAddress) returns()
func (_Subterfuge *SubterfugeSession) SetBurnerAddress(_burnerAddress common.Address) (*types.Transaction, error) {
	return _Subterfuge.Contract.SetBurnerAddress(&_Subterfuge.TransactOpts, _burnerAddress)
}

// SetBurnerAddress is a paid mutator transaction binding the contract method 0x9df806d6.
//
// Solidity: function setBurnerAddress(address _burnerAddress) returns()
func (_Subterfuge *SubterfugeTransactorSession) SetBurnerAddress(_burnerAddress common.Address) (*types.Transaction, error) {
	return _Subterfuge.Contract.SetBurnerAddress(&_Subterfuge.TransactOpts, _burnerAddress)
}

// SetPublicSale is a paid mutator transaction binding the contract method 0x5aca1bb6.
//
// Solidity: function setPublicSale(bool _publicSale) returns()
func (_Subterfuge *SubterfugeTransactor) SetPublicSale(opts *bind.TransactOpts, _publicSale bool) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "setPublicSale", _publicSale)
}

// SetPublicSale is a paid mutator transaction binding the contract method 0x5aca1bb6.
//
// Solidity: function setPublicSale(bool _publicSale) returns()
func (_Subterfuge *SubterfugeSession) SetPublicSale(_publicSale bool) (*types.Transaction, error) {
	return _Subterfuge.Contract.SetPublicSale(&_Subterfuge.TransactOpts, _publicSale)
}

// SetPublicSale is a paid mutator transaction binding the contract method 0x5aca1bb6.
//
// Solidity: function setPublicSale(bool _publicSale) returns()
func (_Subterfuge *SubterfugeTransactorSession) SetPublicSale(_publicSale bool) (*types.Transaction, error) {
	return _Subterfuge.Contract.SetPublicSale(&_Subterfuge.TransactOpts, _publicSale)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Subterfuge *SubterfugeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Subterfuge *SubterfugeSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.TransferFrom(&_Subterfuge.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Subterfuge *SubterfugeTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Subterfuge.Contract.TransferFrom(&_Subterfuge.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subterfuge *SubterfugeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subterfuge *SubterfugeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subterfuge.Contract.TransferOwnership(&_Subterfuge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Subterfuge *SubterfugeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Subterfuge.Contract.TransferOwnership(&_Subterfuge.TransactOpts, newOwner)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Subterfuge *SubterfugeTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subterfuge.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Subterfuge *SubterfugeSession) WithdrawAll() (*types.Transaction, error) {
	return _Subterfuge.Contract.WithdrawAll(&_Subterfuge.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Subterfuge *SubterfugeTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Subterfuge.Contract.WithdrawAll(&_Subterfuge.TransactOpts)
}

// SubterfugeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Subterfuge contract.
type SubterfugeApprovalIterator struct {
	Event *SubterfugeApproval // Event containing the contract specifics and raw log

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
func (it *SubterfugeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubterfugeApproval)
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
		it.Event = new(SubterfugeApproval)
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
func (it *SubterfugeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubterfugeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubterfugeApproval represents a Approval event raised by the Subterfuge contract.
type SubterfugeApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Subterfuge *SubterfugeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SubterfugeApprovalIterator, error) {

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

	logs, sub, err := _Subterfuge.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SubterfugeApprovalIterator{contract: _Subterfuge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Subterfuge *SubterfugeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SubterfugeApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Subterfuge.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubterfugeApproval)
				if err := _Subterfuge.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Subterfuge *SubterfugeFilterer) ParseApproval(log types.Log) (*SubterfugeApproval, error) {
	event := new(SubterfugeApproval)
	if err := _Subterfuge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubterfugeApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Subterfuge contract.
type SubterfugeApprovalForAllIterator struct {
	Event *SubterfugeApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SubterfugeApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubterfugeApprovalForAll)
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
		it.Event = new(SubterfugeApprovalForAll)
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
func (it *SubterfugeApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubterfugeApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubterfugeApprovalForAll represents a ApprovalForAll event raised by the Subterfuge contract.
type SubterfugeApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Subterfuge *SubterfugeFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SubterfugeApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Subterfuge.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SubterfugeApprovalForAllIterator{contract: _Subterfuge.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Subterfuge *SubterfugeFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SubterfugeApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Subterfuge.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubterfugeApprovalForAll)
				if err := _Subterfuge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Subterfuge *SubterfugeFilterer) ParseApprovalForAll(log types.Log) (*SubterfugeApprovalForAll, error) {
	event := new(SubterfugeApprovalForAll)
	if err := _Subterfuge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubterfugeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Subterfuge contract.
type SubterfugeOwnershipTransferredIterator struct {
	Event *SubterfugeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SubterfugeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubterfugeOwnershipTransferred)
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
		it.Event = new(SubterfugeOwnershipTransferred)
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
func (it *SubterfugeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubterfugeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubterfugeOwnershipTransferred represents a OwnershipTransferred event raised by the Subterfuge contract.
type SubterfugeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subterfuge *SubterfugeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SubterfugeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subterfuge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SubterfugeOwnershipTransferredIterator{contract: _Subterfuge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Subterfuge *SubterfugeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SubterfugeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Subterfuge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubterfugeOwnershipTransferred)
				if err := _Subterfuge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Subterfuge *SubterfugeFilterer) ParseOwnershipTransferred(log types.Log) (*SubterfugeOwnershipTransferred, error) {
	event := new(SubterfugeOwnershipTransferred)
	if err := _Subterfuge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubterfugeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Subterfuge contract.
type SubterfugeTransferIterator struct {
	Event *SubterfugeTransfer // Event containing the contract specifics and raw log

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
func (it *SubterfugeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubterfugeTransfer)
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
		it.Event = new(SubterfugeTransfer)
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
func (it *SubterfugeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubterfugeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubterfugeTransfer represents a Transfer event raised by the Subterfuge contract.
type SubterfugeTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Subterfuge *SubterfugeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SubterfugeTransferIterator, error) {

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

	logs, sub, err := _Subterfuge.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SubterfugeTransferIterator{contract: _Subterfuge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Subterfuge *SubterfugeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SubterfugeTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Subterfuge.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubterfugeTransfer)
				if err := _Subterfuge.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Subterfuge *SubterfugeFilterer) ParseTransfer(log types.Log) (*SubterfugeTransfer, error) {
	event := new(SubterfugeTransfer)
	if err := _Subterfuge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
