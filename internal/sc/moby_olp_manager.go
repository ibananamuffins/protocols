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

// MobyOLPManagerMetaData contains all meta data concerning the MobyOLPManager contract.
var MobyOLPManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ActionNotEnabled\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aumInUsdg\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"olpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIOptionsAuthority\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"olpAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aumInUsdg\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"olpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aumAddition\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aumDeduction\",\"type\":\"uint256\"}],\"name\":\"SetAumAdjustment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cooldownDuration\",\"type\":\"uint256\"}],\"name\":\"SetCooldownDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"SetHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"inPrivateMode\",\"type\":\"bool\"}],\"name\":\"SetInPrivateMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"SetVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultPriceFeed\",\"type\":\"address\"}],\"name\":\"SetVaultPriceFeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultUtils\",\"type\":\"address\"}],\"name\":\"SetVaultUtils\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_COOLDOWN_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OLP_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDG_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"_authority\",\"type\":\"address\"}],\"name\":\"__AuthorityUtil_init__\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOlp\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fundingAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minUsdg\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOlp\",\"type\":\"uint256\"}],\"name\":\"addLiquidityForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aumAddition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aumDeduction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getAum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getAumInUsdg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAums\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getOlpAssetAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getOlpAssetUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getOlpMpRpReleaseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getOlpMpRpReleaseUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getTotalOlpAssetUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultUtils\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultPriceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdg\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_olp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cooldownDuration\",\"type\":\"uint256\"},{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"_authority\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastAddedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"olp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_olpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_olpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"removeLiquidityForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_aumAddition\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_aumDeduction\",\"type\":\"uint256\"}],\"name\":\"setAumAdjustment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOptionsAuthority\",\"name\":\"_newAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cooldownDuration\",\"type\":\"uint256\"}],\"name\":\"setCooldownDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"setLastAddedAt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultPriceFeed\",\"type\":\"address\"}],\"name\":\"setVaultPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultUtils\",\"type\":\"address\"}],\"name\":\"setVaultUtils\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdg\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultPriceFeed\",\"outputs\":[{\"internalType\":\"contractIVaultPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultUtils\",\"outputs\":[{\"internalType\":\"contractIVaultUtils\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MobyOLPManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MobyOLPManagerMetaData.ABI instead.
var MobyOLPManagerABI = MobyOLPManagerMetaData.ABI

// MobyOLPManager is an auto generated Go binding around an Ethereum contract.
type MobyOLPManager struct {
	MobyOLPManagerCaller     // Read-only binding to the contract
	MobyOLPManagerTransactor // Write-only binding to the contract
	MobyOLPManagerFilterer   // Log filterer for contract events
}

// MobyOLPManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MobyOLPManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MobyOLPManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MobyOLPManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MobyOLPManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MobyOLPManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MobyOLPManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MobyOLPManagerSession struct {
	Contract     *MobyOLPManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MobyOLPManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MobyOLPManagerCallerSession struct {
	Contract *MobyOLPManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MobyOLPManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MobyOLPManagerTransactorSession struct {
	Contract     *MobyOLPManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MobyOLPManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MobyOLPManagerRaw struct {
	Contract *MobyOLPManager // Generic contract binding to access the raw methods on
}

// MobyOLPManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MobyOLPManagerCallerRaw struct {
	Contract *MobyOLPManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MobyOLPManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MobyOLPManagerTransactorRaw struct {
	Contract *MobyOLPManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMobyOLPManager creates a new instance of MobyOLPManager, bound to a specific deployed contract.
func NewMobyOLPManager(address common.Address, backend bind.ContractBackend) (*MobyOLPManager, error) {
	contract, err := bindMobyOLPManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManager{MobyOLPManagerCaller: MobyOLPManagerCaller{contract: contract}, MobyOLPManagerTransactor: MobyOLPManagerTransactor{contract: contract}, MobyOLPManagerFilterer: MobyOLPManagerFilterer{contract: contract}}, nil
}

// NewMobyOLPManagerCaller creates a new read-only instance of MobyOLPManager, bound to a specific deployed contract.
func NewMobyOLPManagerCaller(address common.Address, caller bind.ContractCaller) (*MobyOLPManagerCaller, error) {
	contract, err := bindMobyOLPManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerCaller{contract: contract}, nil
}

// NewMobyOLPManagerTransactor creates a new write-only instance of MobyOLPManager, bound to a specific deployed contract.
func NewMobyOLPManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MobyOLPManagerTransactor, error) {
	contract, err := bindMobyOLPManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerTransactor{contract: contract}, nil
}

// NewMobyOLPManagerFilterer creates a new log filterer instance of MobyOLPManager, bound to a specific deployed contract.
func NewMobyOLPManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MobyOLPManagerFilterer, error) {
	contract, err := bindMobyOLPManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerFilterer{contract: contract}, nil
}

// bindMobyOLPManager binds a generic wrapper to an already deployed contract.
func bindMobyOLPManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MobyOLPManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MobyOLPManager *MobyOLPManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MobyOLPManager.Contract.MobyOLPManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MobyOLPManager *MobyOLPManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.MobyOLPManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MobyOLPManager *MobyOLPManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.MobyOLPManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MobyOLPManager *MobyOLPManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MobyOLPManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MobyOLPManager *MobyOLPManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MobyOLPManager *MobyOLPManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _MobyOLPManager.Contract.BASISPOINTSDIVISOR(&_MobyOLPManager.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _MobyOLPManager.Contract.BASISPOINTSDIVISOR(&_MobyOLPManager.CallOpts)
}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) MAXCOOLDOWNDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "MAX_COOLDOWN_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) MAXCOOLDOWNDURATION() (*big.Int, error) {
	return _MobyOLPManager.Contract.MAXCOOLDOWNDURATION(&_MobyOLPManager.CallOpts)
}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) MAXCOOLDOWNDURATION() (*big.Int, error) {
	return _MobyOLPManager.Contract.MAXCOOLDOWNDURATION(&_MobyOLPManager.CallOpts)
}

// OLPPRECISION is a free data retrieval call binding the contract method 0x79e62857.
//
// Solidity: function OLP_PRECISION() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) OLPPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "OLP_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OLPPRECISION is a free data retrieval call binding the contract method 0x79e62857.
//
// Solidity: function OLP_PRECISION() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) OLPPRECISION() (*big.Int, error) {
	return _MobyOLPManager.Contract.OLPPRECISION(&_MobyOLPManager.CallOpts)
}

