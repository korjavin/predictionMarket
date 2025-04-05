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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_question\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPayout\",\"type\":\"uint256\"}],\"name\":\"BetClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bettor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"prediction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BetPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bettor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PayoutSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"betFalse\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"betTrue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"falseBets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"falseBettors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBetInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"Question\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ReleaseDate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"IsClosed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"TotalTrueBets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TotalFalseBets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOdds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"trueOdds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"falseOdds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractTinyOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"question\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFalseBets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTrueBets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"trueBets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"trueBettors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f60075f6101000a81548160ff02191690831515021790555034801562000029575f80fd5b5060405162001dce38038062001dce83398181016040528101906200004f919062000313565b42811162000094576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200008b906200040f565b60405180910390fd5b825f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160019081620000e491906200065d565b508060028190555050505062000741565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620001318262000106565b9050919050565b620001438162000125565b81146200014e575f80fd5b50565b5f81519050620001618162000138565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620001b7826200016f565b810181811067ffffffffffffffff82111715620001d957620001d86200017f565b5b80604052505050565b5f620001ed620000f5565b9050620001fb8282620001ac565b919050565b5f67ffffffffffffffff8211156200021d576200021c6200017f565b5b62000228826200016f565b9050602081019050919050565b5f5b838110156200025457808201518184015260208101905062000237565b5f8484015250505050565b5f620002756200026f8462000200565b620001e2565b9050828152602081018484840111156200029457620002936200016b565b5b620002a184828562000235565b509392505050565b5f82601f830112620002c057620002bf62000167565b5b8151620002d28482602086016200025f565b91505092915050565b5f819050919050565b620002ef81620002db565b8114620002fa575f80fd5b50565b5f815190506200030d81620002e4565b92915050565b5f805f606084860312156200032d576200032c620000fe565b5b5f6200033c8682870162000151565b935050602084015167ffffffffffffffff81111562000360576200035f62000102565b5b6200036e86828701620002a9565b92505060406200038186828701620002fd565b9150509250925092565b5f82825260208201905092915050565b7f52656c656173652064617465206d75737420626520696e2074686520667574755f8201527f7265000000000000000000000000000000000000000000000000000000000000602082015250565b5f620003f76022836200038b565b915062000404826200039b565b604082019050919050565b5f6020820190508181035f8301526200042881620003e9565b9050919050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200047e57607f821691505b60208210810362000494576200049362000439565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004f87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620004bb565b620005048683620004bb565b95508019841693508086168417925050509392505050565b5f819050919050565b5f620005456200053f6200053984620002db565b6200051c565b620002db565b9050919050565b5f819050919050565b620005608362000525565b620005786200056f826200054c565b848454620004c7565b825550505050565b5f90565b6200058e62000580565b6200059b81848462000555565b505050565b5b81811015620005c257620005b65f8262000584565b600181019050620005a1565b5050565b601f8211156200061157620005db816200049a565b620005e684620004ac565b81016020851015620005f6578190505b6200060e6200060585620004ac565b830182620005a0565b50505b505050565b5f82821c905092915050565b5f620006335f198460080262000616565b1980831691505092915050565b5f6200064d838362000622565b9150826002028217905092915050565b62000668826200042f565b67ffffffffffffffff8111156200068457620006836200017f565b5b62000690825462000466565b6200069d828285620005c6565b5f60209050601f831160018114620006d3575f8415620006be578287015190505b620006ca858262000640565b86555062000739565b601f198416620006e3866200049a565b5f5b828110156200070c57848901518255600182019150602085019450602081019050620006e5565b868310156200072c578489015162000728601f89168262000622565b8355505b6001600288020188555050505b505050505050565b61167f806200074f5f395ff3fe6080604052600436106100e7575f3560e01c8063724e1d7c11610089578063ace10fb311610058578063ace10fb3146102e2578063b9e3e2db146102ec578063cfe7b77014610316578063eaa968cc14610344576100e7565b8063724e1d7c146102165780637dc0d1d01461025257806392b97ceb1461027c578063a0088530146102a6576100e7565b80634c5be574116100c55780634c5be5741461017b5780635425b94c146101a6578063597e1fb5146101e2578063681fec7f1461020c576100e7565b8063033dfe89146100eb578063162559c7146101155780633fad9ae014610151575b5f80fd5b3480156100f6575f80fd5b506100ff61035a565b60405161010c9190610dcc565b60405180910390f35b348015610120575f80fd5b5061013b60048036038101906101369190610e43565b610360565b6040516101489190610dcc565b60405180910390f35b34801561015c575f80fd5b50610165610375565b6040516101729190610ef8565b60405180910390f35b348015610186575f80fd5b5061018f610401565b60405161019d929190610f18565b60405180910390f35b3480156101b1575f80fd5b506101cc60048036038101906101c79190610f69565b61045b565b6040516101d99190610fa3565b60405180910390f35b3480156101ed575f80fd5b506101f6610496565b6040516102039190610fd6565b60405180910390f35b6102146104a8565b005b348015610221575f80fd5b5061023c60048036038101906102379190610e43565b6104b3565b6040516102499190610dcc565b60405180910390f35b34801561025d575f80fd5b506102666104c8565b604051610273919061104a565b60405180910390f35b348015610287575f80fd5b506102906104eb565b60405161029d9190610dcc565b60405180910390f35b3480156102b1575f80fd5b506102cc60048036038101906102c79190610f69565b6104f1565b6040516102d99190610fa3565b60405180910390f35b6102ea61052c565b005b3480156102f7575f80fd5b50610300610538565b60405161030d9190610dcc565b60405180910390f35b348015610321575f80fd5b5061032a61053e565b60405161033b959493929190611063565b60405180910390f35b34801561034f575f80fd5b506103586105f9565b005b60035481565b6006602052805f5260405f205f915090505481565b60018054610382906110e8565b80601f01602080910402602001604051908101604052809291908181526020018280546103ae906110e8565b80156103f95780601f106103d0576101008083540402835291602001916103f9565b820191905f5260205f20905b8154815290600101906020018083116103dc57829003601f168201915b505050505081565b5f805f6004546003546104149190611145565b90505f810361042a576032809250925050610457565b80606460035461043a9190611178565b61044491906111e6565b92508260646104539190611216565b9150505b9091565b6009818154811061046a575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60075f9054906101000a900460ff1681565b6104b15f6108f9565b565b6005602052805f5260405f205f915090505481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b60088181548110610500575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61053660016108f9565b565b60025481565b60605f805f80600160025460075f9054906101000a900460ff1660035460045484805461056a906110e8565b80601f0160208091040260200160405190810160405280929190818152602001828054610596906110e8565b80156105e15780601f106105b8576101008083540402835291602001916105e1565b820191905f5260205f20905b8154815290600101906020018083116105c457829003601f168201915b50505050509450945094509450945094509091929394565b60075f9054906101000a900460ff1615610648576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063f90611293565b60405180910390fd5b60025442101561068d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068490611321565b60405180910390fd5b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663de2927896040518163ffffffff1660e01b81526004016040805180830381865afa1580156106f6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061071a9190611369565b915091508061075e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610755906113f1565b60405180910390fd5b600160075f6101000a81548160ff0219169083151502179055505f6004546003546107899190611145565b90507f7f83912ab319ee20a87030b0cb15de42af62d5bb98ddb323a7857026d77b26de83826040516107bc92919061140f565b60405180910390a182156108615761085c600880548060200260200160405190810160405280929190818152602001828054801561084c57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610803575b5050505050600560035484610c4d565b6108f4565b6108f360098054806020026020016040519081016040528092919081815260200182805480156108e357602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161089a575b5050505050600660045484610c4d565b5b505050565b60075f9054906101000a900460ff1615610948576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093f90611480565b60405180910390fd5b600254421061098c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610983906114e8565b60405180910390fd5b5f34116109ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c590611576565b60405180910390fd5b8015610ae9575f60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610a7957600833908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b3460055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ac59190611145565b925050819055503460035f828254610add9190611145565b92505081905550610bfa565b5f60065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610b8e57600933908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b3460065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610bda9190611145565b925050819055503460045f828254610bf29190611145565b925050819055505b3373ffffffffffffffffffffffffffffffffffffffff167fe57fa5a543ebcdebf2a7d7102ac913e14454cec006468e43a6916e95bdfcea9e8234604051610c4292919061140f565b60405180910390a250565b5f5b8451811015610dad575f858281518110610c6c57610c6b611594565b5b602002602001015190505f855f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f811115610d98575f858583610ccc9190611178565b610cd691906111e6565b90505f8373ffffffffffffffffffffffffffffffffffffffff1682604051610cfd906115ee565b5f6040518083038185875af1925050503d805f8114610d37576040519150601f19603f3d011682016040523d82523d5f602084013e610d3c565b606091505b505090508015610d95578373ffffffffffffffffffffffffffffffffffffffff167f6c114f03096d92b098f0f087d8fd4a756d362d86a9649a07142404e7fd1d77b583604051610d8c9190610dcc565b60405180910390a25b50505b50508080610da590611602565b915050610c4f565b5050505050565b5f819050919050565b610dc681610db4565b82525050565b5f602082019050610ddf5f830184610dbd565b92915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e1282610de9565b9050919050565b610e2281610e08565b8114610e2c575f80fd5b50565b5f81359050610e3d81610e19565b92915050565b5f60208284031215610e5857610e57610de5565b5b5f610e6584828501610e2f565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610ea5578082015181840152602081019050610e8a565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610eca82610e6e565b610ed48185610e78565b9350610ee4818560208601610e88565b610eed81610eb0565b840191505092915050565b5f6020820190508181035f830152610f108184610ec0565b905092915050565b5f604082019050610f2b5f830185610dbd565b610f386020830184610dbd565b9392505050565b610f4881610db4565b8114610f52575f80fd5b50565b5f81359050610f6381610f3f565b92915050565b5f60208284031215610f7e57610f7d610de5565b5b5f610f8b84828501610f55565b91505092915050565b610f9d81610e08565b82525050565b5f602082019050610fb65f830184610f94565b92915050565b5f8115159050919050565b610fd081610fbc565b82525050565b5f602082019050610fe95f830184610fc7565b92915050565b5f819050919050565b5f61101261100d61100884610de9565b610fef565b610de9565b9050919050565b5f61102382610ff8565b9050919050565b5f61103482611019565b9050919050565b6110448161102a565b82525050565b5f60208201905061105d5f83018461103b565b92915050565b5f60a0820190508181035f83015261107b8188610ec0565b905061108a6020830187610dbd565b6110976040830186610fc7565b6110a46060830185610dbd565b6110b16080830184610dbd565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81905092915050565b50565b5f6115d95f836115c1565b91506115e4826115cb565b5f82019050919050565b5f6115f8826115ce565b9150819050919050565b5f61160c82610db4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361163e5761163d611118565b5b60018201905091905056fea264697066735822122083540f6d7a57530cf186dfb86e3057639aa1c6866c6e0f7c62994d4f05458e9864736f6c63430008140033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address, _question string, _releaseDate *big.Int) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _oracle, _question, _releaseDate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_Contracts *ContractsCaller) Closed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "closed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_Contracts *ContractsSession) Closed() (bool, error) {
	return _Contracts.Contract.Closed(&_Contracts.CallOpts)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_Contracts *ContractsCallerSession) Closed() (bool, error) {
	return _Contracts.Contract.Closed(&_Contracts.CallOpts)
}

