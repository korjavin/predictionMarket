// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TinyOracleMetaData contains all meta data concerning the TinyOracle contract.
var TinyOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_trustedWallets\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"}],\"name\":\"ResultSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"isTrustedWallet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"name\":\"setResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trustedWallets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f8060016101000a81548160ff02191690831515021790555034801562000029575f80fd5b50604051620009503803806200095083398181016040528101906200004f9190620002c8565b5f5b8151811015620000e2576001805f84848151811062000075576200007462000317565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508080620000d9906200037a565b91505062000051565b5050620003c6565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200014782620000ff565b810181811067ffffffffffffffff821117156200016957620001686200010f565b5b80604052505050565b5f6200017d620000ea565b90506200018b82826200013c565b919050565b5f67ffffffffffffffff821115620001ad57620001ac6200010f565b5b602082029050602081019050919050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620001ed82620001c2565b9050919050565b620001ff81620001e1565b81146200020a575f80fd5b50565b5f815190506200021d81620001f4565b92915050565b5f62000239620002338462000190565b62000172565b905080838252602082019050602084028301858111156200025f576200025e620001be565b5b835b818110156200028c57806200027788826200020d565b84526020840193505060208101905062000261565b5050509392505050565b5f82601f830112620002ad57620002ac620000fb565b5b8151620002bf84826020860162000223565b91505092915050565b5f60208284031215620002e057620002df620000f3565b5b5f82015167ffffffffffffffff8111156200030057620002ff620000f7565b5b6200030e8482850162000296565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f819050919050565b5f620003868262000371565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620003bb57620003ba62000344565b5b600182019050919050565b61057c80620003d45f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c8063789f5abe1461004e578063de2927891461007e578063f4b4dc2e1461009d578063ff948fe5146100cd575b5f80fd5b6100686004803603810190610063919061033e565b6100fd565b6040516100759190610383565b60405180910390f35b61008661014f565b60405161009492919061039c565b60405180910390f35b6100b760048036038101906100b291906103ed565b610175565b6040516100c49190610383565b60405180910390f35b6100e760048036038101906100e2919061033e565b6102c3565b6040516100f49190610383565b60405180910390f35b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f805f8054906101000a900460ff165f60019054906101000a900460ff16915091509091565b5f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166101ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f690610498565b60405180910390fd5b5f60019054906101000a900460ff161561024e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024590610500565b60405180910390fd5b815f806101000a81548160ff02191690831515021790555060015f60016101000a81548160ff0219169083151502179055508115157ff08e01198b2d57b28d82ce475fc894dcff0ca9676ac6f8be2b766a732e054a17336040516102b2919061052d565b60405180910390a260019050919050565b6001602052805f5260405f205f915054906101000a900460ff1681565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61030d826102e4565b9050919050565b61031d81610303565b8114610327575f80fd5b50565b5f8135905061033881610314565b92915050565b5f60208284031215610353576103526102e0565b5b5f6103608482850161032a565b91505092915050565b5f8115159050919050565b61037d81610369565b82525050565b5f6020820190506103965f830184610374565b92915050565b5f6040820190506103af5f830185610374565b6103bc6020830184610374565b9392505050565b6103cc81610369565b81146103d6575f80fd5b50565b5f813590506103e7816103c3565b92915050565b5f60208284031215610402576104016102e0565b5b5f61040f848285016103d9565b91505092915050565b5f82825260208201905092915050565b7f4f6e6c7920747275737465642077616c6c6574732063616e20736574207468655f8201527f20726573756c7400000000000000000000000000000000000000000000000000602082015250565b5f610482602783610418565b915061048d82610428565b604082019050919050565b5f6020820190508181035f8301526104af81610476565b9050919050565b7f526573756c742068617320616c7265616479206265656e2073657400000000005f82015250565b5f6104ea601b83610418565b91506104f5826104b6565b602082019050919050565b5f6020820190508181035f830152610517816104de565b9050919050565b61052781610303565b82525050565b5f6020820190506105405f83018461051e565b9291505056fea2646970667358221220cbd201dcab5001deffd0aba213a1ec23941109c64c65554963f321a622cb4e5264736f6c63430008140033",
}

// TinyOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use TinyOracleMetaData.ABI instead.
var TinyOracleABI = TinyOracleMetaData.ABI

// TinyOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TinyOracleMetaData.Bin instead.
var TinyOracleBin = TinyOracleMetaData.Bin

// DeployTinyOracle deploys a new Ethereum contract, binding an instance of TinyOracle to it.
func DeployTinyOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _trustedWallets []common.Address) (common.Address, *types.Transaction, *TinyOracle, error) {
	parsed, err := TinyOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TinyOracleBin), backend, _trustedWallets)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TinyOracle{TinyOracleCaller: TinyOracleCaller{contract: contract}, TinyOracleTransactor: TinyOracleTransactor{contract: contract}, TinyOracleFilterer: TinyOracleFilterer{contract: contract}}, nil
}

