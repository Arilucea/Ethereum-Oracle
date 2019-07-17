// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracleGo

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

// OracleGoABI is the input ABI used to generate the binding from.
const OracleGoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"changePrize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes\"},{\"name\":\"_value\",\"type\":\"bool\"}],\"name\":\"answer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"changeWhiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_message\",\"type\":\"string\"},{\"name\":\"_type\",\"type\":\"string\"}],\"name\":\"makePetition\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes\"},{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"answer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"answer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"mType\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Petition\",\"type\":\"event\"}]"

// OracleGo is an auto generated Go binding around an Ethereum contract.
type OracleGo struct {
	OracleGoCaller     // Read-only binding to the contract
	OracleGoTransactor // Write-only binding to the contract
	OracleGoFilterer   // Log filterer for contract events
}

// OracleGoCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleGoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleGoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleGoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleGoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleGoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleGoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleGoSession struct {
	Contract     *OracleGo         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleGoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleGoCallerSession struct {
	Contract *OracleGoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OracleGoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleGoTransactorSession struct {
	Contract     *OracleGoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OracleGoRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleGoRaw struct {
	Contract *OracleGo // Generic contract binding to access the raw methods on
}

// OracleGoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleGoCallerRaw struct {
	Contract *OracleGoCaller // Generic read-only contract binding to access the raw methods on
}

// OracleGoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleGoTransactorRaw struct {
	Contract *OracleGoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleGo creates a new instance of OracleGo, bound to a specific deployed contract.
func NewOracleGo(address common.Address, backend bind.ContractBackend) (*OracleGo, error) {
	contract, err := bindOracleGo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleGo{OracleGoCaller: OracleGoCaller{contract: contract}, OracleGoTransactor: OracleGoTransactor{contract: contract}, OracleGoFilterer: OracleGoFilterer{contract: contract}}, nil
}

// NewOracleGoCaller creates a new read-only instance of OracleGo, bound to a specific deployed contract.
func NewOracleGoCaller(address common.Address, caller bind.ContractCaller) (*OracleGoCaller, error) {
	contract, err := bindOracleGo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleGoCaller{contract: contract}, nil
}

// NewOracleGoTransactor creates a new write-only instance of OracleGo, bound to a specific deployed contract.
func NewOracleGoTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleGoTransactor, error) {
	contract, err := bindOracleGo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleGoTransactor{contract: contract}, nil
}

// NewOracleGoFilterer creates a new log filterer instance of OracleGo, bound to a specific deployed contract.
func NewOracleGoFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleGoFilterer, error) {
	contract, err := bindOracleGo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleGoFilterer{contract: contract}, nil
}