// OLPPRECISION is a free data retrieval call binding the contract method 0x79e62857.
//
// Solidity: function OLP_PRECISION() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) OLPPRECISION() (*big.Int, error) {
	return _MobyOLPManager.Contract.OLPPRECISION(&_MobyOLPManager.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) PRICEPRECISION() (*big.Int, error) {
	return _MobyOLPManager.Contract.PRICEPRECISION(&_MobyOLPManager.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _MobyOLPManager.Contract.PRICEPRECISION(&_MobyOLPManager.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) USDGDECIMALS() (*big.Int, error) {
	return _MobyOLPManager.Contract.USDGDECIMALS(&_MobyOLPManager.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) USDGDECIMALS() (*big.Int, error) {
	return _MobyOLPManager.Contract.USDGDECIMALS(&_MobyOLPManager.CallOpts)
}

// AumAddition is a free data retrieval call binding the contract method 0x196b68cb.
//
// Solidity: function aumAddition() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) AumAddition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "aumAddition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AumAddition is a free data retrieval call binding the contract method 0x196b68cb.
//
// Solidity: function aumAddition() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) AumAddition() (*big.Int, error) {
	return _MobyOLPManager.Contract.AumAddition(&_MobyOLPManager.CallOpts)
}

// AumAddition is a free data retrieval call binding the contract method 0x196b68cb.
//
// Solidity: function aumAddition() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) AumAddition() (*big.Int, error) {
	return _MobyOLPManager.Contract.AumAddition(&_MobyOLPManager.CallOpts)
}

// AumDeduction is a free data retrieval call binding the contract method 0xb172bb0c.
//
// Solidity: function aumDeduction() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) AumDeduction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "aumDeduction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AumDeduction is a free data retrieval call binding the contract method 0xb172bb0c.
//
// Solidity: function aumDeduction() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) AumDeduction() (*big.Int, error) {
	return _MobyOLPManager.Contract.AumDeduction(&_MobyOLPManager.CallOpts)
}

// AumDeduction is a free data retrieval call binding the contract method 0xb172bb0c.
//
// Solidity: function aumDeduction() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) AumDeduction() (*big.Int, error) {
	return _MobyOLPManager.Contract.AumDeduction(&_MobyOLPManager.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MobyOLPManager *MobyOLPManagerSession) Authority() (common.Address, error) {
	return _MobyOLPManager.Contract.Authority(&_MobyOLPManager.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCallerSession) Authority() (common.Address, error) {
	return _MobyOLPManager.Contract.Authority(&_MobyOLPManager.CallOpts)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) CooldownDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "cooldownDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) CooldownDuration() (*big.Int, error) {
	return _MobyOLPManager.Contract.CooldownDuration(&_MobyOLPManager.CallOpts)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) CooldownDuration() (*big.Int, error) {
	return _MobyOLPManager.Contract.CooldownDuration(&_MobyOLPManager.CallOpts)
}

// GetAum is a free data retrieval call binding the contract method 0x03391476.
//
// Solidity: function getAum(bool _maximise) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) GetAum(opts *bind.CallOpts, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "getAum", _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAum is a free data retrieval call binding the contract method 0x03391476.
//
// Solidity: function getAum(bool _maximise) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) GetAum(_maximise bool) (*big.Int, error) {
	return _MobyOLPManager.Contract.GetAum(&_MobyOLPManager.CallOpts, _maximise)
}

// GetAum is a free data retrieval call binding the contract method 0x03391476.
//
// Solidity: function getAum(bool _maximise) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) GetAum(_maximise bool) (*big.Int, error) {
	return _MobyOLPManager.Contract.GetAum(&_MobyOLPManager.CallOpts, _maximise)
}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool _maximise) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) GetAumInUsdg(opts *bind.CallOpts, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "getAumInUsdg", _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool _maximise) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) GetAumInUsdg(_maximise bool) (*big.Int, error) {
	return _MobyOLPManager.Contract.GetAumInUsdg(&_MobyOLPManager.CallOpts, _maximise)
}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool _maximise) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) GetAumInUsdg(_maximise bool) (*big.Int, error) {
	return _MobyOLPManager.Contract.GetAumInUsdg(&_MobyOLPManager.CallOpts, _maximise)
}

// GetAums is a free data retrieval call binding the contract method 0xed0d1c04.
//
// Solidity: function getAums() view returns(uint256[])
func (_MobyOLPManager *MobyOLPManagerCaller) GetAums(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "getAums")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAums is a free data retrieval call binding the contract method 0xed0d1c04.
//
// Solidity: function getAums() view returns(uint256[])
func (_MobyOLPManager *MobyOLPManagerSession) GetAums() ([]*big.Int, error) {
	return _MobyOLPManager.Contract.GetAums(&_MobyOLPManager.CallOpts)
}

// GetAums is a free data retrieval call binding the contract method 0xed0d1c04.
//
// Solidity: function getAums() view returns(uint256[])
func (_MobyOLPManager *MobyOLPManagerCallerSession) GetAums() ([]*big.Int, error) {
	return _MobyOLPManager.Contract.GetAums(&_MobyOLPManager.CallOpts)
}

// GetOlpAssetAmounts is a free data retrieval call binding the contract method 0xb61df8b2.
//
// Solidity: function getOlpAssetAmounts(address _token) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) GetOlpAssetAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "getOlpAssetAmounts", _token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetOlpAssetAmounts is a free data retrieval call binding the contract method 0xb61df8b2.
//
// Solidity: function getOlpAssetAmounts(address _token) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerSession) GetOlpAssetAmounts(_token common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetOlpAssetAmounts(&_MobyOLPManager.CallOpts, _token)
}

// GetOlpAssetAmounts is a free data retrieval call binding the contract method 0xb61df8b2.
//
// Solidity: function getOlpAssetAmounts(address _token) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) GetOlpAssetAmounts(_token common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetOlpAssetAmounts(&_MobyOLPManager.CallOpts, _token)
}

// GetOlpAssetUsd is a free data retrieval call binding the contract method 0xbc702308.
//
// Solidity: function getOlpAssetUsd(address _token, bool _maximise) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) GetOlpAssetUsd(opts *bind.CallOpts, _token common.Address, _maximise bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "getOlpAssetUsd", _token, _maximise)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetOlpAssetUsd is a free data retrieval call binding the contract method 0xbc702308.
//
// Solidity: function getOlpAssetUsd(address _token, bool _maximise) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerSession) GetOlpAssetUsd(_token common.Address, _maximise bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetOlpAssetUsd(&_MobyOLPManager.CallOpts, _token, _maximise)
}

// GetOlpAssetUsd is a free data retrieval call binding the contract method 0xbc702308.
//
// Solidity: function getOlpAssetUsd(address _token, bool _maximise) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) GetOlpAssetUsd(_token common.Address, _maximise bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetOlpAssetUsd(&_MobyOLPManager.CallOpts, _token, _maximise)
}

// GetOlpMpRpReleaseRate is a free data retrieval call binding the contract method 0x83249e79.
//
// Solidity: function getOlpMpRpReleaseRate(address _token) view returns(uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) GetOlpMpRpReleaseRate(opts *bind.CallOpts, _token common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "getOlpMpRpReleaseRate", _token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetOlpMpRpReleaseRate is a free data retrieval call binding the contract method 0x83249e79.
//
// Solidity: function getOlpMpRpReleaseRate(address _token) view returns(uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerSession) GetOlpMpRpReleaseRate(_token common.Address) (*big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetOlpMpRpReleaseRate(&_MobyOLPManager.CallOpts, _token)
}