// TinyOracle is an auto generated Go binding around an Ethereum contract.
type TinyOracle struct {
	TinyOracleCaller     // Read-only binding to the contract
	TinyOracleTransactor // Write-only binding to the contract
	TinyOracleFilterer   // Log filterer for contract events
}

// TinyOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type TinyOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TinyOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TinyOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TinyOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TinyOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TinyOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TinyOracleSession struct {
	Contract     *TinyOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TinyOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TinyOracleCallerSession struct {
	Contract *TinyOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TinyOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TinyOracleTransactorSession struct {
	Contract     *TinyOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TinyOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type TinyOracleRaw struct {
	Contract *TinyOracle // Generic contract binding to access the raw methods on
}

// TinyOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TinyOracleCallerRaw struct {
	Contract *TinyOracleCaller // Generic read-only contract binding to access the raw methods on
}

// TinyOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TinyOracleTransactorRaw struct {
	Contract *TinyOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTinyOracle creates a new instance of TinyOracle, bound to a specific deployed contract.
func NewTinyOracle(address common.Address, backend bind.ContractBackend) (*TinyOracle, error) {
	contract, err := bindTinyOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TinyOracle{TinyOracleCaller: TinyOracleCaller{contract: contract}, TinyOracleTransactor: TinyOracleTransactor{contract: contract}, TinyOracleFilterer: TinyOracleFilterer{contract: contract}}, nil
}

// NewTinyOracleCaller creates a new read-only instance of TinyOracle, bound to a specific deployed contract.
func NewTinyOracleCaller(address common.Address, caller bind.ContractCaller) (*TinyOracleCaller, error) {
	contract, err := bindTinyOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TinyOracleCaller{contract: contract}, nil
}

// NewTinyOracleTransactor creates a new write-only instance of TinyOracle, bound to a specific deployed contract.
func NewTinyOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*TinyOracleTransactor, error) {
	contract, err := bindTinyOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TinyOracleTransactor{contract: contract}, nil
}

// NewTinyOracleFilterer creates a new log filterer instance of TinyOracle, bound to a specific deployed contract.
func NewTinyOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*TinyOracleFilterer, error) {
	contract, err := bindTinyOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TinyOracleFilterer{contract: contract}, nil
}