// bindOracleGo binds a generic wrapper to an already deployed contract.
func bindOracleGo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleGoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleGo *OracleGoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleGo.Contract.OracleGoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleGo *OracleGoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleGo.Contract.OracleGoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleGo *OracleGoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleGo.Contract.OracleGoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleGo *OracleGoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleGo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleGo *OracleGoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleGo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleGo *OracleGoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleGo.Contract.contract.Transact(opts, method, params...)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_OracleGo *OracleGoCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleGo.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_OracleGo *OracleGoSession) Price() (*big.Int, error) {
	return _OracleGo.Contract.Price(&_OracleGo.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_OracleGo *OracleGoCallerSession) Price() (*big.Int, error) {
	return _OracleGo.Contract.Price(&_OracleGo.CallOpts)
}

// Answer is a paid mutator transaction binding the contract method 0x30655be8.
//
// Solidity: function answer(bytes _id, bool _value) returns()
func (_OracleGo *OracleGoTransactor) Answer(opts *bind.TransactOpts, _id []byte, _value bool) (*types.Transaction, error) {
	return _OracleGo.contract.Transact(opts, "answer", _id, _value)
}

// Answer is a paid mutator transaction binding the contract method 0x30655be8.
//
// Solidity: function answer(bytes _id, bool _value) returns()
func (_OracleGo *OracleGoSession) Answer(_id []byte, _value bool) (*types.Transaction, error) {
	return _OracleGo.Contract.Answer(&_OracleGo.TransactOpts, _id, _value)
}

// Answer is a paid mutator transaction binding the contract method 0x30655be8.
//
// Solidity: function answer(bytes _id, bool _value) returns()
func (_OracleGo *OracleGoTransactorSession) Answer(_id []byte, _value bool) (*types.Transaction, error) {
	return _OracleGo.Contract.Answer(&_OracleGo.TransactOpts, _id, _value)
}

// Answer0 is a paid mutator transaction binding the contract method 0x1e21ee3c.
//
// Solidity: function answer0(bytes _id, string _value) returns()
func (_OracleGo *OracleGoTransactor) Answer0(opts *bind.TransactOpts, _id []byte, _value string) (*types.Transaction, error) {
	return _OracleGo.contract.Transact(opts, "answer0", _id, _value)
}

// Answer0 is a paid mutator transaction binding the contract method 0x1e21ee3c.
//
// Solidity: function answer0(bytes _id, string _value) returns()
func (_OracleGo *OracleGoSession) Answer0(_id []byte, _value string) (*types.Transaction, error) {
	return _OracleGo.Contract.Answer0(&_OracleGo.TransactOpts, _id, _value)
}

// Answer0 is a paid mutator transaction binding the contract method 0x1e21ee3c.
//
// Solidity: function answer0(bytes _id, string _value) returns()
func (_OracleGo *OracleGoTransactorSession) Answer0(_id []byte, _value string) (*types.Transaction, error) {
	return _OracleGo.Contract.Answer0(&_OracleGo.TransactOpts, _id, _value)
}

// Answer1 is a paid mutator transaction binding the contract method 0x62a4950c.
//
// Solidity: function answer1(bytes _id, uint256 _value) returns()
func (_OracleGo *OracleGoTransactor) Answer1(opts *bind.TransactOpts, _id []byte, _value *big.Int) (*types.Transaction, error) {
	return _OracleGo.contract.Transact(opts, "answer1", _id, _value)
}

// Answer1 is a paid mutator transaction binding the contract method 0x62a4950c.
//
// Solidity: function answer1(bytes _id, uint256 _value) returns()
func (_OracleGo *OracleGoSession) Answer1(_id []byte, _value *big.Int) (*types.Transaction, error) {
	return _OracleGo.Contract.Answer1(&_OracleGo.TransactOpts, _id, _value)
}

// Answer1 is a paid mutator transaction binding the contract method 0x62a4950c.
//
// Solidity: function answer1(bytes _id, uint256 _value) returns()
func (_OracleGo *OracleGoTransactorSession) Answer1(_id []byte, _value *big.Int) (*types.Transaction, error) {
	return _OracleGo.Contract.Answer1(&_OracleGo.TransactOpts, _id, _value)
}

// ChangePrize is a paid mutator transaction binding the contract method 0x074cb69f.
//
// Solidity: function changePrize(uint256 _newPrice) returns()
func (_OracleGo *OracleGoTransactor) ChangePrize(opts *bind.TransactOpts, _newPrice *big.Int) (*types.Transaction, error) {
	return _OracleGo.contract.Transact(opts, "changePrize", _newPrice)
}

// ChangePrize is a paid mutator transaction binding the contract method 0x074cb69f.
//
// Solidity: function changePrize(uint256 _newPrice) returns()
func (_OracleGo *OracleGoSession) ChangePrize(_newPrice *big.Int) (*types.Transaction, error) {
	return _OracleGo.Contract.ChangePrize(&_OracleGo.TransactOpts, _newPrice)
}

// ChangePrize is a paid mutator transaction binding the contract method 0x074cb69f.
//
// Solidity: function changePrize(uint256 _newPrice) returns()
func (_OracleGo *OracleGoTransactorSession) ChangePrize(_newPrice *big.Int) (*types.Transaction, error) {
	return _OracleGo.Contract.ChangePrize(&_OracleGo.TransactOpts, _newPrice)
}

// ChangeWhiteList is a paid mutator transaction binding the contract method 0x39b417b7.
//
// Solidity: function changeWhiteList(address _address, bool _state) returns()
func (_OracleGo *OracleGoTransactor) ChangeWhiteList(opts *bind.TransactOpts, _address common.Address, _state bool) (*types.Transaction, error) {
	return _OracleGo.contract.Transact(opts, "changeWhiteList", _address, _state)
}

// ChangeWhiteList is a paid mutator transaction binding the contract method 0x39b417b7.
//
// Solidity: function changeWhiteList(address _address, bool _state) returns()
func (_OracleGo *OracleGoSession) ChangeWhiteList(_address common.Address, _state bool) (*types.Transaction, error) {
	return _OracleGo.Contract.ChangeWhiteList(&_OracleGo.TransactOpts, _address, _state)
}

// ChangeWhiteList is a paid mutator transaction binding the contract method 0x39b417b7.
//
// Solidity: function changeWhiteList(address _address, bool _state) returns()
func (_OracleGo *OracleGoTransactorSession) ChangeWhiteList(_address common.Address, _state bool) (*types.Transaction, error) {
	return _OracleGo.Contract.ChangeWhiteList(&_OracleGo.TransactOpts, _address, _state)
}

// GetEth is a paid mutator transaction binding the contract method 0xcb05b93e.
//
// Solidity: function getEth() returns()
func (_OracleGo *OracleGoTransactor) GetEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleGo.contract.Transact(opts, "getEth")
}