// GetOlpMpRpReleaseRate is a free data retrieval call binding the contract method 0x83249e79.
//
// Solidity: function getOlpMpRpReleaseRate(address _token) view returns(uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) GetOlpMpRpReleaseRate(_token common.Address) (*big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetOlpMpRpReleaseRate(&_MobyOLPManager.CallOpts, _token)
}

// GetOlpMpRpReleaseUsd is a free data retrieval call binding the contract method 0xd8866253.
//
// Solidity: function getOlpMpRpReleaseUsd(bool _maximise) view returns(uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) GetOlpMpRpReleaseUsd(opts *bind.CallOpts, _maximise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "getOlpMpRpReleaseUsd", _maximise)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetOlpMpRpReleaseUsd is a free data retrieval call binding the contract method 0xd8866253.
//
// Solidity: function getOlpMpRpReleaseUsd(bool _maximise) view returns(uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerSession) GetOlpMpRpReleaseUsd(_maximise bool) (*big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetOlpMpRpReleaseUsd(&_MobyOLPManager.CallOpts, _maximise)
}

// GetOlpMpRpReleaseUsd is a free data retrieval call binding the contract method 0xd8866253.
//
// Solidity: function getOlpMpRpReleaseUsd(bool _maximise) view returns(uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) GetOlpMpRpReleaseUsd(_maximise bool) (*big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetOlpMpRpReleaseUsd(&_MobyOLPManager.CallOpts, _maximise)
}

// GetPrice is a free data retrieval call binding the contract method 0xe245b5af.
//
// Solidity: function getPrice(bool _maximise) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) GetPrice(opts *bind.CallOpts, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "getPrice", _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0xe245b5af.
//
// Solidity: function getPrice(bool _maximise) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) GetPrice(_maximise bool) (*big.Int, error) {
	return _MobyOLPManager.Contract.GetPrice(&_MobyOLPManager.CallOpts, _maximise)
}

// GetPrice is a free data retrieval call binding the contract method 0xe245b5af.
//
// Solidity: function getPrice(bool _maximise) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) GetPrice(_maximise bool) (*big.Int, error) {
	return _MobyOLPManager.Contract.GetPrice(&_MobyOLPManager.CallOpts, _maximise)
}

// GetTotalOlpAssetUsd is a free data retrieval call binding the contract method 0x5767dea8.
//
// Solidity: function getTotalOlpAssetUsd(bool _maximise) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) GetTotalOlpAssetUsd(opts *bind.CallOpts, _maximise bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "getTotalOlpAssetUsd", _maximise)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetTotalOlpAssetUsd is a free data retrieval call binding the contract method 0x5767dea8.
//
// Solidity: function getTotalOlpAssetUsd(bool _maximise) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerSession) GetTotalOlpAssetUsd(_maximise bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetTotalOlpAssetUsd(&_MobyOLPManager.CallOpts, _maximise)
}

// GetTotalOlpAssetUsd is a free data retrieval call binding the contract method 0x5767dea8.
//
// Solidity: function getTotalOlpAssetUsd(bool _maximise) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) GetTotalOlpAssetUsd(_maximise bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MobyOLPManager.Contract.GetTotalOlpAssetUsd(&_MobyOLPManager.CallOpts, _maximise)
}

// InPrivateMode is a free data retrieval call binding the contract method 0x070eacee.
//
// Solidity: function inPrivateMode() view returns(bool)
func (_MobyOLPManager *MobyOLPManagerCaller) InPrivateMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "inPrivateMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateMode is a free data retrieval call binding the contract method 0x070eacee.
//
// Solidity: function inPrivateMode() view returns(bool)
func (_MobyOLPManager *MobyOLPManagerSession) InPrivateMode() (bool, error) {
	return _MobyOLPManager.Contract.InPrivateMode(&_MobyOLPManager.CallOpts)
}