// FalseBets is a free data retrieval call binding the contract method 0x162559c7.
//
// Solidity: function falseBets(address ) view returns(uint256)
func (_Contracts *ContractsCaller) FalseBets(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "falseBets", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FalseBets is a free data retrieval call binding the contract method 0x162559c7.
//
// Solidity: function falseBets(address ) view returns(uint256)
func (_Contracts *ContractsSession) FalseBets(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.FalseBets(&_Contracts.CallOpts, arg0)
}

// FalseBets is a free data retrieval call binding the contract method 0x162559c7.
//
// Solidity: function falseBets(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) FalseBets(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.FalseBets(&_Contracts.CallOpts, arg0)
}

// FalseBettors is a free data retrieval call binding the contract method 0x5425b94c.
//
// Solidity: function falseBettors(uint256 ) view returns(address)
func (_Contracts *ContractsCaller) FalseBettors(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "falseBettors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FalseBettors is a free data retrieval call binding the contract method 0x5425b94c.
//
// Solidity: function falseBettors(uint256 ) view returns(address)
func (_Contracts *ContractsSession) FalseBettors(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.FalseBettors(&_Contracts.CallOpts, arg0)
}

// FalseBettors is a free data retrieval call binding the contract method 0x5425b94c.
//
// Solidity: function falseBettors(uint256 ) view returns(address)
func (_Contracts *ContractsCallerSession) FalseBettors(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.FalseBettors(&_Contracts.CallOpts, arg0)
}

// GetBetInfo is a free data retrieval call binding the contract method 0xcfe7b770.
//
// Solidity: function getBetInfo() view returns(string Question, uint256 ReleaseDate, bool IsClosed, uint256 TotalTrueBets, uint256 TotalFalseBets)
func (_Contracts *ContractsCaller) GetBetInfo(opts *bind.CallOpts) (struct {
	Question       string
	ReleaseDate    *big.Int
	IsClosed       bool
	TotalTrueBets  *big.Int
	TotalFalseBets *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBetInfo")

	outstruct := new(struct {
		Question       string
		ReleaseDate    *big.Int
		IsClosed       bool
		TotalTrueBets  *big.Int
		TotalFalseBets *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Question = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ReleaseDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsClosed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.TotalTrueBets = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalFalseBets = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBetInfo is a free data retrieval call binding the contract method 0xcfe7b770.
//
// Solidity: function getBetInfo() view returns(string Question, uint256 ReleaseDate, bool IsClosed, uint256 TotalTrueBets, uint256 TotalFalseBets)
func (_Contracts *ContractsSession) GetBetInfo() (struct {
	Question       string
	ReleaseDate    *big.Int
	IsClosed       bool
	TotalTrueBets  *big.Int
	TotalFalseBets *big.Int
}, error) {
	return _Contracts.Contract.GetBetInfo(&_Contracts.CallOpts)
}

// GetBetInfo is a free data retrieval call binding the contract method 0xcfe7b770.
//
// Solidity: function getBetInfo() view returns(string Question, uint256 ReleaseDate, bool IsClosed, uint256 TotalTrueBets, uint256 TotalFalseBets)
func (_Contracts *ContractsCallerSession) GetBetInfo() (struct {
	Question       string
	ReleaseDate    *big.Int
	IsClosed       bool
	TotalTrueBets  *big.Int
	TotalFalseBets *big.Int
}, error) {
	return _Contracts.Contract.GetBetInfo(&_Contracts.CallOpts)
}

// GetOdds is a free data retrieval call binding the contract method 0x4c5be574.
//
// Solidity: function getOdds() view returns(uint256 trueOdds, uint256 falseOdds)
func (_Contracts *ContractsCaller) GetOdds(opts *bind.CallOpts) (struct {
	TrueOdds  *big.Int
	FalseOdds *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getOdds")

	outstruct := new(struct {
		TrueOdds  *big.Int
		FalseOdds *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TrueOdds = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FalseOdds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOdds is a free data retrieval call binding the contract method 0x4c5be574.
//
// Solidity: function getOdds() view returns(uint256 trueOdds, uint256 falseOdds)
func (_Contracts *ContractsSession) GetOdds() (struct {
	TrueOdds  *big.Int
	FalseOdds *big.Int
}, error) {
	return _Contracts.Contract.GetOdds(&_Contracts.CallOpts)
}

// GetOdds is a free data retrieval call binding the contract method 0x4c5be574.
//
// Solidity: function getOdds() view returns(uint256 trueOdds, uint256 falseOdds)
func (_Contracts *ContractsCallerSession) GetOdds() (struct {
	TrueOdds  *big.Int
	FalseOdds *big.Int
}, error) {
	return _Contracts.Contract.GetOdds(&_Contracts.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contracts *ContractsCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contracts *ContractsSession) Oracle() (common.Address, error) {
	return _Contracts.Contract.Oracle(&_Contracts.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Contracts *ContractsCallerSession) Oracle() (common.Address, error) {
	return _Contracts.Contract.Oracle(&_Contracts.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Contracts *ContractsCaller) Question(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "question")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Contracts *ContractsSession) Question() (string, error) {
	return _Contracts.Contract.Question(&_Contracts.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Contracts *ContractsCallerSession) Question() (string, error) {
	return _Contracts.Contract.Question(&_Contracts.CallOpts)
}

// ReleaseDate is a free data retrieval call binding the contract method 0xb9e3e2db.
//
// Solidity: function releaseDate() view returns(uint256)
func (_Contracts *ContractsCaller) ReleaseDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "releaseDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseDate is a free data retrieval call binding the contract method 0xb9e3e2db.
//
// Solidity: function releaseDate() view returns(uint256)
func (_Contracts *ContractsSession) ReleaseDate() (*big.Int, error) {
	return _Contracts.Contract.ReleaseDate(&_Contracts.CallOpts)
}

// ReleaseDate is a free data retrieval call binding the contract method 0xb9e3e2db.
//
// Solidity: function releaseDate() view returns(uint256)
func (_Contracts *ContractsCallerSession) ReleaseDate() (*big.Int, error) {
	return _Contracts.Contract.ReleaseDate(&_Contracts.CallOpts)
}

// TotalFalseBets is a free data retrieval call binding the contract method 0x92b97ceb.
//
// Solidity: function totalFalseBets() view returns(uint256)
func (_Contracts *ContractsCaller) TotalFalseBets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalFalseBets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFalseBets is a free data retrieval call binding the contract method 0x92b97ceb.
//
// Solidity: function totalFalseBets() view returns(uint256)
func (_Contracts *ContractsSession) TotalFalseBets() (*big.Int, error) {
	return _Contracts.Contract.TotalFalseBets(&_Contracts.CallOpts)
}

// TotalFalseBets is a free data retrieval call binding the contract method 0x92b97ceb.
//
// Solidity: function totalFalseBets() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalFalseBets() (*big.Int, error) {
	return _Contracts.Contract.TotalFalseBets(&_Contracts.CallOpts)
}

// TotalTrueBets is a free data retrieval call binding the contract method 0x033dfe89.
//
// Solidity: function totalTrueBets() view returns(uint256)
func (_Contracts *ContractsCaller) TotalTrueBets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalTrueBets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTrueBets is a free data retrieval call binding the contract method 0x033dfe89.
//
// Solidity: function totalTrueBets() view returns(uint256)
func (_Contracts *ContractsSession) TotalTrueBets() (*big.Int, error) {
	return _Contracts.Contract.TotalTrueBets(&_Contracts.CallOpts)
}

// TotalTrueBets is a free data retrieval call binding the contract method 0x033dfe89.
//
// Solidity: function totalTrueBets() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalTrueBets() (*big.Int, error) {
	return _Contracts.Contract.TotalTrueBets(&_Contracts.CallOpts)
}

// TrueBets is a free data retrieval call binding the contract method 0x724e1d7c.
//
// Solidity: function trueBets(address ) view returns(uint256)
func (_Contracts *ContractsCaller) TrueBets(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "trueBets", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrueBets is a free data retrieval call binding the contract method 0x724e1d7c.
//
// Solidity: function trueBets(address ) view returns(uint256)
func (_Contracts *ContractsSession) TrueBets(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.TrueBets(&_Contracts.CallOpts, arg0)
}

// TrueBets is a free data retrieval call binding the contract method 0x724e1d7c.
//
// Solidity: function trueBets(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) TrueBets(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.TrueBets(&_Contracts.CallOpts, arg0)
}

// TrueBettors is a free data retrieval call binding the contract method 0xa0088530.
//
// Solidity: function trueBettors(uint256 ) view returns(address)
func (_Contracts *ContractsCaller) TrueBettors(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "trueBettors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrueBettors is a free data retrieval call binding the contract method 0xa0088530.
//
// Solidity: function trueBettors(uint256 ) view returns(address)
func (_Contracts *ContractsSession) TrueBettors(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.TrueBettors(&_Contracts.CallOpts, arg0)
}

// TrueBettors is a free data retrieval call binding the contract method 0xa0088530.
//
// Solidity: function trueBettors(uint256 ) view returns(address)
func (_Contracts *ContractsCallerSession) TrueBettors(arg0 *big.Int) (common.Address, error) {
	return _Contracts.Contract.TrueBettors(&_Contracts.CallOpts, arg0)
}

// BetFalse is a paid mutator transaction binding the contract method 0x681fec7f.
//
// Solidity: function betFalse() payable returns()
func (_Contracts *ContractsTransactor) BetFalse(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "betFalse")
}

// BetFalse is a paid mutator transaction binding the contract method 0x681fec7f.
//
// Solidity: function betFalse() payable returns()
func (_Contracts *ContractsSession) BetFalse() (*types.Transaction, error) {
	return _Contracts.Contract.BetFalse(&_Contracts.TransactOpts)
}

// BetFalse is a paid mutator transaction binding the contract method 0x681fec7f.
//
// Solidity: function betFalse() payable returns()
func (_Contracts *ContractsTransactorSession) BetFalse() (*types.Transaction, error) {
	return _Contracts.Contract.BetFalse(&_Contracts.TransactOpts)
}

// BetTrue is a paid mutator transaction binding the contract method 0xace10fb3.
//
// Solidity: function betTrue() payable returns()
func (_Contracts *ContractsTransactor) BetTrue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "betTrue")
}

// BetTrue is a paid mutator transaction binding the contract method 0xace10fb3.
//
// Solidity: function betTrue() payable returns()
func (_Contracts *ContractsSession) BetTrue() (*types.Transaction, error) {
	return _Contracts.Contract.BetTrue(&_Contracts.TransactOpts)
}

// BetTrue is a paid mutator transaction binding the contract method 0xace10fb3.
//
// Solidity: function betTrue() payable returns()
func (_Contracts *ContractsTransactorSession) BetTrue() (*types.Transaction, error) {
	return _Contracts.Contract.BetTrue(&_Contracts.TransactOpts)
}

// CloseBet is a paid mutator transaction binding the contract method 0xeaa968cc.
//
// Solidity: function closeBet() returns()
func (_Contracts *ContractsTransactor) CloseBet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "closeBet")
}

// CloseBet is a paid mutator transaction binding the contract method 0xeaa968cc.
//
// Solidity: function closeBet() returns()
func (_Contracts *ContractsSession) CloseBet() (*types.Transaction, error) {
	return _Contracts.Contract.CloseBet(&_Contracts.TransactOpts)
}

// CloseBet is a paid mutator transaction binding the contract method 0xeaa968cc.
//
// Solidity: function closeBet() returns()
func (_Contracts *ContractsTransactorSession) CloseBet() (*types.Transaction, error) {
	return _Contracts.Contract.CloseBet(&_Contracts.TransactOpts)
}

// ContractsBetClosedIterator is returned from FilterBetClosed and is used to iterate over the raw logs and unpacked data for BetClosed events raised by the Contracts contract.
type ContractsBetClosedIterator struct {
	Event *ContractsBetClosed // Event containing the contract specifics and raw log

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
func (it *ContractsBetClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBetClosed)
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
		it.Event = new(ContractsBetClosed)
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
func (it *ContractsBetClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBetClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBetClosed represents a BetClosed event raised by the Contracts contract.
type ContractsBetClosed struct {
	Result      bool
	TotalPayout *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBetClosed is a free log retrieval operation binding the contract event 0x7f83912ab319ee20a87030b0cb15de42af62d5bb98ddb323a7857026d77b26de.
//
// Solidity: event BetClosed(bool result, uint256 totalPayout)
func (_Contracts *ContractsFilterer) FilterBetClosed(opts *bind.FilterOpts) (*ContractsBetClosedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "BetClosed")
	if err != nil {
		return nil, err
	}
	return &ContractsBetClosedIterator{contract: _Contracts.contract, event: "BetClosed", logs: logs, sub: sub}, nil
}

// WatchBetClosed is a free log subscription operation binding the contract event 0x7f83912ab319ee20a87030b0cb15de42af62d5bb98ddb323a7857026d77b26de.
//
// Solidity: event BetClosed(bool result, uint256 totalPayout)
func (_Contracts *ContractsFilterer) WatchBetClosed(opts *bind.WatchOpts, sink chan<- *ContractsBetClosed) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "BetClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBetClosed)
				if err := _Contracts.contract.UnpackLog(event, "BetClosed", log); err != nil {
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

// ParseBetClosed is a log parse operation binding the contract event 0x7f83912ab319ee20a87030b0cb15de42af62d5bb98ddb323a7857026d77b26de.
//
// Solidity: event BetClosed(bool result, uint256 totalPayout)
func (_Contracts *ContractsFilterer) ParseBetClosed(log types.Log) (*ContractsBetClosed, error) {
	event := new(ContractsBetClosed)
	if err := _Contracts.contract.UnpackLog(event, "BetClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsBetPlacedIterator is returned from FilterBetPlaced and is used to iterate over the raw logs and unpacked data for BetPlaced events raised by the Contracts contract.
type ContractsBetPlacedIterator struct {
	Event *ContractsBetPlaced // Event containing the contract specifics and raw log

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
func (it *ContractsBetPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBetPlaced)
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
		it.Event = new(ContractsBetPlaced)
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
func (it *ContractsBetPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBetPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBetPlaced represents a BetPlaced event raised by the Contracts contract.
type ContractsBetPlaced struct {
	Bettor     common.Address
	Prediction bool
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBetPlaced is a free log retrieval operation binding the contract event 0xe57fa5a543ebcdebf2a7d7102ac913e14454cec006468e43a6916e95bdfcea9e.
//
// Solidity: event BetPlaced(address indexed bettor, bool prediction, uint256 amount)
func (_Contracts *ContractsFilterer) FilterBetPlaced(opts *bind.FilterOpts, bettor []common.Address) (*ContractsBetPlacedIterator, error) {

	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "BetPlaced", bettorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsBetPlacedIterator{contract: _Contracts.contract, event: "BetPlaced", logs: logs, sub: sub}, nil
}

// WatchBetPlaced is a free log subscription operation binding the contract event 0xe57fa5a543ebcdebf2a7d7102ac913e14454cec006468e43a6916e95bdfcea9e.
//
// Solidity: event BetPlaced(address indexed bettor, bool prediction, uint256 amount)
func (_Contracts *ContractsFilterer) WatchBetPlaced(opts *bind.WatchOpts, sink chan<- *ContractsBetPlaced, bettor []common.Address) (event.Subscription, error) {

	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "BetPlaced", bettorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBetPlaced)
				if err := _Contracts.contract.UnpackLog(event, "BetPlaced", log); err != nil {
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

// ParseBetPlaced is a log parse operation binding the contract event 0xe57fa5a543ebcdebf2a7d7102ac913e14454cec006468e43a6916e95bdfcea9e.
//
// Solidity: event BetPlaced(address indexed bettor, bool prediction, uint256 amount)
func (_Contracts *ContractsFilterer) ParseBetPlaced(log types.Log) (*ContractsBetPlaced, error) {
	event := new(ContractsBetPlaced)
	if err := _Contracts.contract.UnpackLog(event, "BetPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPayoutSentIterator is returned from FilterPayoutSent and is used to iterate over the raw logs and unpacked data for PayoutSent events raised by the Contracts contract.
type ContractsPayoutSentIterator struct {
	Event *ContractsPayoutSent // Event containing the contract specifics and raw log

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
func (it *ContractsPayoutSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPayoutSent)
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
		it.Event = new(ContractsPayoutSent)
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
func (it *ContractsPayoutSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPayoutSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPayoutSent represents a PayoutSent event raised by the Contracts contract.
type ContractsPayoutSent struct {
	Bettor common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPayoutSent is a free log retrieval operation binding the contract event 0x6c114f03096d92b098f0f087d8fd4a756d362d86a9649a07142404e7fd1d77b5.
//
// Solidity: event PayoutSent(address indexed bettor, uint256 amount)
func (_Contracts *ContractsFilterer) FilterPayoutSent(opts *bind.FilterOpts, bettor []common.Address) (*ContractsPayoutSentIterator, error) {

	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "PayoutSent", bettorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsPayoutSentIterator{contract: _Contracts.contract, event: "PayoutSent", logs: logs, sub: sub}, nil
}

// WatchPayoutSent is a free log subscription operation binding the contract event 0x6c114f03096d92b098f0f087d8fd4a756d362d86a9649a07142404e7fd1d77b5.
//
// Solidity: event PayoutSent(address indexed bettor, uint256 amount)
func (_Contracts *ContractsFilterer) WatchPayoutSent(opts *bind.WatchOpts, sink chan<- *ContractsPayoutSent, bettor []common.Address) (event.Subscription, error) {

	var bettorRule []interface{}
	for _, bettorItem := range bettor {
		bettorRule = append(bettorRule, bettorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "PayoutSent", bettorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPayoutSent)
				if err := _Contracts.contract.UnpackLog(event, "PayoutSent", log); err != nil {
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

// ParsePayoutSent is a log parse operation binding the contract event 0x6c114f03096d92b098f0f087d8fd4a756d362d86a9649a07142404e7fd1d77b5.
//
// Solidity: event PayoutSent(address indexed bettor, uint256 amount)
func (_Contracts *ContractsFilterer) ParsePayoutSent(log types.Log) (*ContractsPayoutSent, error) {
	event := new(ContractsPayoutSent)
	if err := _Contracts.contract.UnpackLog(event, "PayoutSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinyBetContract is an alias for the Contracts type
type TinyBetContract Contracts

// DeployTinyBetContract deploys a new Ethereum contract, binding an instance of TinyBetContract to it.
func DeployTinyBetContract(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address, _question string, _releaseDate *big.Int) (common.Address, *types.Transaction, *TinyBetContract, error) {
	address, tx, contract, err := DeployContracts(auth, backend, _oracle, _question, _releaseDate)
	return address, tx, (*TinyBetContract)(contract), err
}

// NewTinyBetContract creates a new instance of TinyBetContract, bound to a specific deployed contract.
func NewTinyBetContract(address common.Address, backend bind.ContractBackend) (*TinyBetContract, error) {
	contract, err := NewContracts(address, backend)
	return (*TinyBetContract)(contract), err
}

// This file provides aliases for the generated bindings to match
// the function names expected by handlers.go

// DeployTinyBet deploys a new Ethereum contract, binding an instance of TinyBet to it.
func DeployTinyBet(auth *bind.TransactOpts, backend bind.ContractBackend, _oracle common.Address, _question string, _releaseDate *big.Int) (common.Address, *types.Transaction, *TinyBetContract, error) {
	return DeployTinyBetContract(auth, backend, _oracle, _question, _releaseDate)
}

// NewTinyBet creates a new instance of TinyBet, bound to a specific deployed contract.
func NewTinyBet(address common.Address, backend bind.ContractBackend) (*TinyBetContract, error) {
	return NewTinyBetContract(address, backend)
}
