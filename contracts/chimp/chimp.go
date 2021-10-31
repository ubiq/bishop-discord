// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chimp

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

// CHIMPImageData is an auto generated low-level Go binding around an user-defined struct.
type CHIMPImageData struct {
	PixelChunks [2]*big.Int
	Colors      [4]uint8
	Author      common.Address
}

// ChimpABI is the input ABI used to generate the binding from.
const ChimpABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_burnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DIMENSION_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PALETTE_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PIXEL_CHUNKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnerAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"imageDataForToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pixelChunks\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint8[4]\",\"name\":\"colors\",\"type\":\"uint8[4]\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"}],\"internalType\":\"structCHIMP.ImageData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"pixelChunks\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint8[4]\",\"name\":\"colors\",\"type\":\"uint8[4]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintingActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"palette\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_burnerAddress\",\"type\":\"address\"}],\"name\":\"setBurnerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintPrice\",\"type\":\"uint256\"}],\"name\":\"setMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenSVG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Chimp is an auto generated Go binding around an Ethereum contract.
type Chimp struct {
	ChimpCaller     // Read-only binding to the contract
	ChimpTransactor // Write-only binding to the contract
	ChimpFilterer   // Log filterer for contract events
}

// ChimpCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChimpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChimpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChimpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChimpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChimpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChimpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChimpSession struct {
	Contract     *Chimp            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChimpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChimpCallerSession struct {
	Contract *ChimpCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ChimpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChimpTransactorSession struct {
	Contract     *ChimpTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChimpRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChimpRaw struct {
	Contract *Chimp // Generic contract binding to access the raw methods on
}

// ChimpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChimpCallerRaw struct {
	Contract *ChimpCaller // Generic read-only contract binding to access the raw methods on
}

// ChimpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChimpTransactorRaw struct {
	Contract *ChimpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChimp creates a new instance of Chimp, bound to a specific deployed contract.
func NewChimp(address common.Address, backend bind.ContractBackend) (*Chimp, error) {
	contract, err := bindChimp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chimp{ChimpCaller: ChimpCaller{contract: contract}, ChimpTransactor: ChimpTransactor{contract: contract}, ChimpFilterer: ChimpFilterer{contract: contract}}, nil
}

// NewChimpCaller creates a new read-only instance of Chimp, bound to a specific deployed contract.
func NewChimpCaller(address common.Address, caller bind.ContractCaller) (*ChimpCaller, error) {
	contract, err := bindChimp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChimpCaller{contract: contract}, nil
}

// NewChimpTransactor creates a new write-only instance of Chimp, bound to a specific deployed contract.
func NewChimpTransactor(address common.Address, transactor bind.ContractTransactor) (*ChimpTransactor, error) {
	contract, err := bindChimp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChimpTransactor{contract: contract}, nil
}

// NewChimpFilterer creates a new log filterer instance of Chimp, bound to a specific deployed contract.
func NewChimpFilterer(address common.Address, filterer bind.ContractFilterer) (*ChimpFilterer, error) {
	contract, err := bindChimp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChimpFilterer{contract: contract}, nil
}

// bindChimp binds a generic wrapper to an already deployed contract.
func bindChimp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChimpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chimp *ChimpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chimp.Contract.ChimpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chimp *ChimpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chimp.Contract.ChimpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chimp *ChimpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chimp.Contract.ChimpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chimp *ChimpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chimp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chimp *ChimpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chimp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chimp *ChimpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chimp.Contract.contract.Transact(opts, method, params...)
}

// DIMENSIONSIZE is a free data retrieval call binding the contract method 0x6683363e.
//
// Solidity: function DIMENSION_SIZE() view returns(uint256)
func (_Chimp *ChimpCaller) DIMENSIONSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "DIMENSION_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DIMENSIONSIZE is a free data retrieval call binding the contract method 0x6683363e.
//
// Solidity: function DIMENSION_SIZE() view returns(uint256)
func (_Chimp *ChimpSession) DIMENSIONSIZE() (*big.Int, error) {
	return _Chimp.Contract.DIMENSIONSIZE(&_Chimp.CallOpts)
}