// bindTinyOracle binds a generic wrapper to an already deployed contract.
func bindTinyOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TinyOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TinyOracle *TinyOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TinyOracle.Contract.TinyOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TinyOracle *TinyOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TinyOracle.Contract.TinyOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TinyOracle *TinyOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TinyOracle.Contract.TinyOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TinyOracle *TinyOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TinyOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TinyOracle *TinyOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TinyOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TinyOracle *TinyOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TinyOracle.Contract.contract.Transact(opts, method, params...)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(bool value, bool isSet)
func (_TinyOracle *TinyOracleCaller) GetResult(opts *bind.CallOpts) (struct {
	Value bool
	IsSet bool
}, error) {
	var out []interface{}
	err := _TinyOracle.contract.Call(opts, &out, "getResult")

	outstruct := new(struct {
		Value bool
		IsSet bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsSet = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(bool value, bool isSet)
func (_TinyOracle *TinyOracleSession) GetResult() (struct {
	Value bool
	IsSet bool
}, error) {
	return _TinyOracle.Contract.GetResult(&_TinyOracle.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(bool value, bool isSet)
func (_TinyOracle *TinyOracleCallerSession) GetResult() (struct {
	Value bool
	IsSet bool
}, error) {
	return _TinyOracle.Contract.GetResult(&_TinyOracle.CallOpts)
}

// IsTrustedWallet is a free data retrieval call binding the contract method 0x789f5abe.
//
// Solidity: function isTrustedWallet(address _wallet) view returns(bool)
func (_TinyOracle *TinyOracleCaller) IsTrustedWallet(opts *bind.CallOpts, _wallet common.Address) (bool, error) {
	var out []interface{}
	err := _TinyOracle.contract.Call(opts, &out, "isTrustedWallet", _wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedWallet is a free data retrieval call binding the contract method 0x789f5abe.
//
// Solidity: function isTrustedWallet(address _wallet) view returns(bool)
func (_TinyOracle *TinyOracleSession) IsTrustedWallet(_wallet common.Address) (bool, error) {
	return _TinyOracle.Contract.IsTrustedWallet(&_TinyOracle.CallOpts, _wallet)
}

// IsTrustedWallet is a free data retrieval call binding the contract method 0x789f5abe.
//
// Solidity: function isTrustedWallet(address _wallet) view returns(bool)
func (_TinyOracle *TinyOracleCallerSession) IsTrustedWallet(_wallet common.Address) (bool, error) {
	return _TinyOracle.Contract.IsTrustedWallet(&_TinyOracle.CallOpts, _wallet)
}

// TrustedWallets is a free data retrieval call binding the contract method 0xff948fe5.
//
// Solidity: function trustedWallets(address ) view returns(bool)
func (_TinyOracle *TinyOracleCaller) TrustedWallets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TinyOracle.contract.Call(opts, &out, "trustedWallets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TrustedWallets is a free data retrieval call binding the contract method 0xff948fe5.
//
// Solidity: function trustedWallets(address ) view returns(bool)
func (_TinyOracle *TinyOracleSession) TrustedWallets(arg0 common.Address) (bool, error) {
	return _TinyOracle.Contract.TrustedWallets(&_TinyOracle.CallOpts, arg0)
}

// TrustedWallets is a free data retrieval call binding the contract method 0xff948fe5.
//
// Solidity: function trustedWallets(address ) view returns(bool)
func (_TinyOracle *TinyOracleCallerSession) TrustedWallets(arg0 common.Address) (bool, error) {
	return _TinyOracle.Contract.TrustedWallets(&_TinyOracle.CallOpts, arg0)
}

// SetResult is a paid mutator transaction binding the contract method 0xf4b4dc2e.
//
// Solidity: function setResult(bool _result) returns(bool success)
func (_TinyOracle *TinyOracleTransactor) SetResult(opts *bind.TransactOpts, _result bool) (*types.Transaction, error) {
	return _TinyOracle.contract.Transact(opts, "setResult", _result)
}

// SetResult is a paid mutator transaction binding the contract method 0xf4b4dc2e.
//
// Solidity: function setResult(bool _result) returns(bool success)
func (_TinyOracle *TinyOracleSession) SetResult(_result bool) (*types.Transaction, error) {
	return _TinyOracle.Contract.SetResult(&_TinyOracle.TransactOpts, _result)
}

// SetResult is a paid mutator transaction binding the contract method 0xf4b4dc2e.
//
// Solidity: function setResult(bool _result) returns(bool success)
func (_TinyOracle *TinyOracleTransactorSession) SetResult(_result bool) (*types.Transaction, error) {
	return _TinyOracle.Contract.SetResult(&_TinyOracle.TransactOpts, _result)
}

// TinyOracleResultSetIterator is returned from FilterResultSet and is used to iterate over the raw logs and unpacked data for ResultSet events raised by the TinyOracle contract.
type TinyOracleResultSetIterator struct {
	Event *TinyOracleResultSet // Event containing the contract specifics and raw log

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
func (it *TinyOracleResultSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinyOracleResultSet)
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
		it.Event = new(TinyOracleResultSet)
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
func (it *TinyOracleResultSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinyOracleResultSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinyOracleResultSet represents a ResultSet event raised by the TinyOracle contract.
type TinyOracleResultSet struct {
	Value  bool
	Setter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResultSet is a free log retrieval operation binding the contract event 0xf08e01198b2d57b28d82ce475fc894dcff0ca9676ac6f8be2b766a732e054a17.
//
// Solidity: event ResultSet(bool indexed value, address setter)
func (_TinyOracle *TinyOracleFilterer) FilterResultSet(opts *bind.FilterOpts, value []bool) (*TinyOracleResultSetIterator, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TinyOracle.contract.FilterLogs(opts, "ResultSet", valueRule)
	if err != nil {
		return nil, err
	}
	return &TinyOracleResultSetIterator{contract: _TinyOracle.contract, event: "ResultSet", logs: logs, sub: sub}, nil
}

// WatchResultSet is a free log subscription operation binding the contract event 0xf08e01198b2d57b28d82ce475fc894dcff0ca9676ac6f8be2b766a732e054a17.
//
// Solidity: event ResultSet(bool indexed value, address setter)
func (_TinyOracle *TinyOracleFilterer) WatchResultSet(opts *bind.WatchOpts, sink chan<- *TinyOracleResultSet, value []bool) (event.Subscription, error) {

	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _TinyOracle.contract.WatchLogs(opts, "ResultSet", valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinyOracleResultSet)
				if err := _TinyOracle.contract.UnpackLog(event, "ResultSet", log); err != nil {
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

// ParseResultSet is a log parse operation binding the contract event 0xf08e01198b2d57b28d82ce475fc894dcff0ca9676ac6f8be2b766a732e054a17.
//
// Solidity: event ResultSet(bool indexed value, address setter)
func (_TinyOracle *TinyOracleFilterer) ParseResultSet(log types.Log) (*TinyOracleResultSet, error) {
	event := new(TinyOracleResultSet)
	if err := _TinyOracle.contract.UnpackLog(event, "ResultSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