// InPrivateMode is a free data retrieval call binding the contract method 0x070eacee.
//
// Solidity: function inPrivateMode() view returns(bool)
func (_MobyOLPManager *MobyOLPManagerCallerSession) InPrivateMode() (bool, error) {
	return _MobyOLPManager.Contract.InPrivateMode(&_MobyOLPManager.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_MobyOLPManager *MobyOLPManagerCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_MobyOLPManager *MobyOLPManagerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _MobyOLPManager.Contract.IsHandler(&_MobyOLPManager.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_MobyOLPManager *MobyOLPManagerCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _MobyOLPManager.Contract.IsHandler(&_MobyOLPManager.CallOpts, arg0)
}

// LastAddedAt is a free data retrieval call binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address ) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCaller) LastAddedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "lastAddedAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAddedAt is a free data retrieval call binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address ) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) LastAddedAt(arg0 common.Address) (*big.Int, error) {
	return _MobyOLPManager.Contract.LastAddedAt(&_MobyOLPManager.CallOpts, arg0)
}

// LastAddedAt is a free data retrieval call binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address ) view returns(uint256)
func (_MobyOLPManager *MobyOLPManagerCallerSession) LastAddedAt(arg0 common.Address) (*big.Int, error) {
	return _MobyOLPManager.Contract.LastAddedAt(&_MobyOLPManager.CallOpts, arg0)
}

// Olp is a free data retrieval call binding the contract method 0xa2709a33.
//
// Solidity: function olp() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCaller) Olp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "olp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Olp is a free data retrieval call binding the contract method 0xa2709a33.
//
// Solidity: function olp() view returns(address)
func (_MobyOLPManager *MobyOLPManagerSession) Olp() (common.Address, error) {
	return _MobyOLPManager.Contract.Olp(&_MobyOLPManager.CallOpts)
}

// Olp is a free data retrieval call binding the contract method 0xa2709a33.
//
// Solidity: function olp() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCallerSession) Olp() (common.Address, error) {
	return _MobyOLPManager.Contract.Olp(&_MobyOLPManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MobyOLPManager *MobyOLPManagerSession) Owner() (common.Address, error) {
	return _MobyOLPManager.Contract.Owner(&_MobyOLPManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCallerSession) Owner() (common.Address, error) {
	return _MobyOLPManager.Contract.Owner(&_MobyOLPManager.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_MobyOLPManager *MobyOLPManagerSession) Usdg() (common.Address, error) {
	return _MobyOLPManager.Contract.Usdg(&_MobyOLPManager.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCallerSession) Usdg() (common.Address, error) {
	return _MobyOLPManager.Contract.Usdg(&_MobyOLPManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_MobyOLPManager *MobyOLPManagerSession) Vault() (common.Address, error) {
	return _MobyOLPManager.Contract.Vault(&_MobyOLPManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCallerSession) Vault() (common.Address, error) {
	return _MobyOLPManager.Contract.Vault(&_MobyOLPManager.CallOpts)
}

// VaultPriceFeed is a free data retrieval call binding the contract method 0xeeaa783a.
//
// Solidity: function vaultPriceFeed() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCaller) VaultPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "vaultPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultPriceFeed is a free data retrieval call binding the contract method 0xeeaa783a.
//
// Solidity: function vaultPriceFeed() view returns(address)
func (_MobyOLPManager *MobyOLPManagerSession) VaultPriceFeed() (common.Address, error) {
	return _MobyOLPManager.Contract.VaultPriceFeed(&_MobyOLPManager.CallOpts)
}

// VaultPriceFeed is a free data retrieval call binding the contract method 0xeeaa783a.
//
// Solidity: function vaultPriceFeed() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCallerSession) VaultPriceFeed() (common.Address, error) {
	return _MobyOLPManager.Contract.VaultPriceFeed(&_MobyOLPManager.CallOpts)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCaller) VaultUtils(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MobyOLPManager.contract.Call(opts, &out, "vaultUtils")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_MobyOLPManager *MobyOLPManagerSession) VaultUtils() (common.Address, error) {
	return _MobyOLPManager.Contract.VaultUtils(&_MobyOLPManager.CallOpts)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_MobyOLPManager *MobyOLPManagerCallerSession) VaultUtils() (common.Address, error) {
	return _MobyOLPManager.Contract.VaultUtils(&_MobyOLPManager.CallOpts)
}

// AuthorityUtilInit is a paid mutator transaction binding the contract method 0xbdaf7884.
//
// Solidity: function __AuthorityUtil_init__(address _authority) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) AuthorityUtilInit(opts *bind.TransactOpts, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "__AuthorityUtil_init__", _authority)
}

// AuthorityUtilInit is a paid mutator transaction binding the contract method 0xbdaf7884.
//
// Solidity: function __AuthorityUtil_init__(address _authority) returns()
func (_MobyOLPManager *MobyOLPManagerSession) AuthorityUtilInit(_authority common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.AuthorityUtilInit(&_MobyOLPManager.TransactOpts, _authority)
}

// AuthorityUtilInit is a paid mutator transaction binding the contract method 0xbdaf7884.
//
// Solidity: function __AuthorityUtil_init__(address _authority) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) AuthorityUtilInit(_authority common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.AuthorityUtilInit(&_MobyOLPManager.TransactOpts, _authority)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minOlp) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minOlp *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "addLiquidity", _token, _amount, _minUsdg, _minOlp)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minOlp) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) AddLiquidity(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minOlp *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.AddLiquidity(&_MobyOLPManager.TransactOpts, _token, _amount, _minUsdg, _minOlp)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minOlp) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minOlp *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.AddLiquidity(&_MobyOLPManager.TransactOpts, _token, _amount, _minUsdg, _minOlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minOlp) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerTransactor) AddLiquidityForAccount(opts *bind.TransactOpts, _fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minOlp *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "addLiquidityForAccount", _fundingAccount, _account, _token, _amount, _minUsdg, _minOlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minOlp) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) AddLiquidityForAccount(_fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minOlp *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.AddLiquidityForAccount(&_MobyOLPManager.TransactOpts, _fundingAccount, _account, _token, _amount, _minUsdg, _minOlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minOlp) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerTransactorSession) AddLiquidityForAccount(_fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minOlp *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.AddLiquidityForAccount(&_MobyOLPManager.TransactOpts, _fundingAccount, _account, _token, _amount, _minUsdg, _minOlp)
}

// Initialize is a paid mutator transaction binding the contract method 0xeabcca10.
//
// Solidity: function initialize(address _vault, address _vaultUtils, address _vaultPriceFeed, address _usdg, address _olp, uint256 _cooldownDuration, address _authority) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) Initialize(opts *bind.TransactOpts, _vault common.Address, _vaultUtils common.Address, _vaultPriceFeed common.Address, _usdg common.Address, _olp common.Address, _cooldownDuration *big.Int, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "initialize", _vault, _vaultUtils, _vaultPriceFeed, _usdg, _olp, _cooldownDuration, _authority)
}

// Initialize is a paid mutator transaction binding the contract method 0xeabcca10.
//
// Solidity: function initialize(address _vault, address _vaultUtils, address _vaultPriceFeed, address _usdg, address _olp, uint256 _cooldownDuration, address _authority) returns()
func (_MobyOLPManager *MobyOLPManagerSession) Initialize(_vault common.Address, _vaultUtils common.Address, _vaultPriceFeed common.Address, _usdg common.Address, _olp common.Address, _cooldownDuration *big.Int, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.Initialize(&_MobyOLPManager.TransactOpts, _vault, _vaultUtils, _vaultPriceFeed, _usdg, _olp, _cooldownDuration, _authority)
}

// Initialize is a paid mutator transaction binding the contract method 0xeabcca10.
//
// Solidity: function initialize(address _vault, address _vaultUtils, address _vaultPriceFeed, address _usdg, address _olp, uint256 _cooldownDuration, address _authority) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) Initialize(_vault common.Address, _vaultUtils common.Address, _vaultPriceFeed common.Address, _usdg common.Address, _olp common.Address, _cooldownDuration *big.Int, _authority common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.Initialize(&_MobyOLPManager.TransactOpts, _vault, _vaultUtils, _vaultPriceFeed, _usdg, _olp, _cooldownDuration, _authority)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _olpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerTransactor) RemoveLiquidity(opts *bind.TransactOpts, _tokenOut common.Address, _olpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "removeLiquidity", _tokenOut, _olpAmount, _minOut, _receiver)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _olpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) RemoveLiquidity(_tokenOut common.Address, _olpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.RemoveLiquidity(&_MobyOLPManager.TransactOpts, _tokenOut, _olpAmount, _minOut, _receiver)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _olpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerTransactorSession) RemoveLiquidity(_tokenOut common.Address, _olpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.RemoveLiquidity(&_MobyOLPManager.TransactOpts, _tokenOut, _olpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _olpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerTransactor) RemoveLiquidityForAccount(opts *bind.TransactOpts, _account common.Address, _tokenOut common.Address, _olpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "removeLiquidityForAccount", _account, _tokenOut, _olpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _olpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerSession) RemoveLiquidityForAccount(_account common.Address, _tokenOut common.Address, _olpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.RemoveLiquidityForAccount(&_MobyOLPManager.TransactOpts, _account, _tokenOut, _olpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _olpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_MobyOLPManager *MobyOLPManagerTransactorSession) RemoveLiquidityForAccount(_account common.Address, _tokenOut common.Address, _olpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.RemoveLiquidityForAccount(&_MobyOLPManager.TransactOpts, _account, _tokenOut, _olpAmount, _minOut, _receiver)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MobyOLPManager *MobyOLPManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _MobyOLPManager.Contract.RenounceOwnership(&_MobyOLPManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MobyOLPManager.Contract.RenounceOwnership(&_MobyOLPManager.TransactOpts)
}

// SetAumAdjustment is a paid mutator transaction binding the contract method 0x9116c4ae.
//
// Solidity: function setAumAdjustment(uint256 _aumAddition, uint256 _aumDeduction) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) SetAumAdjustment(opts *bind.TransactOpts, _aumAddition *big.Int, _aumDeduction *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "setAumAdjustment", _aumAddition, _aumDeduction)
}

// SetAumAdjustment is a paid mutator transaction binding the contract method 0x9116c4ae.
//
// Solidity: function setAumAdjustment(uint256 _aumAddition, uint256 _aumDeduction) returns()
func (_MobyOLPManager *MobyOLPManagerSession) SetAumAdjustment(_aumAddition *big.Int, _aumDeduction *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetAumAdjustment(&_MobyOLPManager.TransactOpts, _aumAddition, _aumDeduction)
}