// DIMENSIONSIZE is a free data retrieval call binding the contract method 0x6683363e.
//
// Solidity: function DIMENSION_SIZE() view returns(uint256)
func (_Chimp *ChimpCallerSession) DIMENSIONSIZE() (*big.Int, error) {
	return _Chimp.Contract.DIMENSIONSIZE(&_Chimp.CallOpts)
}

// PALETTESIZE is a free data retrieval call binding the contract method 0x70b1e5b9.
//
// Solidity: function PALETTE_SIZE() view returns(uint256)
func (_Chimp *ChimpCaller) PALETTESIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "PALETTE_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PALETTESIZE is a free data retrieval call binding the contract method 0x70b1e5b9.
//
// Solidity: function PALETTE_SIZE() view returns(uint256)
func (_Chimp *ChimpSession) PALETTESIZE() (*big.Int, error) {
	return _Chimp.Contract.PALETTESIZE(&_Chimp.CallOpts)
}

// PALETTESIZE is a free data retrieval call binding the contract method 0x70b1e5b9.
//
// Solidity: function PALETTE_SIZE() view returns(uint256)
func (_Chimp *ChimpCallerSession) PALETTESIZE() (*big.Int, error) {
	return _Chimp.Contract.PALETTESIZE(&_Chimp.CallOpts)
}

// PIXELCHUNKS is a free data retrieval call binding the contract method 0x9b2ee47f.
//
// Solidity: function PIXEL_CHUNKS() view returns(uint256)
func (_Chimp *ChimpCaller) PIXELCHUNKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "PIXEL_CHUNKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PIXELCHUNKS is a free data retrieval call binding the contract method 0x9b2ee47f.
//
// Solidity: function PIXEL_CHUNKS() view returns(uint256)
func (_Chimp *ChimpSession) PIXELCHUNKS() (*big.Int, error) {
	return _Chimp.Contract.PIXELCHUNKS(&_Chimp.CallOpts)
}

