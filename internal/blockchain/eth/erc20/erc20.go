// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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
	_ = abi.ConvertType
)

// Erc20MetaData contains all meta data concerning the Erc20 contract.
var Erc20MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_upgradedAddress\",\"type\":\"address\"}],\"name\":\"deprecate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deprecated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradedAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maximumFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_totalSupplyasd\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBasisPoints\",\"type\":\"uint256\"},{\"name\":\"newMaxFee\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"basisPointsRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_UINT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blackListedUser\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"Deprecate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"feeBasisPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"Params\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]",
}

// Erc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20MetaData.ABI instead.
var Erc20ABI = Erc20MetaData.ABI

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct {
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
}

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct {
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct {
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct {
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct {
	Contract *Erc20 // Generic contract binding to access the raw methods on
}

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct {
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct {
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Caller{contract: contract}, nil
}

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Transactor{contract: contract}, nil
}

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Filterer{contract: contract}, nil
}

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Erc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transact(opts, method, params...)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Erc20 *Erc20Caller) MAXUINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "MAX_UINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Erc20 *Erc20Session) MAXUINT() (*big.Int, error) {
	return _Erc20.Contract.MAXUINT(&_Erc20.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Erc20 *Erc20CallerSession) MAXUINT() (*big.Int, error) {
	return _Erc20.Contract.MAXUINT(&_Erc20.CallOpts)
}

// TotalSupplyasd is a free data retrieval call binding the contract method 0xc0448cce.
//
// Solidity: function _totalSupplyasd() view returns(uint256)
func (_Erc20 *Erc20Caller) TotalSupplyasd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "_totalSupplyasd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyasd is a free data retrieval call binding the contract method 0xc0448cce.
//
// Solidity: function _totalSupplyasd() view returns(uint256)
func (_Erc20 *Erc20Session) TotalSupplyasd() (*big.Int, error) {
	return _Erc20.Contract.TotalSupplyasd(&_Erc20.CallOpts)
}

// TotalSupplyasd is a free data retrieval call binding the contract method 0xc0448cce.
//
// Solidity: function _totalSupplyasd() view returns(uint256)
func (_Erc20 *Erc20CallerSession) TotalSupplyasd() (*big.Int, error) {
	return _Erc20.Contract.TotalSupplyasd(&_Erc20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Erc20 *Erc20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Erc20 *Erc20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Erc20 *Erc20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Erc20 *Erc20Caller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "allowed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Erc20 *Erc20Session) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowed(&_Erc20.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowed(&_Erc20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Erc20 *Erc20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Erc20 *Erc20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Erc20 *Erc20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, who)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Erc20 *Erc20Caller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Erc20 *Erc20Session) Balances(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.Balances(&_Erc20.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Erc20 *Erc20CallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Erc20.Contract.Balances(&_Erc20.CallOpts, arg0)
}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() view returns(uint256)
func (_Erc20 *Erc20Caller) BasisPointsRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "basisPointsRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() view returns(uint256)
func (_Erc20 *Erc20Session) BasisPointsRate() (*big.Int, error) {
	return _Erc20.Contract.BasisPointsRate(&_Erc20.CallOpts)
}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() view returns(uint256)
func (_Erc20 *Erc20CallerSession) BasisPointsRate() (*big.Int, error) {
	return _Erc20.Contract.BasisPointsRate(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20 *Erc20Session) Decimals() (*big.Int, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20 *Erc20CallerSession) Decimals() (*big.Int, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_Erc20 *Erc20Caller) Deprecated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "deprecated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_Erc20 *Erc20Session) Deprecated() (bool, error) {
	return _Erc20.Contract.Deprecated(&_Erc20.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_Erc20 *Erc20CallerSession) Deprecated() (bool, error) {
	return _Erc20.Contract.Deprecated(&_Erc20.CallOpts)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_Erc20 *Erc20Caller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_Erc20 *Erc20Session) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _Erc20.Contract.GetBlackListStatus(&_Erc20.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_Erc20 *Erc20CallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _Erc20.Contract.GetBlackListStatus(&_Erc20.CallOpts, _maker)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Erc20 *Erc20Caller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Erc20 *Erc20Session) GetOwner() (common.Address, error) {
	return _Erc20.Contract.GetOwner(&_Erc20.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Erc20 *Erc20CallerSession) GetOwner() (common.Address, error) {
	return _Erc20.Contract.GetOwner(&_Erc20.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_Erc20 *Erc20Caller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_Erc20 *Erc20Session) IsBlackListed(arg0 common.Address) (bool, error) {
	return _Erc20.Contract.IsBlackListed(&_Erc20.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_Erc20 *Erc20CallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _Erc20.Contract.IsBlackListed(&_Erc20.CallOpts, arg0)
}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() view returns(uint256)
func (_Erc20 *Erc20Caller) MaximumFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "maximumFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() view returns(uint256)
func (_Erc20 *Erc20Session) MaximumFee() (*big.Int, error) {
	return _Erc20.Contract.MaximumFee(&_Erc20.CallOpts)
}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() view returns(uint256)
func (_Erc20 *Erc20CallerSession) MaximumFee() (*big.Int, error) {
	return _Erc20.Contract.MaximumFee(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Session) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20CallerSession) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20Session) Owner() (common.Address, error) {
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc20 *Erc20CallerSession) Owner() (common.Address, error) {
	return _Erc20.Contract.Owner(&_Erc20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc20 *Erc20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc20 *Erc20Session) Paused() (bool, error) {
	return _Erc20.Contract.Paused(&_Erc20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Erc20 *Erc20CallerSession) Paused() (bool, error) {
	return _Erc20.Contract.Paused(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Session) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20CallerSession) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Session) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_Erc20 *Erc20Caller) UpgradedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "upgradedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_Erc20 *Erc20Session) UpgradedAddress() (common.Address, error) {
	return _Erc20.Contract.UpgradedAddress(&_Erc20.CallOpts)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_Erc20 *Erc20CallerSession) UpgradedAddress() (common.Address, error) {
	return _Erc20.Contract.UpgradedAddress(&_Erc20.CallOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_Erc20 *Erc20Transactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_Erc20 *Erc20Session) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.AddBlackList(&_Erc20.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_Erc20 *Erc20TransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.AddBlackList(&_Erc20.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_Erc20 *Erc20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_Erc20 *Erc20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, _spender, _value)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_Erc20 *Erc20Transactor) Deprecate(opts *bind.TransactOpts, _upgradedAddress common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "deprecate", _upgradedAddress)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_Erc20 *Erc20Session) Deprecate(_upgradedAddress common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.Deprecate(&_Erc20.TransactOpts, _upgradedAddress)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_Erc20 *Erc20TransactorSession) Deprecate(_upgradedAddress common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.Deprecate(&_Erc20.TransactOpts, _upgradedAddress)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_Erc20 *Erc20Transactor) DestroyBlackFunds(opts *bind.TransactOpts, _blackListedUser common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "destroyBlackFunds", _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_Erc20 *Erc20Session) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DestroyBlackFunds(&_Erc20.TransactOpts, _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_Erc20 *Erc20TransactorSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DestroyBlackFunds(&_Erc20.TransactOpts, _blackListedUser)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_Erc20 *Erc20Transactor) Issue(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "issue", amount)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_Erc20 *Erc20Session) Issue(amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Issue(&_Erc20.TransactOpts, amount)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) Issue(amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Issue(&_Erc20.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc20 *Erc20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc20 *Erc20Session) Pause() (*types.Transaction, error) {
	return _Erc20.Contract.Pause(&_Erc20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc20 *Erc20TransactorSession) Pause() (*types.Transaction, error) {
	return _Erc20.Contract.Pause(&_Erc20.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_Erc20 *Erc20Transactor) Redeem(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "redeem", amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_Erc20 *Erc20Session) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Redeem(&_Erc20.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_Erc20 *Erc20TransactorSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Redeem(&_Erc20.TransactOpts, amount)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_Erc20 *Erc20Transactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_Erc20 *Erc20Session) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.RemoveBlackList(&_Erc20.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_Erc20 *Erc20TransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.RemoveBlackList(&_Erc20.TransactOpts, _clearedUser)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_Erc20 *Erc20Transactor) SetParams(opts *bind.TransactOpts, newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "setParams", newBasisPoints, newMaxFee)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_Erc20 *Erc20Session) SetParams(newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SetParams(&_Erc20.TransactOpts, newBasisPoints, newMaxFee)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_Erc20 *Erc20TransactorSession) SetParams(newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.SetParams(&_Erc20.TransactOpts, newBasisPoints, newMaxFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_Erc20 *Erc20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_Erc20 *Erc20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_Erc20 *Erc20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_Erc20 *Erc20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Erc20 *Erc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.TransferOwnership(&_Erc20.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Erc20 *Erc20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Erc20 *Erc20Session) Unpause() (*types.Transaction, error) {
	return _Erc20.Contract.Unpause(&_Erc20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Erc20 *Erc20TransactorSession) Unpause() (*types.Transaction, error) {
	return _Erc20.Contract.Unpause(&_Erc20.TransactOpts)
}

// Erc20AddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the Erc20 contract.
type Erc20AddedBlackListIterator struct {
	Event *Erc20AddedBlackList // Event containing the contract specifics and raw log

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
func (it *Erc20AddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20AddedBlackList)
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
		it.Event = new(Erc20AddedBlackList)
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
func (it *Erc20AddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20AddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20AddedBlackList represents a AddedBlackList event raised by the Erc20 contract.
type Erc20AddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_Erc20 *Erc20Filterer) FilterAddedBlackList(opts *bind.FilterOpts) (*Erc20AddedBlackListIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &Erc20AddedBlackListIterator{contract: _Erc20.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_Erc20 *Erc20Filterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *Erc20AddedBlackList) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20AddedBlackList)
				if err := _Erc20.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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

// ParseAddedBlackList is a log parse operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_Erc20 *Erc20Filterer) ParseAddedBlackList(log types.Log) (*Erc20AddedBlackList, error) {
	event := new(Erc20AddedBlackList)
	if err := _Erc20.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20 contract.
type Erc20ApprovalIterator struct {
	Event *Erc20Approval // Event containing the contract specifics and raw log

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
func (it *Erc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Approval)
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
		it.Event = new(Erc20Approval)
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
func (it *Erc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Approval represents a Approval event raised by the Erc20 contract.
type Erc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20ApprovalIterator{contract: _Erc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Approval)
				if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseApproval(log types.Log) (*Erc20Approval, error) {
	event := new(Erc20Approval)
	if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20DeprecateIterator is returned from FilterDeprecate and is used to iterate over the raw logs and unpacked data for Deprecate events raised by the Erc20 contract.
type Erc20DeprecateIterator struct {
	Event *Erc20Deprecate // Event containing the contract specifics and raw log

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
func (it *Erc20DeprecateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Deprecate)
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
		it.Event = new(Erc20Deprecate)
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
func (it *Erc20DeprecateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20DeprecateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Deprecate represents a Deprecate event raised by the Erc20 contract.
type Erc20Deprecate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeprecate is a free log retrieval operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_Erc20 *Erc20Filterer) FilterDeprecate(opts *bind.FilterOpts) (*Erc20DeprecateIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Deprecate")
	if err != nil {
		return nil, err
	}
	return &Erc20DeprecateIterator{contract: _Erc20.contract, event: "Deprecate", logs: logs, sub: sub}, nil
}

// WatchDeprecate is a free log subscription operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_Erc20 *Erc20Filterer) WatchDeprecate(opts *bind.WatchOpts, sink chan<- *Erc20Deprecate) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Deprecate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Deprecate)
				if err := _Erc20.contract.UnpackLog(event, "Deprecate", log); err != nil {
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

// ParseDeprecate is a log parse operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_Erc20 *Erc20Filterer) ParseDeprecate(log types.Log) (*Erc20Deprecate, error) {
	event := new(Erc20Deprecate)
	if err := _Erc20.contract.UnpackLog(event, "Deprecate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20DestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the Erc20 contract.
type Erc20DestroyedBlackFundsIterator struct {
	Event *Erc20DestroyedBlackFunds // Event containing the contract specifics and raw log

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
func (it *Erc20DestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20DestroyedBlackFunds)
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
		it.Event = new(Erc20DestroyedBlackFunds)
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
func (it *Erc20DestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20DestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20DestroyedBlackFunds represents a DestroyedBlackFunds event raised by the Erc20 contract.
type Erc20DestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_Erc20 *Erc20Filterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*Erc20DestroyedBlackFundsIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &Erc20DestroyedBlackFundsIterator{contract: _Erc20.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_Erc20 *Erc20Filterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *Erc20DestroyedBlackFunds) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20DestroyedBlackFunds)
				if err := _Erc20.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
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

// ParseDestroyedBlackFunds is a log parse operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_Erc20 *Erc20Filterer) ParseDestroyedBlackFunds(log types.Log) (*Erc20DestroyedBlackFunds, error) {
	event := new(Erc20DestroyedBlackFunds)
	if err := _Erc20.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20IssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the Erc20 contract.
type Erc20IssueIterator struct {
	Event *Erc20Issue // Event containing the contract specifics and raw log

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
func (it *Erc20IssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Issue)
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
		it.Event = new(Erc20Issue)
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
func (it *Erc20IssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20IssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Issue represents a Issue event raised by the Erc20 contract.
type Erc20Issue struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_Erc20 *Erc20Filterer) FilterIssue(opts *bind.FilterOpts) (*Erc20IssueIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return &Erc20IssueIterator{contract: _Erc20.contract, event: "Issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_Erc20 *Erc20Filterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *Erc20Issue) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Issue)
				if err := _Erc20.contract.UnpackLog(event, "Issue", log); err != nil {
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

// ParseIssue is a log parse operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_Erc20 *Erc20Filterer) ParseIssue(log types.Log) (*Erc20Issue, error) {
	event := new(Erc20Issue)
	if err := _Erc20.contract.UnpackLog(event, "Issue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20ParamsIterator is returned from FilterParams and is used to iterate over the raw logs and unpacked data for Params events raised by the Erc20 contract.
type Erc20ParamsIterator struct {
	Event *Erc20Params // Event containing the contract specifics and raw log

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
func (it *Erc20ParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Params)
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
		it.Event = new(Erc20Params)
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
func (it *Erc20ParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20ParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Params represents a Params event raised by the Erc20 contract.
type Erc20Params struct {
	FeeBasisPoints *big.Int
	MaxFee         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterParams is a free log retrieval operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_Erc20 *Erc20Filterer) FilterParams(opts *bind.FilterOpts) (*Erc20ParamsIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Params")
	if err != nil {
		return nil, err
	}
	return &Erc20ParamsIterator{contract: _Erc20.contract, event: "Params", logs: logs, sub: sub}, nil
}

// WatchParams is a free log subscription operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_Erc20 *Erc20Filterer) WatchParams(opts *bind.WatchOpts, sink chan<- *Erc20Params) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Params")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Params)
				if err := _Erc20.contract.UnpackLog(event, "Params", log); err != nil {
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

// ParseParams is a log parse operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_Erc20 *Erc20Filterer) ParseParams(log types.Log) (*Erc20Params, error) {
	event := new(Erc20Params)
	if err := _Erc20.contract.UnpackLog(event, "Params", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20PauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Erc20 contract.
type Erc20PauseIterator struct {
	Event *Erc20Pause // Event containing the contract specifics and raw log

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
func (it *Erc20PauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Pause)
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
		it.Event = new(Erc20Pause)
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
func (it *Erc20PauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20PauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Pause represents a Pause event raised by the Erc20 contract.
type Erc20Pause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Erc20 *Erc20Filterer) FilterPause(opts *bind.FilterOpts) (*Erc20PauseIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &Erc20PauseIterator{contract: _Erc20.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Erc20 *Erc20Filterer) WatchPause(opts *bind.WatchOpts, sink chan<- *Erc20Pause) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Pause)
				if err := _Erc20.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Erc20 *Erc20Filterer) ParsePause(log types.Log) (*Erc20Pause, error) {
	event := new(Erc20Pause)
	if err := _Erc20.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20RedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Erc20 contract.
type Erc20RedeemIterator struct {
	Event *Erc20Redeem // Event containing the contract specifics and raw log

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
func (it *Erc20RedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Redeem)
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
		it.Event = new(Erc20Redeem)
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
func (it *Erc20RedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20RedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Redeem represents a Redeem event raised by the Erc20 contract.
type Erc20Redeem struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_Erc20 *Erc20Filterer) FilterRedeem(opts *bind.FilterOpts) (*Erc20RedeemIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &Erc20RedeemIterator{contract: _Erc20.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_Erc20 *Erc20Filterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *Erc20Redeem) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Redeem)
				if err := _Erc20.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_Erc20 *Erc20Filterer) ParseRedeem(log types.Log) (*Erc20Redeem, error) {
	event := new(Erc20Redeem)
	if err := _Erc20.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20RemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the Erc20 contract.
type Erc20RemovedBlackListIterator struct {
	Event *Erc20RemovedBlackList // Event containing the contract specifics and raw log

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
func (it *Erc20RemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20RemovedBlackList)
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
		it.Event = new(Erc20RemovedBlackList)
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
func (it *Erc20RemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20RemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20RemovedBlackList represents a RemovedBlackList event raised by the Erc20 contract.
type Erc20RemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_Erc20 *Erc20Filterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*Erc20RemovedBlackListIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &Erc20RemovedBlackListIterator{contract: _Erc20.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_Erc20 *Erc20Filterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *Erc20RemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20RemovedBlackList)
				if err := _Erc20.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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

// ParseRemovedBlackList is a log parse operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_Erc20 *Erc20Filterer) ParseRemovedBlackList(log types.Log) (*Erc20RemovedBlackList, error) {
	event := new(Erc20RemovedBlackList)
	if err := _Erc20.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20 contract.
type Erc20TransferIterator struct {
	Event *Erc20Transfer // Event containing the contract specifics and raw log

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
func (it *Erc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Transfer)
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
		it.Event = new(Erc20Transfer)
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
func (it *Erc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Transfer represents a Transfer event raised by the Erc20 contract.
type Erc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20TransferIterator{contract: _Erc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Transfer)
				if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseTransfer(log types.Log) (*Erc20Transfer, error) {
	event := new(Erc20Transfer)
	if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20UnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Erc20 contract.
type Erc20UnpauseIterator struct {
	Event *Erc20Unpause // Event containing the contract specifics and raw log

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
func (it *Erc20UnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Unpause)
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
		it.Event = new(Erc20Unpause)
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
func (it *Erc20UnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20UnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Unpause represents a Unpause event raised by the Erc20 contract.
type Erc20Unpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Erc20 *Erc20Filterer) FilterUnpause(opts *bind.FilterOpts) (*Erc20UnpauseIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &Erc20UnpauseIterator{contract: _Erc20.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Erc20 *Erc20Filterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *Erc20Unpause) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Unpause)
				if err := _Erc20.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Erc20 *Erc20Filterer) ParseUnpause(log types.Log) (*Erc20Unpause, error) {
	event := new(Erc20Unpause)
	if err := _Erc20.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
