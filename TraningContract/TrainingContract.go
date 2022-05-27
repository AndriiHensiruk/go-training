// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TraningContract

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

// TrainingContractLocation is an auto generated low-level Go binding around an user-defined struct.
type TrainingContractLocation struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	Province     string
	CountryCode  string
	PostalCode   string
}

// TrainingContractMetaData contains all meta data concerning the TrainingContract contract.
var TrainingContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"addressLine1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"addressLine2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"city\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"province\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"countryCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"postalCode\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structTrainingContract.Location\",\"name\":\"l\",\"type\":\"tuple\"}],\"name\":\"LOCATION\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MONEY_RECEIVED\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"giveMeMoney\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"addressLine1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"addressLine2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"city\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"province\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"countryCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"postalCode\",\"type\":\"string\"}],\"internalType\":\"structTrainingContract.Location\",\"name\":\"location_\",\"type\":\"tuple\"}],\"name\":\"setLocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TrainingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TrainingContractMetaData.ABI instead.
var TrainingContractABI = TrainingContractMetaData.ABI

// TrainingContract is an auto generated Go binding around an Ethereum contract.
type TrainingContract struct {
	TrainingContractCaller     // Read-only binding to the contract
	TrainingContractTransactor // Write-only binding to the contract
	TrainingContractFilterer   // Log filterer for contract events
}

// TrainingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrainingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrainingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrainingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrainingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrainingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrainingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrainingContractSession struct {
	Contract     *TrainingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrainingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrainingContractCallerSession struct {
	Contract *TrainingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TrainingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrainingContractTransactorSession struct {
	Contract     *TrainingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TrainingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrainingContractRaw struct {
	Contract *TrainingContract // Generic contract binding to access the raw methods on
}

// TrainingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrainingContractCallerRaw struct {
	Contract *TrainingContractCaller // Generic read-only contract binding to access the raw methods on
}

// TrainingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrainingContractTransactorRaw struct {
	Contract *TrainingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrainingContract creates a new instance of TrainingContract, bound to a specific deployed contract.
func NewTrainingContract(address common.Address, backend bind.ContractBackend) (*TrainingContract, error) {
	contract, err := bindTrainingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrainingContract{TrainingContractCaller: TrainingContractCaller{contract: contract}, TrainingContractTransactor: TrainingContractTransactor{contract: contract}, TrainingContractFilterer: TrainingContractFilterer{contract: contract}}, nil
}

// NewTrainingContractCaller creates a new read-only instance of TrainingContract, bound to a specific deployed contract.
func NewTrainingContractCaller(address common.Address, caller bind.ContractCaller) (*TrainingContractCaller, error) {
	contract, err := bindTrainingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrainingContractCaller{contract: contract}, nil
}

// NewTrainingContractTransactor creates a new write-only instance of TrainingContract, bound to a specific deployed contract.
func NewTrainingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TrainingContractTransactor, error) {
	contract, err := bindTrainingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrainingContractTransactor{contract: contract}, nil
}

// NewTrainingContractFilterer creates a new log filterer instance of TrainingContract, bound to a specific deployed contract.
func NewTrainingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TrainingContractFilterer, error) {
	contract, err := bindTrainingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrainingContractFilterer{contract: contract}, nil
}

// bindTrainingContract binds a generic wrapper to an already deployed contract.
func bindTrainingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrainingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrainingContract *TrainingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrainingContract.Contract.TrainingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrainingContract *TrainingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrainingContract.Contract.TrainingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrainingContract *TrainingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrainingContract.Contract.TrainingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrainingContract *TrainingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrainingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrainingContract *TrainingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrainingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrainingContract *TrainingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrainingContract.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_TrainingContract *TrainingContractCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TrainingContract.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_TrainingContract *TrainingContractSession) GetName() (string, error) {
	return _TrainingContract.Contract.GetName(&_TrainingContract.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_TrainingContract *TrainingContractCallerSession) GetName() (string, error) {
	return _TrainingContract.Contract.GetName(&_TrainingContract.CallOpts)
}

// GetLocation is a paid mutator transaction binding the contract method 0xce2ce3fc.
//
// Solidity: function getLocation() returns()
func (_TrainingContract *TrainingContractTransactor) GetLocation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrainingContract.contract.Transact(opts, "getLocation")
}

// GetLocation is a paid mutator transaction binding the contract method 0xce2ce3fc.
//
// Solidity: function getLocation() returns()
func (_TrainingContract *TrainingContractSession) GetLocation() (*types.Transaction, error) {
	return _TrainingContract.Contract.GetLocation(&_TrainingContract.TransactOpts)
}

