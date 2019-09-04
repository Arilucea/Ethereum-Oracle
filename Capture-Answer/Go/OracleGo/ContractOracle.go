// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"changePrize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"changeWhiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_message\",\"type\":\"string\"},{\"name\":\"_type\",\"type\":\"string\"}],\"name\":\"makePetition\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"mType\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Petition\",\"type\":\"event\"}]"

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Oracle *OracleCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Oracle *OracleSession) Price() (*big.Int, error) {
	return _Oracle.Contract.Price(&_Oracle.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Oracle *OracleCallerSession) Price() (*big.Int, error) {
	return _Oracle.Contract.Price(&_Oracle.CallOpts)
}

// ChangePrize is a paid mutator transaction binding the contract method 0x074cb69f.
//
// Solidity: function changePrize(uint256 _newPrice) returns()
func (_Oracle *OracleTransactor) ChangePrize(opts *bind.TransactOpts, _newPrice *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "changePrize", _newPrice)
}

// ChangePrize is a paid mutator transaction binding the contract method 0x074cb69f.
//
// Solidity: function changePrize(uint256 _newPrice) returns()
func (_Oracle *OracleSession) ChangePrize(_newPrice *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ChangePrize(&_Oracle.TransactOpts, _newPrice)
}

// ChangePrize is a paid mutator transaction binding the contract method 0x074cb69f.
//
// Solidity: function changePrize(uint256 _newPrice) returns()
func (_Oracle *OracleTransactorSession) ChangePrize(_newPrice *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.ChangePrize(&_Oracle.TransactOpts, _newPrice)
}

// ChangeWhiteList is a paid mutator transaction binding the contract method 0x39b417b7.
//
// Solidity: function changeWhiteList(address _address, bool _state) returns()
func (_Oracle *OracleTransactor) ChangeWhiteList(opts *bind.TransactOpts, _address common.Address, _state bool) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "changeWhiteList", _address, _state)
}

// ChangeWhiteList is a paid mutator transaction binding the contract method 0x39b417b7.
//
// Solidity: function changeWhiteList(address _address, bool _state) returns()
func (_Oracle *OracleSession) ChangeWhiteList(_address common.Address, _state bool) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeWhiteList(&_Oracle.TransactOpts, _address, _state)
}

// ChangeWhiteList is a paid mutator transaction binding the contract method 0x39b417b7.
//
// Solidity: function changeWhiteList(address _address, bool _state) returns()
func (_Oracle *OracleTransactorSession) ChangeWhiteList(_address common.Address, _state bool) (*types.Transaction, error) {
	return _Oracle.Contract.ChangeWhiteList(&_Oracle.TransactOpts, _address, _state)
}

// GetEth is a paid mutator transaction binding the contract method 0xcb05b93e.
//
// Solidity: function getEth() returns()
func (_Oracle *OracleTransactor) GetEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "getEth")
}

// GetEth is a paid mutator transaction binding the contract method 0xcb05b93e.
//
// Solidity: function getEth() returns()
func (_Oracle *OracleSession) GetEth() (*types.Transaction, error) {
	return _Oracle.Contract.GetEth(&_Oracle.TransactOpts)
}

// GetEth is a paid mutator transaction binding the contract method 0xcb05b93e.
//
// Solidity: function getEth() returns()
func (_Oracle *OracleTransactorSession) GetEth() (*types.Transaction, error) {
	return _Oracle.Contract.GetEth(&_Oracle.TransactOpts)
}

// MakePetition is a paid mutator transaction binding the contract method 0x4d3044fe.
//
// Solidity: function makePetition(string _message, string _type) returns()
func (_Oracle *OracleTransactor) MakePetition(opts *bind.TransactOpts, _message string, _type string) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "makePetition", _message, _type)
}

// MakePetition is a paid mutator transaction binding the contract method 0x4d3044fe.
//
// Solidity: function makePetition(string _message, string _type) returns()
func (_Oracle *OracleSession) MakePetition(_message string, _type string) (*types.Transaction, error) {
	return _Oracle.Contract.MakePetition(&_Oracle.TransactOpts, _message, _type)
}

// MakePetition is a paid mutator transaction binding the contract method 0x4d3044fe.
//
// Solidity: function makePetition(string _message, string _type) returns()
func (_Oracle *OracleTransactorSession) MakePetition(_message string, _type string) (*types.Transaction, error) {
	return _Oracle.Contract.MakePetition(&_Oracle.TransactOpts, _message, _type)
}

// OraclePetitionIterator is returned from FilterPetition and is used to iterate over the raw logs and unpacked data for Petition events raised by the Oracle contract.
type OraclePetitionIterator struct {
	Event *OraclePetition // Event containing the contract specifics and raw log

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
func (it *OraclePetitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OraclePetition)
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
		it.Event = new(OraclePetition)
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
func (it *OraclePetitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OraclePetitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OraclePetition represents a Petition event raised by the Oracle contract.
type OraclePetition struct {
	Caller  common.Address
	MType   string
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPetition is a free log retrieval operation binding the contract event 0xc659bc090075346ac7a96402907b87a7f9b4292b51589a09794a3e6e649d1e2a.
//
// Solidity: event Petition(address caller, string mType, string message)
func (_Oracle *OracleFilterer) FilterPetition(opts *bind.FilterOpts) (*OraclePetitionIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Petition")
	if err != nil {
		return nil, err
	}
	return &OraclePetitionIterator{contract: _Oracle.contract, event: "Petition", logs: logs, sub: sub}, nil
}

// WatchPetition is a free log subscription operation binding the contract event 0xc659bc090075346ac7a96402907b87a7f9b4292b51589a09794a3e6e649d1e2a.
//
// Solidity: event Petition(address caller, string mType, string message)
func (_Oracle *OracleFilterer) WatchPetition(opts *bind.WatchOpts, sink chan<- *OraclePetition) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Petition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OraclePetition)
				if err := _Oracle.contract.UnpackLog(event, "Petition", log); err != nil {
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

// ParsePetition is a log parse operation binding the contract event 0xc659bc090075346ac7a96402907b87a7f9b4292b51589a09794a3e6e649d1e2a.
//
// Solidity: event Petition(address caller, string mType, string message)
func (_Oracle *OracleFilterer) ParsePetition(log types.Log) (*OraclePetition, error) {
	event := new(OraclePetition)
	if err := _Oracle.contract.UnpackLog(event, "Petition", log); err != nil {
		return nil, err
	}
	return event, nil
}