// SetAumAdjustment is a paid mutator transaction binding the contract method 0x9116c4ae.
//
// Solidity: function setAumAdjustment(uint256 _aumAddition, uint256 _aumDeduction) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) SetAumAdjustment(_aumAddition *big.Int, _aumDeduction *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetAumAdjustment(&_MobyOLPManager.TransactOpts, _aumAddition, _aumDeduction)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _newAuthority) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) SetAuthority(opts *bind.TransactOpts, _newAuthority common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "setAuthority", _newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _newAuthority) returns()
func (_MobyOLPManager *MobyOLPManagerSession) SetAuthority(_newAuthority common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetAuthority(&_MobyOLPManager.TransactOpts, _newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _newAuthority) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) SetAuthority(_newAuthority common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetAuthority(&_MobyOLPManager.TransactOpts, _newAuthority)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) SetCooldownDuration(opts *bind.TransactOpts, _cooldownDuration *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "setCooldownDuration", _cooldownDuration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_MobyOLPManager *MobyOLPManagerSession) SetCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetCooldownDuration(&_MobyOLPManager.TransactOpts, _cooldownDuration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) SetCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetCooldownDuration(&_MobyOLPManager.TransactOpts, _cooldownDuration)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_MobyOLPManager *MobyOLPManagerSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetHandler(&_MobyOLPManager.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetHandler(&_MobyOLPManager.TransactOpts, _handler, _isActive)
}

// SetInPrivateMode is a paid mutator transaction binding the contract method 0x6a86da19.
//
// Solidity: function setInPrivateMode(bool _inPrivateMode) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) SetInPrivateMode(opts *bind.TransactOpts, _inPrivateMode bool) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "setInPrivateMode", _inPrivateMode)
}

// SetInPrivateMode is a paid mutator transaction binding the contract method 0x6a86da19.
//
// Solidity: function setInPrivateMode(bool _inPrivateMode) returns()
func (_MobyOLPManager *MobyOLPManagerSession) SetInPrivateMode(_inPrivateMode bool) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetInPrivateMode(&_MobyOLPManager.TransactOpts, _inPrivateMode)
}

// SetInPrivateMode is a paid mutator transaction binding the contract method 0x6a86da19.
//
// Solidity: function setInPrivateMode(bool _inPrivateMode) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) SetInPrivateMode(_inPrivateMode bool) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetInPrivateMode(&_MobyOLPManager.TransactOpts, _inPrivateMode)
}

// SetLastAddedAt is a paid mutator transaction binding the contract method 0x0cf22696.
//
// Solidity: function setLastAddedAt(address _account, uint256 _timestamp) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) SetLastAddedAt(opts *bind.TransactOpts, _account common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "setLastAddedAt", _account, _timestamp)
}

// SetLastAddedAt is a paid mutator transaction binding the contract method 0x0cf22696.
//
// Solidity: function setLastAddedAt(address _account, uint256 _timestamp) returns()
func (_MobyOLPManager *MobyOLPManagerSession) SetLastAddedAt(_account common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetLastAddedAt(&_MobyOLPManager.TransactOpts, _account, _timestamp)
}

// SetLastAddedAt is a paid mutator transaction binding the contract method 0x0cf22696.
//
// Solidity: function setLastAddedAt(address _account, uint256 _timestamp) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) SetLastAddedAt(_account common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetLastAddedAt(&_MobyOLPManager.TransactOpts, _account, _timestamp)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) SetVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "setVault", _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_MobyOLPManager *MobyOLPManagerSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetVault(&_MobyOLPManager.TransactOpts, _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetVault(&_MobyOLPManager.TransactOpts, _vault)
}

// SetVaultPriceFeed is a paid mutator transaction binding the contract method 0x238aafb7.
//
// Solidity: function setVaultPriceFeed(address _vaultPriceFeed) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) SetVaultPriceFeed(opts *bind.TransactOpts, _vaultPriceFeed common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "setVaultPriceFeed", _vaultPriceFeed)
}

// SetVaultPriceFeed is a paid mutator transaction binding the contract method 0x238aafb7.
//
// Solidity: function setVaultPriceFeed(address _vaultPriceFeed) returns()
func (_MobyOLPManager *MobyOLPManagerSession) SetVaultPriceFeed(_vaultPriceFeed common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetVaultPriceFeed(&_MobyOLPManager.TransactOpts, _vaultPriceFeed)
}

// SetVaultPriceFeed is a paid mutator transaction binding the contract method 0x238aafb7.
//
// Solidity: function setVaultPriceFeed(address _vaultPriceFeed) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) SetVaultPriceFeed(_vaultPriceFeed common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetVaultPriceFeed(&_MobyOLPManager.TransactOpts, _vaultPriceFeed)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) SetVaultUtils(opts *bind.TransactOpts, _vaultUtils common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "setVaultUtils", _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_MobyOLPManager *MobyOLPManagerSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetVaultUtils(&_MobyOLPManager.TransactOpts, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.SetVaultUtils(&_MobyOLPManager.TransactOpts, _vaultUtils)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MobyOLPManager *MobyOLPManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MobyOLPManager *MobyOLPManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.TransferOwnership(&_MobyOLPManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MobyOLPManager *MobyOLPManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MobyOLPManager.Contract.TransferOwnership(&_MobyOLPManager.TransactOpts, newOwner)
}

// MobyOLPManagerAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the MobyOLPManager contract.
type MobyOLPManagerAddLiquidityIterator struct {
	Event *MobyOLPManagerAddLiquidity // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerAddLiquidity)
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
		it.Event = new(MobyOLPManagerAddLiquidity)
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
func (it *MobyOLPManagerAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerAddLiquidity represents a AddLiquidity event raised by the MobyOLPManager contract.
type MobyOLPManagerAddLiquidity struct {
	Account    common.Address
	Token      common.Address
	Amount     *big.Int
	AumInUsdg  *big.Int
	OlpSupply  *big.Int
	UsdgAmount *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x38dc38b96482be64113daffd8d464ebda93e856b70ccfc605e69ccf892ab981e.
//
// Solidity: event AddLiquidity(address indexed account, address indexed token, uint256 amount, uint256 aumInUsdg, uint256 olpSupply, uint256 usdgAmount, uint256 mintAmount)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterAddLiquidity(opts *bind.FilterOpts, account []common.Address, token []common.Address) (*MobyOLPManagerAddLiquidityIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "AddLiquidity", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerAddLiquidityIterator{contract: _MobyOLPManager.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x38dc38b96482be64113daffd8d464ebda93e856b70ccfc605e69ccf892ab981e.
//
// Solidity: event AddLiquidity(address indexed account, address indexed token, uint256 amount, uint256 aumInUsdg, uint256 olpSupply, uint256 usdgAmount, uint256 mintAmount)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerAddLiquidity, account []common.Address, token []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "AddLiquidity", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerAddLiquidity)
				if err := _MobyOLPManager.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x38dc38b96482be64113daffd8d464ebda93e856b70ccfc605e69ccf892ab981e.
//
// Solidity: event AddLiquidity(address indexed account, address indexed token, uint256 amount, uint256 aumInUsdg, uint256 olpSupply, uint256 usdgAmount, uint256 mintAmount)
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseAddLiquidity(log types.Log) (*MobyOLPManagerAddLiquidity, error) {
	event := new(MobyOLPManagerAddLiquidity)
	if err := _MobyOLPManager.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the MobyOLPManager contract.
type MobyOLPManagerAuthorityUpdatedIterator struct {
	Event *MobyOLPManagerAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerAuthorityUpdated)
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
		it.Event = new(MobyOLPManagerAuthorityUpdated)
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
func (it *MobyOLPManagerAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerAuthorityUpdated represents a AuthorityUpdated event raised by the MobyOLPManager contract.
type MobyOLPManagerAuthorityUpdated struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address indexed authority)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts, authority []common.Address) (*MobyOLPManagerAuthorityUpdatedIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "AuthorityUpdated", authorityRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerAuthorityUpdatedIterator{contract: _MobyOLPManager.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address indexed authority)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerAuthorityUpdated, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "AuthorityUpdated", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerAuthorityUpdated)
				if err := _MobyOLPManager.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseAuthorityUpdated(log types.Log) (*MobyOLPManagerAuthorityUpdated, error) {
	event := new(MobyOLPManagerAuthorityUpdated)
	if err := _MobyOLPManager.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MobyOLPManager contract.
type MobyOLPManagerOwnershipTransferredIterator struct {
	Event *MobyOLPManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerOwnershipTransferred)
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
		it.Event = new(MobyOLPManagerOwnershipTransferred)
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
func (it *MobyOLPManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerOwnershipTransferred represents a OwnershipTransferred event raised by the MobyOLPManager contract.
type MobyOLPManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MobyOLPManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerOwnershipTransferredIterator{contract: _MobyOLPManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerOwnershipTransferred)
				if err := _MobyOLPManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseOwnershipTransferred(log types.Log) (*MobyOLPManagerOwnershipTransferred, error) {
	event := new(MobyOLPManagerOwnershipTransferred)
	if err := _MobyOLPManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the MobyOLPManager contract.
type MobyOLPManagerRemoveLiquidityIterator struct {
	Event *MobyOLPManagerRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerRemoveLiquidity)
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
		it.Event = new(MobyOLPManagerRemoveLiquidity)
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
func (it *MobyOLPManagerRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerRemoveLiquidity represents a RemoveLiquidity event raised by the MobyOLPManager contract.
type MobyOLPManagerRemoveLiquidity struct {
	Account    common.Address
	Token      common.Address
	OlpAmount  *big.Int
	AumInUsdg  *big.Int
	OlpSupply  *big.Int
	UsdgAmount *big.Int
	AmountOut  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x87b9679bb9a4944bafa98c267e7cd4a00ab29fed48afdefae25f0fca5da27940.
//
// Solidity: event RemoveLiquidity(address indexed account, address indexed token, uint256 olpAmount, uint256 aumInUsdg, uint256 olpSupply, uint256 usdgAmount, uint256 amountOut)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, account []common.Address, token []common.Address) (*MobyOLPManagerRemoveLiquidityIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "RemoveLiquidity", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerRemoveLiquidityIterator{contract: _MobyOLPManager.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x87b9679bb9a4944bafa98c267e7cd4a00ab29fed48afdefae25f0fca5da27940.
//
// Solidity: event RemoveLiquidity(address indexed account, address indexed token, uint256 olpAmount, uint256 aumInUsdg, uint256 olpSupply, uint256 usdgAmount, uint256 amountOut)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerRemoveLiquidity, account []common.Address, token []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "RemoveLiquidity", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerRemoveLiquidity)
				if err := _MobyOLPManager.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x87b9679bb9a4944bafa98c267e7cd4a00ab29fed48afdefae25f0fca5da27940.
//
// Solidity: event RemoveLiquidity(address indexed account, address indexed token, uint256 olpAmount, uint256 aumInUsdg, uint256 olpSupply, uint256 usdgAmount, uint256 amountOut)
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseRemoveLiquidity(log types.Log) (*MobyOLPManagerRemoveLiquidity, error) {
	event := new(MobyOLPManagerRemoveLiquidity)
	if err := _MobyOLPManager.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerSetAumAdjustmentIterator is returned from FilterSetAumAdjustment and is used to iterate over the raw logs and unpacked data for SetAumAdjustment events raised by the MobyOLPManager contract.
type MobyOLPManagerSetAumAdjustmentIterator struct {
	Event *MobyOLPManagerSetAumAdjustment // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerSetAumAdjustmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerSetAumAdjustment)
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
		it.Event = new(MobyOLPManagerSetAumAdjustment)
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
func (it *MobyOLPManagerSetAumAdjustmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerSetAumAdjustmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerSetAumAdjustment represents a SetAumAdjustment event raised by the MobyOLPManager contract.
type MobyOLPManagerSetAumAdjustment struct {
	AumAddition  *big.Int
	AumDeduction *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetAumAdjustment is a free log retrieval operation binding the contract event 0x6053f2196ef566290f5a0c7e42705ab4fe5dd826a82eb2c6d4cf0133adfcd22b.
//
// Solidity: event SetAumAdjustment(uint256 aumAddition, uint256 aumDeduction)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterSetAumAdjustment(opts *bind.FilterOpts) (*MobyOLPManagerSetAumAdjustmentIterator, error) {

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "SetAumAdjustment")
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerSetAumAdjustmentIterator{contract: _MobyOLPManager.contract, event: "SetAumAdjustment", logs: logs, sub: sub}, nil
}

// WatchSetAumAdjustment is a free log subscription operation binding the contract event 0x6053f2196ef566290f5a0c7e42705ab4fe5dd826a82eb2c6d4cf0133adfcd22b.
//
// Solidity: event SetAumAdjustment(uint256 aumAddition, uint256 aumDeduction)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchSetAumAdjustment(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerSetAumAdjustment) (event.Subscription, error) {

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "SetAumAdjustment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerSetAumAdjustment)
				if err := _MobyOLPManager.contract.UnpackLog(event, "SetAumAdjustment", log); err != nil {
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

// ParseSetAumAdjustment is a log parse operation binding the contract event 0x6053f2196ef566290f5a0c7e42705ab4fe5dd826a82eb2c6d4cf0133adfcd22b.
//
// Solidity: event SetAumAdjustment(uint256 aumAddition, uint256 aumDeduction)
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseSetAumAdjustment(log types.Log) (*MobyOLPManagerSetAumAdjustment, error) {
	event := new(MobyOLPManagerSetAumAdjustment)
	if err := _MobyOLPManager.contract.UnpackLog(event, "SetAumAdjustment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerSetCooldownDurationIterator is returned from FilterSetCooldownDuration and is used to iterate over the raw logs and unpacked data for SetCooldownDuration events raised by the MobyOLPManager contract.
type MobyOLPManagerSetCooldownDurationIterator struct {
	Event *MobyOLPManagerSetCooldownDuration // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerSetCooldownDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerSetCooldownDuration)
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
		it.Event = new(MobyOLPManagerSetCooldownDuration)
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
func (it *MobyOLPManagerSetCooldownDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerSetCooldownDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerSetCooldownDuration represents a SetCooldownDuration event raised by the MobyOLPManager contract.
type MobyOLPManagerSetCooldownDuration struct {
	CooldownDuration *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetCooldownDuration is a free log retrieval operation binding the contract event 0xd11d0ebdd6af0dec0dd89018cf0fd5ceb1791ceaf849b900d840af10435abc39.
//
// Solidity: event SetCooldownDuration(uint256 cooldownDuration)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterSetCooldownDuration(opts *bind.FilterOpts) (*MobyOLPManagerSetCooldownDurationIterator, error) {

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "SetCooldownDuration")
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerSetCooldownDurationIterator{contract: _MobyOLPManager.contract, event: "SetCooldownDuration", logs: logs, sub: sub}, nil
}

// WatchSetCooldownDuration is a free log subscription operation binding the contract event 0xd11d0ebdd6af0dec0dd89018cf0fd5ceb1791ceaf849b900d840af10435abc39.
//
// Solidity: event SetCooldownDuration(uint256 cooldownDuration)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchSetCooldownDuration(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerSetCooldownDuration) (event.Subscription, error) {

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "SetCooldownDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerSetCooldownDuration)
				if err := _MobyOLPManager.contract.UnpackLog(event, "SetCooldownDuration", log); err != nil {
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

// ParseSetCooldownDuration is a log parse operation binding the contract event 0xd11d0ebdd6af0dec0dd89018cf0fd5ceb1791ceaf849b900d840af10435abc39.
//
// Solidity: event SetCooldownDuration(uint256 cooldownDuration)
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseSetCooldownDuration(log types.Log) (*MobyOLPManagerSetCooldownDuration, error) {
	event := new(MobyOLPManagerSetCooldownDuration)
	if err := _MobyOLPManager.contract.UnpackLog(event, "SetCooldownDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerSetHandlerIterator is returned from FilterSetHandler and is used to iterate over the raw logs and unpacked data for SetHandler events raised by the MobyOLPManager contract.
type MobyOLPManagerSetHandlerIterator struct {
	Event *MobyOLPManagerSetHandler // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerSetHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerSetHandler)
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
		it.Event = new(MobyOLPManagerSetHandler)
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
func (it *MobyOLPManagerSetHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerSetHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerSetHandler represents a SetHandler event raised by the MobyOLPManager contract.
type MobyOLPManagerSetHandler struct {
	Handler  common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetHandler is a free log retrieval operation binding the contract event 0xd373464a39404e5f98fdb4b152b8ba9a094561e97e5c4b4ea11eb18cd9e6695e.
//
// Solidity: event SetHandler(address indexed handler, bool isActive)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterSetHandler(opts *bind.FilterOpts, handler []common.Address) (*MobyOLPManagerSetHandlerIterator, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "SetHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerSetHandlerIterator{contract: _MobyOLPManager.contract, event: "SetHandler", logs: logs, sub: sub}, nil
}

// WatchSetHandler is a free log subscription operation binding the contract event 0xd373464a39404e5f98fdb4b152b8ba9a094561e97e5c4b4ea11eb18cd9e6695e.
//
// Solidity: event SetHandler(address indexed handler, bool isActive)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchSetHandler(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerSetHandler, handler []common.Address) (event.Subscription, error) {

	var handlerRule []interface{}
	for _, handlerItem := range handler {
		handlerRule = append(handlerRule, handlerItem)
	}

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "SetHandler", handlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerSetHandler)
				if err := _MobyOLPManager.contract.UnpackLog(event, "SetHandler", log); err != nil {
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
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseSetHandler(log types.Log) (*MobyOLPManagerSetHandler, error) {
	event := new(MobyOLPManagerSetHandler)
	if err := _MobyOLPManager.contract.UnpackLog(event, "SetHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerSetInPrivateModeIterator is returned from FilterSetInPrivateMode and is used to iterate over the raw logs and unpacked data for SetInPrivateMode events raised by the MobyOLPManager contract.
type MobyOLPManagerSetInPrivateModeIterator struct {
	Event *MobyOLPManagerSetInPrivateMode // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerSetInPrivateModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerSetInPrivateMode)
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
		it.Event = new(MobyOLPManagerSetInPrivateMode)
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
func (it *MobyOLPManagerSetInPrivateModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerSetInPrivateModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerSetInPrivateMode represents a SetInPrivateMode event raised by the MobyOLPManager contract.
type MobyOLPManagerSetInPrivateMode struct {
	InPrivateMode bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetInPrivateMode is a free log retrieval operation binding the contract event 0x103ae9832f052879845a50017fc76388964d64ed04d8246c05c222d6a2d77268.
//
// Solidity: event SetInPrivateMode(bool inPrivateMode)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterSetInPrivateMode(opts *bind.FilterOpts) (*MobyOLPManagerSetInPrivateModeIterator, error) {

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "SetInPrivateMode")
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerSetInPrivateModeIterator{contract: _MobyOLPManager.contract, event: "SetInPrivateMode", logs: logs, sub: sub}, nil
}

// WatchSetInPrivateMode is a free log subscription operation binding the contract event 0x103ae9832f052879845a50017fc76388964d64ed04d8246c05c222d6a2d77268.
//
// Solidity: event SetInPrivateMode(bool inPrivateMode)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchSetInPrivateMode(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerSetInPrivateMode) (event.Subscription, error) {

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "SetInPrivateMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerSetInPrivateMode)
				if err := _MobyOLPManager.contract.UnpackLog(event, "SetInPrivateMode", log); err != nil {
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

// ParseSetInPrivateMode is a log parse operation binding the contract event 0x103ae9832f052879845a50017fc76388964d64ed04d8246c05c222d6a2d77268.
//
// Solidity: event SetInPrivateMode(bool inPrivateMode)
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseSetInPrivateMode(log types.Log) (*MobyOLPManagerSetInPrivateMode, error) {
	event := new(MobyOLPManagerSetInPrivateMode)
	if err := _MobyOLPManager.contract.UnpackLog(event, "SetInPrivateMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerSetVaultIterator is returned from FilterSetVault and is used to iterate over the raw logs and unpacked data for SetVault events raised by the MobyOLPManager contract.
type MobyOLPManagerSetVaultIterator struct {
	Event *MobyOLPManagerSetVault // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerSetVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerSetVault)
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
		it.Event = new(MobyOLPManagerSetVault)
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
func (it *MobyOLPManagerSetVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerSetVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerSetVault represents a SetVault event raised by the MobyOLPManager contract.
type MobyOLPManagerSetVault struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetVault is a free log retrieval operation binding the contract event 0xd459c7242e23d490831b5676a611c4342d899d28f342d89ae80793e56a930f30.
//
// Solidity: event SetVault(address indexed vault)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterSetVault(opts *bind.FilterOpts, vault []common.Address) (*MobyOLPManagerSetVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "SetVault", vaultRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerSetVaultIterator{contract: _MobyOLPManager.contract, event: "SetVault", logs: logs, sub: sub}, nil
}

// WatchSetVault is a free log subscription operation binding the contract event 0xd459c7242e23d490831b5676a611c4342d899d28f342d89ae80793e56a930f30.
//
// Solidity: event SetVault(address indexed vault)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchSetVault(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerSetVault, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "SetVault", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerSetVault)
				if err := _MobyOLPManager.contract.UnpackLog(event, "SetVault", log); err != nil {
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

// ParseSetVault is a log parse operation binding the contract event 0xd459c7242e23d490831b5676a611c4342d899d28f342d89ae80793e56a930f30.
//
// Solidity: event SetVault(address indexed vault)
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseSetVault(log types.Log) (*MobyOLPManagerSetVault, error) {
	event := new(MobyOLPManagerSetVault)
	if err := _MobyOLPManager.contract.UnpackLog(event, "SetVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerSetVaultPriceFeedIterator is returned from FilterSetVaultPriceFeed and is used to iterate over the raw logs and unpacked data for SetVaultPriceFeed events raised by the MobyOLPManager contract.
type MobyOLPManagerSetVaultPriceFeedIterator struct {
	Event *MobyOLPManagerSetVaultPriceFeed // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerSetVaultPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerSetVaultPriceFeed)
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
		it.Event = new(MobyOLPManagerSetVaultPriceFeed)
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
func (it *MobyOLPManagerSetVaultPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerSetVaultPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerSetVaultPriceFeed represents a SetVaultPriceFeed event raised by the MobyOLPManager contract.
type MobyOLPManagerSetVaultPriceFeed struct {
	VaultPriceFeed common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetVaultPriceFeed is a free log retrieval operation binding the contract event 0x1f9dce6f8c70451ff3c461fd9c19dab1268b83aea65f5410ed6c4196cace264f.
//
// Solidity: event SetVaultPriceFeed(address indexed vaultPriceFeed)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterSetVaultPriceFeed(opts *bind.FilterOpts, vaultPriceFeed []common.Address) (*MobyOLPManagerSetVaultPriceFeedIterator, error) {

	var vaultPriceFeedRule []interface{}
	for _, vaultPriceFeedItem := range vaultPriceFeed {
		vaultPriceFeedRule = append(vaultPriceFeedRule, vaultPriceFeedItem)
	}

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "SetVaultPriceFeed", vaultPriceFeedRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerSetVaultPriceFeedIterator{contract: _MobyOLPManager.contract, event: "SetVaultPriceFeed", logs: logs, sub: sub}, nil
}

// WatchSetVaultPriceFeed is a free log subscription operation binding the contract event 0x1f9dce6f8c70451ff3c461fd9c19dab1268b83aea65f5410ed6c4196cace264f.
//
// Solidity: event SetVaultPriceFeed(address indexed vaultPriceFeed)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchSetVaultPriceFeed(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerSetVaultPriceFeed, vaultPriceFeed []common.Address) (event.Subscription, error) {

	var vaultPriceFeedRule []interface{}
	for _, vaultPriceFeedItem := range vaultPriceFeed {
		vaultPriceFeedRule = append(vaultPriceFeedRule, vaultPriceFeedItem)
	}

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "SetVaultPriceFeed", vaultPriceFeedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerSetVaultPriceFeed)
				if err := _MobyOLPManager.contract.UnpackLog(event, "SetVaultPriceFeed", log); err != nil {
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

// ParseSetVaultPriceFeed is a log parse operation binding the contract event 0x1f9dce6f8c70451ff3c461fd9c19dab1268b83aea65f5410ed6c4196cace264f.
//
// Solidity: event SetVaultPriceFeed(address indexed vaultPriceFeed)
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseSetVaultPriceFeed(log types.Log) (*MobyOLPManagerSetVaultPriceFeed, error) {
	event := new(MobyOLPManagerSetVaultPriceFeed)
	if err := _MobyOLPManager.contract.UnpackLog(event, "SetVaultPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MobyOLPManagerSetVaultUtilsIterator is returned from FilterSetVaultUtils and is used to iterate over the raw logs and unpacked data for SetVaultUtils events raised by the MobyOLPManager contract.
type MobyOLPManagerSetVaultUtilsIterator struct {
	Event *MobyOLPManagerSetVaultUtils // Event containing the contract specifics and raw log

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
func (it *MobyOLPManagerSetVaultUtilsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MobyOLPManagerSetVaultUtils)
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
		it.Event = new(MobyOLPManagerSetVaultUtils)
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
func (it *MobyOLPManagerSetVaultUtilsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MobyOLPManagerSetVaultUtilsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MobyOLPManagerSetVaultUtils represents a SetVaultUtils event raised by the MobyOLPManager contract.
type MobyOLPManagerSetVaultUtils struct {
	VaultUtils common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetVaultUtils is a free log retrieval operation binding the contract event 0x01f1858824776edf07c0cb2b9f2cefdd2c4b3ad503a8636e217f3669866345a4.
//
// Solidity: event SetVaultUtils(address indexed vaultUtils)
func (_MobyOLPManager *MobyOLPManagerFilterer) FilterSetVaultUtils(opts *bind.FilterOpts, vaultUtils []common.Address) (*MobyOLPManagerSetVaultUtilsIterator, error) {

	var vaultUtilsRule []interface{}
	for _, vaultUtilsItem := range vaultUtils {
		vaultUtilsRule = append(vaultUtilsRule, vaultUtilsItem)
	}

	logs, sub, err := _MobyOLPManager.contract.FilterLogs(opts, "SetVaultUtils", vaultUtilsRule)
	if err != nil {
		return nil, err
	}
	return &MobyOLPManagerSetVaultUtilsIterator{contract: _MobyOLPManager.contract, event: "SetVaultUtils", logs: logs, sub: sub}, nil
}

// WatchSetVaultUtils is a free log subscription operation binding the contract event 0x01f1858824776edf07c0cb2b9f2cefdd2c4b3ad503a8636e217f3669866345a4.
//
// Solidity: event SetVaultUtils(address indexed vaultUtils)
func (_MobyOLPManager *MobyOLPManagerFilterer) WatchSetVaultUtils(opts *bind.WatchOpts, sink chan<- *MobyOLPManagerSetVaultUtils, vaultUtils []common.Address) (event.Subscription, error) {

	var vaultUtilsRule []interface{}
	for _, vaultUtilsItem := range vaultUtils {
		vaultUtilsRule = append(vaultUtilsRule, vaultUtilsItem)
	}

	logs, sub, err := _MobyOLPManager.contract.WatchLogs(opts, "SetVaultUtils", vaultUtilsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MobyOLPManagerSetVaultUtils)
				if err := _MobyOLPManager.contract.UnpackLog(event, "SetVaultUtils", log); err != nil {
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

// ParseSetVaultUtils is a log parse operation binding the contract event 0x01f1858824776edf07c0cb2b9f2cefdd2c4b3ad503a8636e217f3669866345a4.
//
// Solidity: event SetVaultUtils(address indexed vaultUtils)
func (_MobyOLPManager *MobyOLPManagerFilterer) ParseSetVaultUtils(log types.Log) (*MobyOLPManagerSetVaultUtils, error) {
	event := new(MobyOLPManagerSetVaultUtils)
	if err := _MobyOLPManager.contract.UnpackLog(event, "SetVaultUtils", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