// GetLocation is a paid mutator transaction binding the contract method 0xce2ce3fc.
//
// Solidity: function getLocation() returns()
func (_TrainingContract *TrainingContractTransactorSession) GetLocation() (*types.Transaction, error) {
	return _TrainingContract.Contract.GetLocation(&_TrainingContract.TransactOpts)
}

// GiveMeMoney is a paid mutator transaction binding the contract method 0x7aa1469f.
//
// Solidity: function giveMeMoney() payable returns()
func (_TrainingContract *TrainingContractTransactor) GiveMeMoney(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrainingContract.contract.Transact(opts, "giveMeMoney")
}

// GiveMeMoney is a paid mutator transaction binding the contract method 0x7aa1469f.
//
// Solidity: function giveMeMoney() payable returns()
func (_TrainingContract *TrainingContractSession) GiveMeMoney() (*types.Transaction, error) {
	return _TrainingContract.Contract.GiveMeMoney(&_TrainingContract.TransactOpts)
}

// GiveMeMoney is a paid mutator transaction binding the contract method 0x7aa1469f.
//
// Solidity: function giveMeMoney() payable returns()
func (_TrainingContract *TrainingContractTransactorSession) GiveMeMoney() (*types.Transaction, error) {
	return _TrainingContract.Contract.GiveMeMoney(&_TrainingContract.TransactOpts)
}

// SetLocation is a paid mutator transaction binding the contract method 0x0f86b760.
//
// Solidity: function setLocation((string,string,string,string,string,string) location_) returns()
func (_TrainingContract *TrainingContractTransactor) SetLocation(opts *bind.TransactOpts, location_ TrainingContractLocation) (*types.Transaction, error) {
	return _TrainingContract.contract.Transact(opts, "setLocation", location_)
}

// SetLocation is a paid mutator transaction binding the contract method 0x0f86b760.
//
// Solidity: function setLocation((string,string,string,string,string,string) location_) returns()
func (_TrainingContract *TrainingContractSession) SetLocation(location_ TrainingContractLocation) (*types.Transaction, error) {
	return _TrainingContract.Contract.SetLocation(&_TrainingContract.TransactOpts, location_)
}