// PIXELCHUNKS is a free data retrieval call binding the contract method 0x9b2ee47f.
//
// Solidity: function PIXEL_CHUNKS() view returns(uint256)
func (_Chimp *ChimpCallerSession) PIXELCHUNKS() (*big.Int, error) {
	return _Chimp.Contract.PIXELCHUNKS(&_Chimp.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Chimp *ChimpCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Chimp *ChimpSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Chimp.Contract.BalanceOf(&_Chimp.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Chimp *ChimpCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Chimp.Contract.BalanceOf(&_Chimp.CallOpts, owner)
}

// BurnerAddress is a free data retrieval call binding the contract method 0xe6293e23.
//
// Solidity: function burnerAddress() view returns(address)
func (_Chimp *ChimpCaller) BurnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "burnerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnerAddress is a free data retrieval call binding the contract method 0xe6293e23.
//
// Solidity: function burnerAddress() view returns(address)
func (_Chimp *ChimpSession) BurnerAddress() (common.Address, error) {
	return _Chimp.Contract.BurnerAddress(&_Chimp.CallOpts)
}

// BurnerAddress is a free data retrieval call binding the contract method 0xe6293e23.
//
// Solidity: function burnerAddress() view returns(address)
func (_Chimp *ChimpCallerSession) BurnerAddress() (common.Address, error) {
	return _Chimp.Contract.BurnerAddress(&_Chimp.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Chimp *ChimpCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Chimp *ChimpSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Chimp.Contract.GetApproved(&_Chimp.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Chimp *ChimpCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Chimp.Contract.GetApproved(&_Chimp.CallOpts, tokenId)
}

// ImageDataForToken is a free data retrieval call binding the contract method 0x0dc65adb.
//
// Solidity: function imageDataForToken(uint256 tokenId) view returns((uint256[2],uint8[4],address))
func (_Chimp *ChimpCaller) ImageDataForToken(opts *bind.CallOpts, tokenId *big.Int) (CHIMPImageData, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "imageDataForToken", tokenId)

	if err != nil {
		return *new(CHIMPImageData), err
	}

	out0 := *abi.ConvertType(out[0], new(CHIMPImageData)).(*CHIMPImageData)

	return out0, err

}

// ImageDataForToken is a free data retrieval call binding the contract method 0x0dc65adb.
//
// Solidity: function imageDataForToken(uint256 tokenId) view returns((uint256[2],uint8[4],address))
func (_Chimp *ChimpSession) ImageDataForToken(tokenId *big.Int) (CHIMPImageData, error) {
	return _Chimp.Contract.ImageDataForToken(&_Chimp.CallOpts, tokenId)
}

// ImageDataForToken is a free data retrieval call binding the contract method 0x0dc65adb.
//
// Solidity: function imageDataForToken(uint256 tokenId) view returns((uint256[2],uint8[4],address))
func (_Chimp *ChimpCallerSession) ImageDataForToken(tokenId *big.Int) (CHIMPImageData, error) {
	return _Chimp.Contract.ImageDataForToken(&_Chimp.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Chimp *ChimpCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Chimp *ChimpSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Chimp.Contract.IsApprovedForAll(&_Chimp.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Chimp *ChimpCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Chimp.Contract.IsApprovedForAll(&_Chimp.CallOpts, owner, operator)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Chimp *ChimpCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Chimp *ChimpSession) MintPrice() (*big.Int, error) {
	return _Chimp.Contract.MintPrice(&_Chimp.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Chimp *ChimpCallerSession) MintPrice() (*big.Int, error) {
	return _Chimp.Contract.MintPrice(&_Chimp.CallOpts)
}

// MintingActive is a free data retrieval call binding the contract method 0x31f9c919.
//
// Solidity: function mintingActive() view returns(bool)
func (_Chimp *ChimpCaller) MintingActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "mintingActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintingActive is a free data retrieval call binding the contract method 0x31f9c919.
//
// Solidity: function mintingActive() view returns(bool)
func (_Chimp *ChimpSession) MintingActive() (bool, error) {
	return _Chimp.Contract.MintingActive(&_Chimp.CallOpts)
}

// MintingActive is a free data retrieval call binding the contract method 0x31f9c919.
//
// Solidity: function mintingActive() view returns(bool)
func (_Chimp *ChimpCallerSession) MintingActive() (bool, error) {
	return _Chimp.Contract.MintingActive(&_Chimp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Chimp *ChimpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Chimp *ChimpSession) Name() (string, error) {
	return _Chimp.Contract.Name(&_Chimp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Chimp *ChimpCallerSession) Name() (string, error) {
	return _Chimp.Contract.Name(&_Chimp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Chimp *ChimpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Chimp *ChimpSession) Owner() (common.Address, error) {
	return _Chimp.Contract.Owner(&_Chimp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Chimp *ChimpCallerSession) Owner() (common.Address, error) {
	return _Chimp.Contract.Owner(&_Chimp.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Chimp *ChimpCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Chimp *ChimpSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Chimp.Contract.OwnerOf(&_Chimp.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Chimp *ChimpCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Chimp.Contract.OwnerOf(&_Chimp.CallOpts, tokenId)
}

// Palette is a free data retrieval call binding the contract method 0x9cb71ef8.
//
// Solidity: function palette(uint256 ) view returns(string)
func (_Chimp *ChimpCaller) Palette(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "palette", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Palette is a free data retrieval call binding the contract method 0x9cb71ef8.
//
// Solidity: function palette(uint256 ) view returns(string)
func (_Chimp *ChimpSession) Palette(arg0 *big.Int) (string, error) {
	return _Chimp.Contract.Palette(&_Chimp.CallOpts, arg0)
}

// Palette is a free data retrieval call binding the contract method 0x9cb71ef8.
//
// Solidity: function palette(uint256 ) view returns(string)
func (_Chimp *ChimpCallerSession) Palette(arg0 *big.Int) (string, error) {
	return _Chimp.Contract.Palette(&_Chimp.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Chimp *ChimpCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Chimp *ChimpSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Chimp.Contract.SupportsInterface(&_Chimp.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Chimp *ChimpCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Chimp.Contract.SupportsInterface(&_Chimp.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Chimp *ChimpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Chimp *ChimpSession) Symbol() (string, error) {
	return _Chimp.Contract.Symbol(&_Chimp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Chimp *ChimpCallerSession) Symbol() (string, error) {
	return _Chimp.Contract.Symbol(&_Chimp.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Chimp *ChimpCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Chimp *ChimpSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Chimp.Contract.TokenByIndex(&_Chimp.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Chimp *ChimpCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Chimp.Contract.TokenByIndex(&_Chimp.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Chimp *ChimpCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Chimp *ChimpSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Chimp.Contract.TokenOfOwnerByIndex(&_Chimp.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Chimp *ChimpCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Chimp.Contract.TokenOfOwnerByIndex(&_Chimp.CallOpts, owner, index)
}

// TokenSVG is a free data retrieval call binding the contract method 0x9bac5f7a.
//
// Solidity: function tokenSVG(uint256 tokenId) view returns(string)
func (_Chimp *ChimpCaller) TokenSVG(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "tokenSVG", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenSVG is a free data retrieval call binding the contract method 0x9bac5f7a.
//
// Solidity: function tokenSVG(uint256 tokenId) view returns(string)
func (_Chimp *ChimpSession) TokenSVG(tokenId *big.Int) (string, error) {
	return _Chimp.Contract.TokenSVG(&_Chimp.CallOpts, tokenId)
}

// TokenSVG is a free data retrieval call binding the contract method 0x9bac5f7a.
//
// Solidity: function tokenSVG(uint256 tokenId) view returns(string)
func (_Chimp *ChimpCallerSession) TokenSVG(tokenId *big.Int) (string, error) {
	return _Chimp.Contract.TokenSVG(&_Chimp.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Chimp *ChimpCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Chimp *ChimpSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Chimp.Contract.TokenURI(&_Chimp.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Chimp *ChimpCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Chimp.Contract.TokenURI(&_Chimp.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Chimp *ChimpCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chimp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Chimp *ChimpSession) TotalSupply() (*big.Int, error) {
	return _Chimp.Contract.TotalSupply(&_Chimp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Chimp *ChimpCallerSession) TotalSupply() (*big.Int, error) {
	return _Chimp.Contract.TotalSupply(&_Chimp.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Chimp *ChimpTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Chimp *ChimpSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Chimp.Contract.Approve(&_Chimp.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Chimp *ChimpTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Chimp.Contract.Approve(&_Chimp.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x8e782bbb.
//
// Solidity: function mint(uint256[2] pixelChunks, uint8[4] colors) payable returns()
func (_Chimp *ChimpTransactor) Mint(opts *bind.TransactOpts, pixelChunks [2]*big.Int, colors [4]uint8) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "mint", pixelChunks, colors)
}

// Mint is a paid mutator transaction binding the contract method 0x8e782bbb.
//
// Solidity: function mint(uint256[2] pixelChunks, uint8[4] colors) payable returns()
func (_Chimp *ChimpSession) Mint(pixelChunks [2]*big.Int, colors [4]uint8) (*types.Transaction, error) {
	return _Chimp.Contract.Mint(&_Chimp.TransactOpts, pixelChunks, colors)
}

// Mint is a paid mutator transaction binding the contract method 0x8e782bbb.
//
// Solidity: function mint(uint256[2] pixelChunks, uint8[4] colors) payable returns()
func (_Chimp *ChimpTransactorSession) Mint(pixelChunks [2]*big.Int, colors [4]uint8) (*types.Transaction, error) {
	return _Chimp.Contract.Mint(&_Chimp.TransactOpts, pixelChunks, colors)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Chimp *ChimpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Chimp *ChimpSession) RenounceOwnership() (*types.Transaction, error) {
	return _Chimp.Contract.RenounceOwnership(&_Chimp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Chimp *ChimpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Chimp.Contract.RenounceOwnership(&_Chimp.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Chimp *ChimpTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Chimp *ChimpSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Chimp.Contract.SafeTransferFrom(&_Chimp.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Chimp *ChimpTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Chimp.Contract.SafeTransferFrom(&_Chimp.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Chimp *ChimpTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Chimp *ChimpSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Chimp.Contract.SafeTransferFrom0(&_Chimp.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Chimp *ChimpTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Chimp.Contract.SafeTransferFrom0(&_Chimp.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Chimp *ChimpTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Chimp *ChimpSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Chimp.Contract.SetApprovalForAll(&_Chimp.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Chimp *ChimpTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Chimp.Contract.SetApprovalForAll(&_Chimp.TransactOpts, operator, approved)
}

// SetBurnerAddress is a paid mutator transaction binding the contract method 0x9df806d6.
//
// Solidity: function setBurnerAddress(address _burnerAddress) returns()
func (_Chimp *ChimpTransactor) SetBurnerAddress(opts *bind.TransactOpts, _burnerAddress common.Address) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "setBurnerAddress", _burnerAddress)
}

// SetBurnerAddress is a paid mutator transaction binding the contract method 0x9df806d6.
//
// Solidity: function setBurnerAddress(address _burnerAddress) returns()
func (_Chimp *ChimpSession) SetBurnerAddress(_burnerAddress common.Address) (*types.Transaction, error) {
	return _Chimp.Contract.SetBurnerAddress(&_Chimp.TransactOpts, _burnerAddress)
}

// SetBurnerAddress is a paid mutator transaction binding the contract method 0x9df806d6.
//
// Solidity: function setBurnerAddress(address _burnerAddress) returns()
func (_Chimp *ChimpTransactorSession) SetBurnerAddress(_burnerAddress common.Address) (*types.Transaction, error) {
	return _Chimp.Contract.SetBurnerAddress(&_Chimp.TransactOpts, _burnerAddress)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_Chimp *ChimpTransactor) SetMintPrice(opts *bind.TransactOpts, _mintPrice *big.Int) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "setMintPrice", _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_Chimp *ChimpSession) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _Chimp.Contract.SetMintPrice(&_Chimp.TransactOpts, _mintPrice)
}

// SetMintPrice is a paid mutator transaction binding the contract method 0xf4a0a528.
//
// Solidity: function setMintPrice(uint256 _mintPrice) returns()
func (_Chimp *ChimpTransactorSession) SetMintPrice(_mintPrice *big.Int) (*types.Transaction, error) {
	return _Chimp.Contract.SetMintPrice(&_Chimp.TransactOpts, _mintPrice)
}

// ToggleActive is a paid mutator transaction binding the contract method 0x29c68dc1.
//
// Solidity: function toggleActive() returns()
func (_Chimp *ChimpTransactor) ToggleActive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "toggleActive")
}

// ToggleActive is a paid mutator transaction binding the contract method 0x29c68dc1.
//
// Solidity: function toggleActive() returns()
func (_Chimp *ChimpSession) ToggleActive() (*types.Transaction, error) {
	return _Chimp.Contract.ToggleActive(&_Chimp.TransactOpts)
}

// ToggleActive is a paid mutator transaction binding the contract method 0x29c68dc1.
//
// Solidity: function toggleActive() returns()
func (_Chimp *ChimpTransactorSession) ToggleActive() (*types.Transaction, error) {
	return _Chimp.Contract.ToggleActive(&_Chimp.TransactOpts)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Chimp *ChimpTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Chimp *ChimpSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Chimp.Contract.TransferFrom(&_Chimp.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Chimp *ChimpTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Chimp.Contract.TransferFrom(&_Chimp.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Chimp *ChimpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Chimp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Chimp *ChimpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Chimp.Contract.TransferOwnership(&_Chimp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Chimp *ChimpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Chimp.Contract.TransferOwnership(&_Chimp.TransactOpts, newOwner)
}

// ChimpApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Chimp contract.
type ChimpApprovalIterator struct {
	Event *ChimpApproval // Event containing the contract specifics and raw log

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
func (it *ChimpApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChimpApproval)
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
		it.Event = new(ChimpApproval)
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
func (it *ChimpApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChimpApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChimpApproval represents a Approval event raised by the Chimp contract.
type ChimpApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Chimp *ChimpFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ChimpApprovalIterator, error) {

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

	logs, sub, err := _Chimp.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ChimpApprovalIterator{contract: _Chimp.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Chimp *ChimpFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ChimpApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Chimp.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChimpApproval)
				if err := _Chimp.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Chimp *ChimpFilterer) ParseApproval(log types.Log) (*ChimpApproval, error) {
	event := new(ChimpApproval)
	if err := _Chimp.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChimpApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Chimp contract.
type ChimpApprovalForAllIterator struct {
	Event *ChimpApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ChimpApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChimpApprovalForAll)
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
		it.Event = new(ChimpApprovalForAll)
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
func (it *ChimpApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChimpApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChimpApprovalForAll represents a ApprovalForAll event raised by the Chimp contract.
type ChimpApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Chimp *ChimpFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ChimpApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Chimp.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ChimpApprovalForAllIterator{contract: _Chimp.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Chimp *ChimpFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ChimpApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Chimp.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChimpApprovalForAll)
				if err := _Chimp.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Chimp *ChimpFilterer) ParseApprovalForAll(log types.Log) (*ChimpApprovalForAll, error) {
	event := new(ChimpApprovalForAll)
	if err := _Chimp.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChimpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Chimp contract.
type ChimpOwnershipTransferredIterator struct {
	Event *ChimpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChimpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChimpOwnershipTransferred)
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
		it.Event = new(ChimpOwnershipTransferred)
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
func (it *ChimpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChimpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChimpOwnershipTransferred represents a OwnershipTransferred event raised by the Chimp contract.
type ChimpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Chimp *ChimpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChimpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Chimp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChimpOwnershipTransferredIterator{contract: _Chimp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Chimp *ChimpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChimpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Chimp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChimpOwnershipTransferred)
				if err := _Chimp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Chimp *ChimpFilterer) ParseOwnershipTransferred(log types.Log) (*ChimpOwnershipTransferred, error) {
	event := new(ChimpOwnershipTransferred)
	if err := _Chimp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChimpTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Chimp contract.
type ChimpTransferIterator struct {
	Event *ChimpTransfer // Event containing the contract specifics and raw log

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
func (it *ChimpTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChimpTransfer)
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
		it.Event = new(ChimpTransfer)
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
func (it *ChimpTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChimpTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChimpTransfer represents a Transfer event raised by the Chimp contract.
type ChimpTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Chimp *ChimpFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ChimpTransferIterator, error) {

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

	logs, sub, err := _Chimp.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ChimpTransferIterator{contract: _Chimp.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Chimp *ChimpFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ChimpTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Chimp.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChimpTransfer)
				if err := _Chimp.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Chimp *ChimpFilterer) ParseTransfer(log types.Log) (*ChimpTransfer, error) {
	event := new(ChimpTransfer)
	if err := _Chimp.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
