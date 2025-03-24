// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sc

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

// MobyOLPMetaData contains all meta data concerning the MobyOLP contract.
var MobyOLPMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIOptionsAuthority\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"SetHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"inPrivateTransferMode\",\"type\":\"bool\"}],\"name\":\"SetInPrivateTransferMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"SetInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"SetMinter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"yieldTrackers\",\"type\":\"address[]\"}],\"name\":\"SetYieldTrackers\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNonStakingAccount\",\"type\":\"bool\"}],\"name\":\"UpdateNonStakingAccount\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"_authority\",\"type\":\"address\"}],\"name\":\"__AuthorityUtil_init__\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"_authority\",\"type\":\"address\"}],\"name\":\"__BaseToken_init__\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"_authority\",\"type\":\"address\"}],\"name\":\"__MintableBaseToken_init__\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addNonStakingAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateTransferMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"_authority\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonStakingAccounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonStakingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"recoverClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeNonStakingAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"_newAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateTransferMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateTransferMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"setInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_yieldTrackers\",\"type\":\"address[]\"}],\"name\":\"setYieldTrackers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"stakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"yieldTrackers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MobyOLPABI is the input ABI used to generate the binding from.
// Deprecated: Use MobyOLPMetaData.ABI instead.
var MobyOLPABI = MobyOLPMetaData.ABI

// MobyOLP is an auto generated Go binding around an Ethereum contract.
type MobyOLP struct {
	MobyOLPCaller     // Read-only binding to the contract
	MobyOLPTransactor // Write-only binding to the contract
	MobyOLPFilterer   // Log filterer for contract events
}

// MobyOLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type MobyOLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MobyOLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MobyOLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MobyOLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MobyOLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MobyOLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MobyOLPSession struct {
	Contract     *MobyOLP          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MobyOLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MobyOLPCallerSession struct {
	Contract *MobyOLPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MobyOLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MobyOLPTransactorSession struct {
	Contract     *MobyOLPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MobyOLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type MobyOLPRaw struct {
	Contract *MobyOLP // Generic contract binding to access the raw methods on
}

// MobyOLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MobyOLPCallerRaw struct {
	Contract *MobyOLPCaller // Generic read-only contract binding to access the raw methods on
}

// MobyOLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MobyOLPTransactorRaw struct {
	Contract *MobyOLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMobyOLP creates a new instance of MobyOLP, bound to a specific deployed contract.
func NewMobyOLP(address common.Address, backend bind.ContractBackend) (*MobyOLP, error) {
	contract, err := bindMobyOLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MobyOLP{MobyOLPCaller: MobyOLPCaller{contract: contract}, MobyOLPTransactor: MobyOLPTransactor{contract: contract}, MobyOLPFilterer: MobyOLPFilterer{contract: contract}}, nil
}

// NewMobyOLPCaller creates a new read-only instance of MobyOLP, bound to a specific deployed contract.
func NewMobyOLPCaller(address common.Address, caller bind.ContractCaller) (*MobyOLPCaller, error) {
	contract, err := bindMobyOLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MobyOLPCaller{contract: contract}, nil
}

// NewMobyOLPTransactor creates a new write-only instance of MobyOLP, bound to a specific deployed contract.
func NewMobyOLPTransactor(address common.Address, transactor bind.ContractTransactor) (*MobyOLPTransactor, error) {
	contract, err := bindMobyOLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MobyOLPTransactor{contract: contract}, nil
}

// NewMobyOLPFilterer creates a new log filterer instance of MobyOLP, bound to a specific deployed contract.
func NewMobyOLPFilterer(address common.Address, filterer bind.ContractFilterer) (*MobyOLPFilterer, error) {
	contract, err := bindMobyOLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MobyOLPFilterer{contract: contract}, nil
}

// bindMobyOLP binds a generic wrapper to an already deployed contract.
func bindMobyOLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MobyOLPMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MobyOLP *MobyOLPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MobyOLP.Contract.MobyOLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MobyOLP *MobyOLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MobyOLP.Contract.MobyOLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MobyOLP *MobyOLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MobyOLP.Contract.MobyOLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MobyOLP *MobyOLPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MobyOLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MobyOLP *MobyOLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MobyOLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MobyOLP *MobyOLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MobyOLP.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_MobyOLP *MobyOLPCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_MobyOLP *MobyOLPSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.Allowance(&_MobyOLP.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_MobyOLP *MobyOLPCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.Allowance(&_MobyOLP.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_MobyOLP *MobyOLPCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_MobyOLP *MobyOLPSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.Allowances(&_MobyOLP.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_MobyOLP *MobyOLPCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.Allowances(&_MobyOLP.CallOpts, arg0, arg1)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MobyOLP *MobyOLPCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MobyOLP *MobyOLPSession) Authority() (common.Address, error) {
	return _MobyOLP.Contract.Authority(&_MobyOLP.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MobyOLP *MobyOLPCallerSession) Authority() (common.Address, error) {
	return _MobyOLP.Contract.Authority(&_MobyOLP.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_MobyOLP *MobyOLPCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_MobyOLP *MobyOLPSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.BalanceOf(&_MobyOLP.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_MobyOLP *MobyOLPCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.BalanceOf(&_MobyOLP.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_MobyOLP *MobyOLPCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_MobyOLP *MobyOLPSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.Balances(&_MobyOLP.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_MobyOLP *MobyOLPCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.Balances(&_MobyOLP.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MobyOLP *MobyOLPCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MobyOLP *MobyOLPSession) Decimals() (uint8, error) {
	return _MobyOLP.Contract.Decimals(&_MobyOLP.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MobyOLP *MobyOLPCallerSession) Decimals() (uint8, error) {
	return _MobyOLP.Contract.Decimals(&_MobyOLP.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(string)
func (_MobyOLP *MobyOLPCaller) Id(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(string)
func (_MobyOLP *MobyOLPSession) Id() (string, error) {
	return _MobyOLP.Contract.Id(&_MobyOLP.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(string)
func (_MobyOLP *MobyOLPCallerSession) Id() (string, error) {
	return _MobyOLP.Contract.Id(&_MobyOLP.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_MobyOLP *MobyOLPCaller) InPrivateTransferMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "inPrivateTransferMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_MobyOLP *MobyOLPSession) InPrivateTransferMode() (bool, error) {
	return _MobyOLP.Contract.InPrivateTransferMode(&_MobyOLP.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_MobyOLP *MobyOLPCallerSession) InPrivateTransferMode() (bool, error) {
	return _MobyOLP.Contract.InPrivateTransferMode(&_MobyOLP.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_MobyOLP *MobyOLPCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_MobyOLP *MobyOLPSession) IsHandler(arg0 common.Address) (bool, error) {
	return _MobyOLP.Contract.IsHandler(&_MobyOLP.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_MobyOLP *MobyOLPCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _MobyOLP.Contract.IsHandler(&_MobyOLP.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_MobyOLP *MobyOLPCaller) IsMinter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "isMinter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_MobyOLP *MobyOLPSession) IsMinter(arg0 common.Address) (bool, error) {
	return _MobyOLP.Contract.IsMinter(&_MobyOLP.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_MobyOLP *MobyOLPCallerSession) IsMinter(arg0 common.Address) (bool, error) {
	return _MobyOLP.Contract.IsMinter(&_MobyOLP.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MobyOLP *MobyOLPCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MobyOLP *MobyOLPSession) Name() (string, error) {
	return _MobyOLP.Contract.Name(&_MobyOLP.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MobyOLP *MobyOLPCallerSession) Name() (string, error) {
	return _MobyOLP.Contract.Name(&_MobyOLP.CallOpts)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_MobyOLP *MobyOLPCaller) NonStakingAccounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "nonStakingAccounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_MobyOLP *MobyOLPSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _MobyOLP.Contract.NonStakingAccounts(&_MobyOLP.CallOpts, arg0)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_MobyOLP *MobyOLPCallerSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _MobyOLP.Contract.NonStakingAccounts(&_MobyOLP.CallOpts, arg0)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_MobyOLP *MobyOLPCaller) NonStakingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "nonStakingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_MobyOLP *MobyOLPSession) NonStakingSupply() (*big.Int, error) {
	return _MobyOLP.Contract.NonStakingSupply(&_MobyOLP.CallOpts)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_MobyOLP *MobyOLPCallerSession) NonStakingSupply() (*big.Int, error) {
	return _MobyOLP.Contract.NonStakingSupply(&_MobyOLP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MobyOLP *MobyOLPCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MobyOLP *MobyOLPSession) Owner() (common.Address, error) {
	return _MobyOLP.Contract.Owner(&_MobyOLP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MobyOLP *MobyOLPCallerSession) Owner() (common.Address, error) {
	return _MobyOLP.Contract.Owner(&_MobyOLP.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_MobyOLP *MobyOLPCaller) StakedBalance(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "stakedBalance", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_MobyOLP *MobyOLPSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.StakedBalance(&_MobyOLP.CallOpts, _account)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_MobyOLP *MobyOLPCallerSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _MobyOLP.Contract.StakedBalance(&_MobyOLP.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MobyOLP *MobyOLPCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MobyOLP *MobyOLPSession) Symbol() (string, error) {
	return _MobyOLP.Contract.Symbol(&_MobyOLP.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MobyOLP *MobyOLPCallerSession) Symbol() (string, error) {
	return _MobyOLP.Contract.Symbol(&_MobyOLP.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_MobyOLP *MobyOLPCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_MobyOLP *MobyOLPSession) TotalStaked() (*big.Int, error) {
	return _MobyOLP.Contract.TotalStaked(&_MobyOLP.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_MobyOLP *MobyOLPCallerSession) TotalStaked() (*big.Int, error) {
	return _MobyOLP.Contract.TotalStaked(&_MobyOLP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MobyOLP *MobyOLPCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MobyOLP *MobyOLPSession) TotalSupply() (*big.Int, error) {
	return _MobyOLP.Contract.TotalSupply(&_MobyOLP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MobyOLP *MobyOLPCallerSession) TotalSupply() (*big.Int, error) {
	return _MobyOLP.Contract.TotalSupply(&_MobyOLP.CallOpts)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_MobyOLP *MobyOLPCaller) YieldTrackers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MobyOLP.contract.Call(opts, &out, "yieldTrackers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_MobyOLP *MobyOLPSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _MobyOLP.Contract.YieldTrackers(&_MobyOLP.CallOpts, arg0)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_MobyOLP *MobyOLPCallerSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _MobyOLP.Contract.YieldTrackers(&_MobyOLP.CallOpts, arg0)
}

// AuthorityUtilInit is a paid mutator transaction binding the contract method 0xbdaf7884.
//
// Solidity: function __AuthorityUtil_init__(address _authority) returns()
func (_MobyOLP *MobyOLPTransactor) AuthorityUtilInit(opts *bind.TransactOpts, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "__AuthorityUtil_init__", _authority)
}

// AuthorityUtilInit is a paid mutator transaction binding the contract method 0xbdaf7884.
//
// Solidity: function __AuthorityUtil_init__(address _authority) returns()
func (_MobyOLP *MobyOLPSession) AuthorityUtilInit(_authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.AuthorityUtilInit(&_MobyOLP.TransactOpts, _authority)
}

// AuthorityUtilInit is a paid mutator transaction binding the contract method 0xbdaf7884.
//
// Solidity: function __AuthorityUtil_init__(address _authority) returns()
func (_MobyOLP *MobyOLPTransactorSession) AuthorityUtilInit(_authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.AuthorityUtilInit(&_MobyOLP.TransactOpts, _authority)
}

// BaseTokenInit is a paid mutator transaction binding the contract method 0xfcc76648.
//
// Solidity: function __BaseToken_init__(string _name, string _symbol, uint256 _initialSupply, address _authority) returns()
func (_MobyOLP *MobyOLPTransactor) BaseTokenInit(opts *bind.TransactOpts, _name string, _symbol string, _initialSupply *big.Int, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "__BaseToken_init__", _name, _symbol, _initialSupply, _authority)
}

// BaseTokenInit is a paid mutator transaction binding the contract method 0xfcc76648.
//
// Solidity: function __BaseToken_init__(string _name, string _symbol, uint256 _initialSupply, address _authority) returns()
func (_MobyOLP *MobyOLPSession) BaseTokenInit(_name string, _symbol string, _initialSupply *big.Int, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.BaseTokenInit(&_MobyOLP.TransactOpts, _name, _symbol, _initialSupply, _authority)
}

// BaseTokenInit is a paid mutator transaction binding the contract method 0xfcc76648.
//
// Solidity: function __BaseToken_init__(string _name, string _symbol, uint256 _initialSupply, address _authority) returns()
func (_MobyOLP *MobyOLPTransactorSession) BaseTokenInit(_name string, _symbol string, _initialSupply *big.Int, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.BaseTokenInit(&_MobyOLP.TransactOpts, _name, _symbol, _initialSupply, _authority)
}

// MintableBaseTokenInit is a paid mutator transaction binding the contract method 0xab1d235c.
//
// Solidity: function __MintableBaseToken_init__(string _name, string _symbol, uint256 _initialSupply, address _authority) returns()
func (_MobyOLP *MobyOLPTransactor) MintableBaseTokenInit(opts *bind.TransactOpts, _name string, _symbol string, _initialSupply *big.Int, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "__MintableBaseToken_init__", _name, _symbol, _initialSupply, _authority)
}

// MintableBaseTokenInit is a paid mutator transaction binding the contract method 0xab1d235c.
//
// Solidity: function __MintableBaseToken_init__(string _name, string _symbol, uint256 _initialSupply, address _authority) returns()
func (_MobyOLP *MobyOLPSession) MintableBaseTokenInit(_name string, _symbol string, _initialSupply *big.Int, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.MintableBaseTokenInit(&_MobyOLP.TransactOpts, _name, _symbol, _initialSupply, _authority)
}

// MintableBaseTokenInit is a paid mutator transaction binding the contract method 0xab1d235c.
//
// Solidity: function __MintableBaseToken_init__(string _name, string _symbol, uint256 _initialSupply, address _authority) returns()
func (_MobyOLP *MobyOLPTransactorSession) MintableBaseTokenInit(_name string, _symbol string, _initialSupply *big.Int, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.MintableBaseTokenInit(&_MobyOLP.TransactOpts, _name, _symbol, _initialSupply, _authority)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_MobyOLP *MobyOLPTransactor) AddNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "addNonStakingAccount", _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_MobyOLP *MobyOLPSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.AddNonStakingAccount(&_MobyOLP.TransactOpts, _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_MobyOLP *MobyOLPTransactorSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.AddNonStakingAccount(&_MobyOLP.TransactOpts, _account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_MobyOLP *MobyOLPTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_MobyOLP *MobyOLPSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.Approve(&_MobyOLP.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_MobyOLP *MobyOLPTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.Approve(&_MobyOLP.TransactOpts, _spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_MobyOLP *MobyOLPTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_MobyOLP *MobyOLPSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.Burn(&_MobyOLP.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_MobyOLP *MobyOLPTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.Burn(&_MobyOLP.TransactOpts, _account, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_MobyOLP *MobyOLPTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "claim", _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_MobyOLP *MobyOLPSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.Claim(&_MobyOLP.TransactOpts, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_MobyOLP *MobyOLPTransactorSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.Claim(&_MobyOLP.TransactOpts, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _authority) returns()
func (_MobyOLP *MobyOLPTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "initialize", _name, _symbol, _authority)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _authority) returns()
func (_MobyOLP *MobyOLPSession) Initialize(_name string, _symbol string, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.Initialize(&_MobyOLP.TransactOpts, _name, _symbol, _authority)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _authority) returns()
func (_MobyOLP *MobyOLPTransactorSession) Initialize(_name string, _symbol string, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.Initialize(&_MobyOLP.TransactOpts, _name, _symbol, _authority)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MobyOLP *MobyOLPTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MobyOLP *MobyOLPSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.Mint(&_MobyOLP.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MobyOLP *MobyOLPTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.Mint(&_MobyOLP.TransactOpts, _account, _amount)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_MobyOLP *MobyOLPTransactor) RecoverClaim(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "recoverClaim", _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_MobyOLP *MobyOLPSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.RecoverClaim(&_MobyOLP.TransactOpts, _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_MobyOLP *MobyOLPTransactorSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.RecoverClaim(&_MobyOLP.TransactOpts, _account, _receiver)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_MobyOLP *MobyOLPTransactor) RemoveNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "removeNonStakingAccount", _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_MobyOLP *MobyOLPSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.RemoveNonStakingAccount(&_MobyOLP.TransactOpts, _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_MobyOLP *MobyOLPTransactorSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.RemoveNonStakingAccount(&_MobyOLP.TransactOpts, _account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MobyOLP *MobyOLPTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MobyOLP *MobyOLPSession) RenounceOwnership() (*types.Transaction, error) {
	return _MobyOLP.Contract.RenounceOwnership(&_MobyOLP.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MobyOLP *MobyOLPTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MobyOLP.Contract.RenounceOwnership(&_MobyOLP.TransactOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _newAuthority) returns()
func (_MobyOLP *MobyOLPTransactor) SetAuthority(opts *bind.TransactOpts, _newAuthority common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "setAuthority", _newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _newAuthority) returns()
func (_MobyOLP *MobyOLPSession) SetAuthority(_newAuthority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetAuthority(&_MobyOLP.TransactOpts, _newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _newAuthority) returns()
func (_MobyOLP *MobyOLPTransactorSession) SetAuthority(_newAuthority common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetAuthority(&_MobyOLP.TransactOpts, _newAuthority)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_MobyOLP *MobyOLPTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_MobyOLP *MobyOLPSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetHandler(&_MobyOLP.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_MobyOLP *MobyOLPTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetHandler(&_MobyOLP.TransactOpts, _handler, _isActive)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_MobyOLP *MobyOLPTransactor) SetInPrivateTransferMode(opts *bind.TransactOpts, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "setInPrivateTransferMode", _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_MobyOLP *MobyOLPSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetInPrivateTransferMode(&_MobyOLP.TransactOpts, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_MobyOLP *MobyOLPTransactorSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetInPrivateTransferMode(&_MobyOLP.TransactOpts, _inPrivateTransferMode)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_MobyOLP *MobyOLPTransactor) SetInfo(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "setInfo", _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_MobyOLP *MobyOLPSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetInfo(&_MobyOLP.TransactOpts, _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_MobyOLP *MobyOLPTransactorSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetInfo(&_MobyOLP.TransactOpts, _name, _symbol)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_MobyOLP *MobyOLPTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "setMinter", _minter, _isActive)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_MobyOLP *MobyOLPSession) SetMinter(_minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetMinter(&_MobyOLP.TransactOpts, _minter, _isActive)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_MobyOLP *MobyOLPTransactorSession) SetMinter(_minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetMinter(&_MobyOLP.TransactOpts, _minter, _isActive)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_MobyOLP *MobyOLPTransactor) SetYieldTrackers(opts *bind.TransactOpts, _yieldTrackers []common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "setYieldTrackers", _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_MobyOLP *MobyOLPSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetYieldTrackers(&_MobyOLP.TransactOpts, _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_MobyOLP *MobyOLPTransactorSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.SetYieldTrackers(&_MobyOLP.TransactOpts, _yieldTrackers)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_MobyOLP *MobyOLPTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_MobyOLP *MobyOLPSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.Transfer(&_MobyOLP.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_MobyOLP *MobyOLPTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.Transfer(&_MobyOLP.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_MobyOLP *MobyOLPTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_MobyOLP *MobyOLPSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.TransferFrom(&_MobyOLP.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_MobyOLP *MobyOLPTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.TransferFrom(&_MobyOLP.TransactOpts, _sender, _recipient, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MobyOLP *MobyOLPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MobyOLP *MobyOLPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.TransferOwnership(&_MobyOLP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MobyOLP *MobyOLPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MobyOLP.Contract.TransferOwnership(&_MobyOLP.TransactOpts, newOwner)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_MobyOLP *MobyOLPTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_MobyOLP *MobyOLPSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.WithdrawToken(&_MobyOLP.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_MobyOLP *MobyOLPTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MobyOLP.Contract.WithdrawToken(&_MobyOLP.TransactOpts, _token, _account, _amount)
}

// MobyOLPApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MobyOLP contract.
type MobyOLPApprovalIterator struct {
	Event *MobyOLPApproval // Event containing the contract specifics and raw log

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
func (it *MobyOLPApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPApproval)
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
		it.Event = new(MobyOLPApproval)
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
func (it *MobyOLPApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPApproval represents a Approval event raised by the MobyOLP contract.
type MobyOLPApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MobyOLP *MobyOLPFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MobyOLPApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPApprovalIterator{contract: _MobyOLP.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MobyOLP *MobyOLPFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MobyOLPApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPApproval)
				if err := _MobyOLP.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MobyOLP *MobyOLPFilterer) ParseApproval(log types.Log) (*MobyOLPApproval, error) {
	event := new(MobyOLPApproval)
	if err := _MobyOLP.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the MobyOLP contract.
type MobyOLPAuthorityUpdatedIterator struct {
	Event *MobyOLPAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *MobyOLPAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPAuthorityUpdated)
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
		it.Event = new(MobyOLPAuthorityUpdated)
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
func (it *MobyOLPAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPAuthorityUpdated represents a AuthorityUpdated event raised by the MobyOLP contract.
type MobyOLPAuthorityUpdated struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address indexed authority)
func (_MobyOLP *MobyOLPFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts, authority []common.Address) (*MobyOLPAuthorityUpdatedIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "AuthorityUpdated", authorityRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPAuthorityUpdatedIterator{contract: _MobyOLP.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address indexed authority)
func (_MobyOLP *MobyOLPFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *MobyOLPAuthorityUpdated, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "AuthorityUpdated", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPAuthorityUpdated)
				if err := _MobyOLP.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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

// ParseAuthorityUpdated is a log parse operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address indexed authority)
func (_MobyOLP *MobyOLPFilterer) ParseAuthorityUpdated(log types.Log) (*MobyOLPAuthorityUpdated, error) {
	event := new(MobyOLPAuthorityUpdated)
	if err := _MobyOLP.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MobyOLP contract.
type MobyOLPOwnershipTransferredIterator struct {
	Event *MobyOLPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MobyOLPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPOwnershipTransferred)
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
		it.Event = new(MobyOLPOwnershipTransferred)
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
func (it *MobyOLPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPOwnershipTransferred represents a OwnershipTransferred event raised by the MobyOLP contract.
type MobyOLPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MobyOLP *MobyOLPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MobyOLPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPOwnershipTransferredIterator{contract: _MobyOLP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MobyOLP *MobyOLPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MobyOLPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPOwnershipTransferred)
				if err := _MobyOLP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MobyOLP *MobyOLPFilterer) ParseOwnershipTransferred(log types.Log) (*MobyOLPOwnershipTransferred, error) {
	event := new(MobyOLPOwnershipTransferred)
	if err := _MobyOLP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPSetHandlerIterator is returned from FilterSetHandler and is used to iterate over the raw logs and unpacked data for SetHandler events raised by the MobyOLP contract.
type MobyOLPSetHandlerIterator struct {
	Event *MobyOLPSetHandler // Event containing the contract specifics and raw log

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
func (it *MobyOLPSetHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPSetHandler)
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
		it.Event = new(MobyOLPSetHandler)
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
func (it *MobyOLPSetHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPSetHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPSetHandler represents a SetHandler event raised by the MobyOLP contract.
type MobyOLPSetHandler struct {
	Handler  common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetHandler is a free log retrieval operation binding the contract event 0xd373464a39404e5f98fdb4b152b8ba9a094561e97e5c4b4ea11eb18cd9e6695e.
//
// Solidity: event SetHandler(address indexed handler, bool isActive)
func (_MobyOLP *MobyOLPFilterer) FilterSetHandler(opts *bind.FilterOpts, handler []common.Address) (*MobyOLPSetHandlerIterator, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "SetHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPSetHandlerIterator{contract: _MobyOLP.contract, event: "SetHandler", logs: logs, sub: sub}, nil
}

// WatchSetHandler is a free log subscription operation binding the contract event 0xd373464a39404e5f98fdb4b152b8ba9a094561e97e5c4b4ea11eb18cd9e6695e.
//
// Solidity: event SetHandler(address indexed handler, bool isActive)
func (_MobyOLP *MobyOLPFilterer) WatchSetHandler(opts *bind.WatchOpts, sink chan<- *MobyOLPSetHandler, handler []common.Address) (event.Subscription, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "SetHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPSetHandler)
				if err := _MobyOLP.contract.UnpackLog(event, "SetHandler", log); err != nil {
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

// ParseSetHandler is a log parse operation binding the contract event 0xd373464a39404e5f98fdb4b152b8ba9a094561e97e5c4b4ea11eb18cd9e6695e.
//
// Solidity: event SetHandler(address indexed handler, bool isActive)
func (_MobyOLP *MobyOLPFilterer) ParseSetHandler(log types.Log) (*MobyOLPSetHandler, error) {
	event := new(MobyOLPSetHandler)
	if err := _MobyOLP.contract.UnpackLog(event, "SetHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPSetInPrivateTransferModeIterator is returned from FilterSetInPrivateTransferMode and is used to iterate over the raw logs and unpacked data for SetInPrivateTransferMode events raised by the MobyOLP contract.
type MobyOLPSetInPrivateTransferModeIterator struct {
	Event *MobyOLPSetInPrivateTransferMode // Event containing the contract specifics and raw log

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
func (it *MobyOLPSetInPrivateTransferModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPSetInPrivateTransferMode)
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
		it.Event = new(MobyOLPSetInPrivateTransferMode)
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
func (it *MobyOLPSetInPrivateTransferModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPSetInPrivateTransferModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPSetInPrivateTransferMode represents a SetInPrivateTransferMode event raised by the MobyOLP contract.
type MobyOLPSetInPrivateTransferMode struct {
	InPrivateTransferMode bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetInPrivateTransferMode is a free log retrieval operation binding the contract event 0x74c079335468f366bb20ab8e81de390d41699b16e4f47c18049c3fcb724346ad.
//
// Solidity: event SetInPrivateTransferMode(bool indexed inPrivateTransferMode)
func (_MobyOLP *MobyOLPFilterer) FilterSetInPrivateTransferMode(opts *bind.FilterOpts, inPrivateTransferMode []bool) (*MobyOLPSetInPrivateTransferModeIterator, error) {

	var inPrivateTransferModeRule []interface{}
	for _, inPrivateTransferModeItem := range inPrivateTransferMode {
		inPrivateTransferModeRule = append(inPrivateTransferModeRule, inPrivateTransferModeItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "SetInPrivateTransferMode", inPrivateTransferModeRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPSetInPrivateTransferModeIterator{contract: _MobyOLP.contract, event: "SetInPrivateTransferMode", logs: logs, sub: sub}, nil
}

// WatchSetInPrivateTransferMode is a free log subscription operation binding the contract event 0x74c079335468f366bb20ab8e81de390d41699b16e4f47c18049c3fcb724346ad.
//
// Solidity: event SetInPrivateTransferMode(bool indexed inPrivateTransferMode)
func (_MobyOLP *MobyOLPFilterer) WatchSetInPrivateTransferMode(opts *bind.WatchOpts, sink chan<- *MobyOLPSetInPrivateTransferMode, inPrivateTransferMode []bool) (event.Subscription, error) {

	var inPrivateTransferModeRule []interface{}
	for _, inPrivateTransferModeItem := range inPrivateTransferMode {
		inPrivateTransferModeRule = append(inPrivateTransferModeRule, inPrivateTransferModeItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "SetInPrivateTransferMode", inPrivateTransferModeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPSetInPrivateTransferMode)
				if err := _MobyOLP.contract.UnpackLog(event, "SetInPrivateTransferMode", log); err != nil {
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

// ParseSetInPrivateTransferMode is a log parse operation binding the contract event 0x74c079335468f366bb20ab8e81de390d41699b16e4f47c18049c3fcb724346ad.
//
// Solidity: event SetInPrivateTransferMode(bool indexed inPrivateTransferMode)
func (_MobyOLP *MobyOLPFilterer) ParseSetInPrivateTransferMode(log types.Log) (*MobyOLPSetInPrivateTransferMode, error) {
	event := new(MobyOLPSetInPrivateTransferMode)
	if err := _MobyOLP.contract.UnpackLog(event, "SetInPrivateTransferMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPSetInfoIterator is returned from FilterSetInfo and is used to iterate over the raw logs and unpacked data for SetInfo events raised by the MobyOLP contract.
type MobyOLPSetInfoIterator struct {
	Event *MobyOLPSetInfo // Event containing the contract specifics and raw log

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
func (it *MobyOLPSetInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPSetInfo)
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
		it.Event = new(MobyOLPSetInfo)
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
func (it *MobyOLPSetInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPSetInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPSetInfo represents a SetInfo event raised by the MobyOLP contract.
type MobyOLPSetInfo struct {
	Name   common.Hash
	Symbol common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetInfo is a free log retrieval operation binding the contract event 0x54bc9bd839ed82c38d3fac3bb1a3c878aa254a502c9b221d6b640db21c4ac7e3.
//
// Solidity: event SetInfo(string indexed name, string indexed symbol)
func (_MobyOLP *MobyOLPFilterer) FilterSetInfo(opts *bind.FilterOpts, name []string, symbol []string) (*MobyOLPSetInfoIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "SetInfo", nameRule, symbolRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPSetInfoIterator{contract: _MobyOLP.contract, event: "SetInfo", logs: logs, sub: sub}, nil
}

// WatchSetInfo is a free log subscription operation binding the contract event 0x54bc9bd839ed82c38d3fac3bb1a3c878aa254a502c9b221d6b640db21c4ac7e3.
//
// Solidity: event SetInfo(string indexed name, string indexed symbol)
func (_MobyOLP *MobyOLPFilterer) WatchSetInfo(opts *bind.WatchOpts, sink chan<- *MobyOLPSetInfo, name []string, symbol []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "SetInfo", nameRule, symbolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPSetInfo)
				if err := _MobyOLP.contract.UnpackLog(event, "SetInfo", log); err != nil {
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

// ParseSetInfo is a log parse operation binding the contract event 0x54bc9bd839ed82c38d3fac3bb1a3c878aa254a502c9b221d6b640db21c4ac7e3.
//
// Solidity: event SetInfo(string indexed name, string indexed symbol)
func (_MobyOLP *MobyOLPFilterer) ParseSetInfo(log types.Log) (*MobyOLPSetInfo, error) {
	event := new(MobyOLPSetInfo)
	if err := _MobyOLP.contract.UnpackLog(event, "SetInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPSetMinterIterator is returned from FilterSetMinter and is used to iterate over the raw logs and unpacked data for SetMinter events raised by the MobyOLP contract.
type MobyOLPSetMinterIterator struct {
	Event *MobyOLPSetMinter // Event containing the contract specifics and raw log

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
func (it *MobyOLPSetMinterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPSetMinter)
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
		it.Event = new(MobyOLPSetMinter)
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
func (it *MobyOLPSetMinterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPSetMinterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPSetMinter represents a SetMinter event raised by the MobyOLP contract.
type MobyOLPSetMinter struct {
	Minter   common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetMinter is a free log retrieval operation binding the contract event 0x1f96bc657d385fd83da973a43f2ad969e6d96b6779b779571a7306db7ca1cd00.
//
// Solidity: event SetMinter(address indexed minter, bool isActive)
func (_MobyOLP *MobyOLPFilterer) FilterSetMinter(opts *bind.FilterOpts, minter []common.Address) (*MobyOLPSetMinterIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "SetMinter", minterRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPSetMinterIterator{contract: _MobyOLP.contract, event: "SetMinter", logs: logs, sub: sub}, nil
}

// WatchSetMinter is a free log subscription operation binding the contract event 0x1f96bc657d385fd83da973a43f2ad969e6d96b6779b779571a7306db7ca1cd00.
//
// Solidity: event SetMinter(address indexed minter, bool isActive)
func (_MobyOLP *MobyOLPFilterer) WatchSetMinter(opts *bind.WatchOpts, sink chan<- *MobyOLPSetMinter, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "SetMinter", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPSetMinter)
				if err := _MobyOLP.contract.UnpackLog(event, "SetMinter", log); err != nil {
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

// ParseSetMinter is a log parse operation binding the contract event 0x1f96bc657d385fd83da973a43f2ad969e6d96b6779b779571a7306db7ca1cd00.
//
// Solidity: event SetMinter(address indexed minter, bool isActive)
func (_MobyOLP *MobyOLPFilterer) ParseSetMinter(log types.Log) (*MobyOLPSetMinter, error) {
	event := new(MobyOLPSetMinter)
	if err := _MobyOLP.contract.UnpackLog(event, "SetMinter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPSetYieldTrackersIterator is returned from FilterSetYieldTrackers and is used to iterate over the raw logs and unpacked data for SetYieldTrackers events raised by the MobyOLP contract.
type MobyOLPSetYieldTrackersIterator struct {
	Event *MobyOLPSetYieldTrackers // Event containing the contract specifics and raw log

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
func (it *MobyOLPSetYieldTrackersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPSetYieldTrackers)
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
		it.Event = new(MobyOLPSetYieldTrackers)
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
func (it *MobyOLPSetYieldTrackersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPSetYieldTrackersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPSetYieldTrackers represents a SetYieldTrackers event raised by the MobyOLP contract.
type MobyOLPSetYieldTrackers struct {
	YieldTrackers []common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetYieldTrackers is a free log retrieval operation binding the contract event 0x2495af74690b3dc8df76a2f12a034c96e4552b58b1824d96dced9b8e19590bef.
//
// Solidity: event SetYieldTrackers(address[] indexed yieldTrackers)
func (_MobyOLP *MobyOLPFilterer) FilterSetYieldTrackers(opts *bind.FilterOpts, yieldTrackers [][]common.Address) (*MobyOLPSetYieldTrackersIterator, error) {

	var yieldTrackersRule []interface{}
	for _, yieldTrackersItem := range yieldTrackers {
		yieldTrackersRule = append(yieldTrackersRule, yieldTrackersItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "SetYieldTrackers", yieldTrackersRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPSetYieldTrackersIterator{contract: _MobyOLP.contract, event: "SetYieldTrackers", logs: logs, sub: sub}, nil
}

// WatchSetYieldTrackers is a free log subscription operation binding the contract event 0x2495af74690b3dc8df76a2f12a034c96e4552b58b1824d96dced9b8e19590bef.
//
// Solidity: event SetYieldTrackers(address[] indexed yieldTrackers)
func (_MobyOLP *MobyOLPFilterer) WatchSetYieldTrackers(opts *bind.WatchOpts, sink chan<- *MobyOLPSetYieldTrackers, yieldTrackers [][]common.Address) (event.Subscription, error) {

	var yieldTrackersRule []interface{}
	for _, yieldTrackersItem := range yieldTrackers {
		yieldTrackersRule = append(yieldTrackersRule, yieldTrackersItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "SetYieldTrackers", yieldTrackersRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPSetYieldTrackers)
				if err := _MobyOLP.contract.UnpackLog(event, "SetYieldTrackers", log); err != nil {
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

// ParseSetYieldTrackers is a log parse operation binding the contract event 0x2495af74690b3dc8df76a2f12a034c96e4552b58b1824d96dced9b8e19590bef.
//
// Solidity: event SetYieldTrackers(address[] indexed yieldTrackers)
func (_MobyOLP *MobyOLPFilterer) ParseSetYieldTrackers(log types.Log) (*MobyOLPSetYieldTrackers, error) {
	event := new(MobyOLPSetYieldTrackers)
	if err := _MobyOLP.contract.UnpackLog(event, "SetYieldTrackers", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MobyOLP contract.
type MobyOLPTransferIterator struct {
	Event *MobyOLPTransfer // Event containing the contract specifics and raw log

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
func (it *MobyOLPTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPTransfer)
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
		it.Event = new(MobyOLPTransfer)
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
func (it *MobyOLPTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPTransfer represents a Transfer event raised by the MobyOLP contract.
type MobyOLPTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MobyOLP *MobyOLPFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MobyOLPTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPTransferIterator{contract: _MobyOLP.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MobyOLP *MobyOLPFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MobyOLPTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPTransfer)
				if err := _MobyOLP.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MobyOLP *MobyOLPFilterer) ParseTransfer(log types.Log) (*MobyOLPTransfer, error) {
	event := new(MobyOLPTransfer)
	if err := _MobyOLP.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPUpdateNonStakingAccountIterator is returned from FilterUpdateNonStakingAccount and is used to iterate over the raw logs and unpacked data for UpdateNonStakingAccount events raised by the MobyOLP contract.
type MobyOLPUpdateNonStakingAccountIterator struct {
	Event *MobyOLPUpdateNonStakingAccount // Event containing the contract specifics and raw log

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
func (it *MobyOLPUpdateNonStakingAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPUpdateNonStakingAccount)
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
		it.Event = new(MobyOLPUpdateNonStakingAccount)
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
func (it *MobyOLPUpdateNonStakingAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPUpdateNonStakingAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPUpdateNonStakingAccount represents a UpdateNonStakingAccount event raised by the MobyOLP contract.
type MobyOLPUpdateNonStakingAccount struct {
	Account             common.Address
	IsNonStakingAccount bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateNonStakingAccount is a free log retrieval operation binding the contract event 0x7887d715d63654633e6baadffa7c791c36784ccceb00b6c190706f317f8225bc.
//
// Solidity: event UpdateNonStakingAccount(address indexed account, bool isNonStakingAccount)
func (_MobyOLP *MobyOLPFilterer) FilterUpdateNonStakingAccount(opts *bind.FilterOpts, account []common.Address) (*MobyOLPUpdateNonStakingAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MobyOLP.contract.FilterLogs(opts, "UpdateNonStakingAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPUpdateNonStakingAccountIterator{contract: _MobyOLP.contract, event: "UpdateNonStakingAccount", logs: logs, sub: sub}, nil
}

// WatchUpdateNonStakingAccount is a free log subscription operation binding the contract event 0x7887d715d63654633e6baadffa7c791c36784ccceb00b6c190706f317f8225bc.
//
// Solidity: event UpdateNonStakingAccount(address indexed account, bool isNonStakingAccount)
func (_MobyOLP *MobyOLPFilterer) WatchUpdateNonStakingAccount(opts *bind.WatchOpts, sink chan<- *MobyOLPUpdateNonStakingAccount, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MobyOLP.contract.WatchLogs(opts, "UpdateNonStakingAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPUpdateNonStakingAccount)
				if err := _MobyOLP.contract.UnpackLog(event, "UpdateNonStakingAccount", log); err != nil {
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

// ParseUpdateNonStakingAccount is a log parse operation binding the contract event 0x7887d715d63654633e6baadffa7c791c36784ccceb00b6c190706f317f8225bc.
//
// Solidity: event UpdateNonStakingAccount(address indexed account, bool isNonStakingAccount)
func (_MobyOLP *MobyOLPFilterer) ParseUpdateNonStakingAccount(log types.Log) (*MobyOLPUpdateNonStakingAccount, error) {
	event := new(MobyOLPUpdateNonStakingAccount)
	if err := _MobyOLP.contract.UnpackLog(event, "UpdateNonStakingAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