// SetLocation is a paid mutator transaction binding the contract method 0x0f86b760.
//
// Solidity: function setLocation((string,string,string,string,string,string) location_) returns()
func (_TrainingContract *TrainingContractTransactorSession) SetLocation(location_ TrainingContractLocation) (*types.Transaction, error) {
	return _TrainingContract.Contract.SetLocation(&_TrainingContract.TransactOpts, location_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_TrainingContract *TrainingContractTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _TrainingContract.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_TrainingContract *TrainingContractSession) SetName(name_ string) (*types.Transaction, error) {
	return _TrainingContract.Contract.SetName(&_TrainingContract.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_TrainingContract *TrainingContractTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _TrainingContract.Contract.SetName(&_TrainingContract.TransactOpts, name_)
}

// TrainingContractLOCATIONIterator is returned from FilterLOCATION and is used to iterate over the raw logs and unpacked data for LOCATION events raised by the TrainingContract contract.
type TrainingContractLOCATIONIterator struct {
	Event *TrainingContractLOCATION // Event containing the contract specifics and raw log

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
func (it *TrainingContractLOCATIONIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrainingContractLOCATION)
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
		it.Event = new(TrainingContractLOCATION)
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
func (it *TrainingContractLOCATIONIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrainingContractLOCATIONIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrainingContractLOCATION represents a LOCATION event raised by the TrainingContract contract.
type TrainingContractLOCATION struct {
	L   TrainingContractLocation
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLOCATION is a free log retrieval operation binding the contract event 0x9bce99fb72b0420725cee6e006a7cc95f4a0e0dfce21fbdc74bb9b089fae7c7b.
//
// Solidity: event LOCATION((string,string,string,string,string,string) l)
func (_TrainingContract *TrainingContractFilterer) FilterLOCATION(opts *bind.FilterOpts) (*TrainingContractLOCATIONIterator, error) {

	logs, sub, err := _TrainingContract.contract.FilterLogs(opts, "LOCATION")
	if err != nil {
		return nil, err
	}
	return &TrainingContractLOCATIONIterator{contract: _TrainingContract.contract, event: "LOCATION", logs: logs, sub: sub}, nil
}

// WatchLOCATION is a free log subscription operation binding the contract event 0x9bce99fb72b0420725cee6e006a7cc95f4a0e0dfce21fbdc74bb9b089fae7c7b.
//
// Solidity: event LOCATION((string,string,string,string,string,string) l)
func (_TrainingContract *TrainingContractFilterer) WatchLOCATION(opts *bind.WatchOpts, sink chan<- *TrainingContractLOCATION) (event.Subscription, error) {

	logs, sub, err := _TrainingContract.contract.WatchLogs(opts, "LOCATION")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrainingContractLOCATION)
				if err := _TrainingContract.contract.UnpackLog(event, "LOCATION", log); err != nil {
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

// ParseLOCATION is a log parse operation binding the contract event 0x9bce99fb72b0420725cee6e006a7cc95f4a0e0dfce21fbdc74bb9b089fae7c7b.
//
// Solidity: event LOCATION((string,string,string,string,string,string) l)
func (_TrainingContract *TrainingContractFilterer) ParseLOCATION(log types.Log) (*TrainingContractLOCATION, error) {
	event := new(TrainingContractLOCATION)
	if err := _TrainingContract.contract.UnpackLog(event, "LOCATION", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrainingContractMONEYRECEIVEDIterator is returned from FilterMONEYRECEIVED and is used to iterate over the raw logs and unpacked data for MONEYRECEIVED events raised by the TrainingContract contract.
type TrainingContractMONEYRECEIVEDIterator struct {
	Event *TrainingContractMONEYRECEIVED // Event containing the contract specifics and raw log

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
func (it *TrainingContractMONEYRECEIVEDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrainingContractMONEYRECEIVED)
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
		it.Event = new(TrainingContractMONEYRECEIVED)
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
func (it *TrainingContractMONEYRECEIVEDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrainingContractMONEYRECEIVEDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrainingContractMONEYRECEIVED represents a MONEYRECEIVED event raised by the TrainingContract contract.
type TrainingContractMONEYRECEIVED struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMONEYRECEIVED is a free log retrieval operation binding the contract event 0x6ab7f172b6c2a4245900b6bbfe74c4aecfa72d060b7cc6fff64d7300ab36f941.
//
// Solidity: event MONEY_RECEIVED(address arg0, uint256 arg1)
func (_TrainingContract *TrainingContractFilterer) FilterMONEYRECEIVED(opts *bind.FilterOpts) (*TrainingContractMONEYRECEIVEDIterator, error) {

	logs, sub, err := _TrainingContract.contract.FilterLogs(opts, "MONEY_RECEIVED")
	if err != nil {
		return nil, err
	}
	return &TrainingContractMONEYRECEIVEDIterator{contract: _TrainingContract.contract, event: "MONEY_RECEIVED", logs: logs, sub: sub}, nil
}

// WatchMONEYRECEIVED is a free log subscription operation binding the contract event 0x6ab7f172b6c2a4245900b6bbfe74c4aecfa72d060b7cc6fff64d7300ab36f941.
//
// Solidity: event MONEY_RECEIVED(address arg0, uint256 arg1)
func (_TrainingContract *TrainingContractFilterer) WatchMONEYRECEIVED(opts *bind.WatchOpts, sink chan<- *TrainingContractMONEYRECEIVED) (event.Subscription, error) {

	logs, sub, err := _TrainingContract.contract.WatchLogs(opts, "MONEY_RECEIVED")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrainingContractMONEYRECEIVED)
				if err := _TrainingContract.contract.UnpackLog(event, "MONEY_RECEIVED", log); err != nil {
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

// ParseMONEYRECEIVED is a log parse operation binding the contract event 0x6ab7f172b6c2a4245900b6bbfe74c4aecfa72d060b7cc6fff64d7300ab36f941.
//
// Solidity: event MONEY_RECEIVED(address arg0, uint256 arg1)
func (_TrainingContract *TrainingContractFilterer) ParseMONEYRECEIVED(log types.Log) (*TrainingContractMONEYRECEIVED, error) {
	event := new(TrainingContractMONEYRECEIVED)
	if err := _TrainingContract.contract.UnpackLog(event, "MONEY_RECEIVED", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