// GetEth is a paid mutator transaction binding the contract method 0xcb05b93e.
//
// Solidity: function getEth() returns()
func (_OracleGo *OracleGoSession) GetEth() (*types.Transaction, error) {
	return _OracleGo.Contract.GetEth(&_OracleGo.TransactOpts)
}

// GetEth is a paid mutator transaction binding the contract method 0xcb05b93e.
//
// Solidity: function getEth() returns()
func (_OracleGo *OracleGoTransactorSession) GetEth() (*types.Transaction, error) {
	return _OracleGo.Contract.GetEth(&_OracleGo.TransactOpts)
}

// MakePetition is a paid mutator transaction binding the contract method 0x4d3044fe.
//
// Solidity: function makePetition(string _message, string _type) returns()
func (_OracleGo *OracleGoTransactor) MakePetition(opts *bind.TransactOpts, _message string, _type string) (*types.Transaction, error) {
	return _OracleGo.contract.Transact(opts, "makePetition", _message, _type)
}

// MakePetition is a paid mutator transaction binding the contract method 0x4d3044fe.
//
// Solidity: function makePetition(string _message, string _type) returns()
func (_OracleGo *OracleGoSession) MakePetition(_message string, _type string) (*types.Transaction, error) {
	return _OracleGo.Contract.MakePetition(&_OracleGo.TransactOpts, _message, _type)
}

// MakePetition is a paid mutator transaction binding the contract method 0x4d3044fe.
//
// Solidity: function makePetition(string _message, string _type) returns()
func (_OracleGo *OracleGoTransactorSession) MakePetition(_message string, _type string) (*types.Transaction, error) {
	return _OracleGo.Contract.MakePetition(&_OracleGo.TransactOpts, _message, _type)
}

// OracleGoPetitionIterator is returned from FilterPetition and is used to iterate over the raw logs and unpacked data for Petition events raised by the OracleGo contract.
type OracleGoPetitionIterator struct {
	Event *OracleGoPetition // Event containing the contract specifics and raw log

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
func (it *OracleGoPetitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleGoPetition)
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
		it.Event = new(OracleGoPetition)
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
func (it *OracleGoPetitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleGoPetitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleGoPetition represents a Petition event raised by the OracleGo contract.
type OracleGoPetition struct {
	Id      []byte
	MType   string
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPetition is a free log retrieval operation binding the contract event 0xbd8b7313fd1af282e2644201bb3dd87abc00cf6d8e512d120ae3be15bc2066ab.
//
// Solidity: event Petition(bytes id, string mType, string message)
func (_OracleGo *OracleGoFilterer) FilterPetition(opts *bind.FilterOpts) (*OracleGoPetitionIterator, error) {

	logs, sub, err := _OracleGo.contract.FilterLogs(opts, "Petition")
	if err != nil {
		return nil, err
	}
	return &OracleGoPetitionIterator{contract: _OracleGo.contract, event: "Petition", logs: logs, sub: sub}, nil
}

// WatchPetition is a free log subscription operation binding the contract event 0xbd8b7313fd1af282e2644201bb3dd87abc00cf6d8e512d120ae3be15bc2066ab.
//
// Solidity: event Petition(bytes id, string mType, string message)
func (_OracleGo *OracleGoFilterer) WatchPetition(opts *bind.WatchOpts, sink chan<- *OracleGoPetition) (event.Subscription, error) {

	logs, sub, err := _OracleGo.contract.WatchLogs(opts, "Petition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleGoPetition)
				if err := _OracleGo.contract.UnpackLog(event, "Petition", log); err != nil {
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

// ParsePetition is a log parse operation binding the contract event 0xbd8b7313fd1af282e2644201bb3dd87abc00cf6d8e512d120ae3be15bc2066ab.
//
// Solidity: event Petition(bytes id, string mType, string message)
func (_OracleGo *OracleGoFilterer) ParsePetition(log types.Log) (*OracleGoPetition, error) {
	event := new(OracleGoPetition)
	if err := _OracleGo.contract.UnpackLog(event, "Petition", log); err != nil {
		return nil, err
	}
	return event, nil
}
