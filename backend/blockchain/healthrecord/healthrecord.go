// Code generated - DO NOT EDIT.
// This is a simplified example of a Go binding for a HealthRecord contract.

package healthrecord

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// HealthRecordABI is the input ABI used to generate the binding from.
const HealthRecordABI = `[
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "patientName",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "dataHash",
				"type": "string"
			}
		],
		"name": "AddRecord",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

// HealthRecordBin is the compiled bytecode used for deploying new contracts.
const HealthRecordBin = "0x608060405234801561001057600080fd5b506101c8806100206000396000f3fe60806040526004361061002f5760003560e01c80630b5a56aa146100345780631c2c8c0c14610063575b600080fd5b61004a6004803603602081101561004a57600080fd5b50356100b9565b005b61006b6100e1565b6040516100789190610137565b60405180910390f35b60005481565b60008054905090565b6000813590506100d3816101a3565b92915050565b6000602082840312156100ef576100ee61019d565b5b60006100fd848285016100c4565b91505092915050565b61010f81610157565b82525050565b600060208201905061012a6000830184610106565b92915050565b6000819050919050565b61014681610133565b811461015157600080fd5b50565b6000813590506101638161017f565b92915050565b6000806040838503121561017f5761017e61019d565b5b600061018d85828601610156565b925050602061019e85828601610156565b915050925092905056"

// HealthRecordABIParsed is the parsed ABI used by the binding.
var HealthRecordABIParsed abi.ABI

// init 解析 ABI 并初始化全局变量
func init() {
	var err error
	HealthRecordABIParsed, err = abi.JSON(strings.NewReader(HealthRecordABI))
	if err != nil {
		panic(err)
	}
}

// HealthRecord is an auto-generated Go binding around an Ethereum contract.
type HealthRecord struct {
	Address  common.Address
	Contract *bind.BoundContract
}

// NewHealthRecord creates a new instance of HealthRecord, bound to a specific deployed contract.
func NewHealthRecord(address common.Address, backend bind.ContractBackend) (*HealthRecord, error) {
	contract := bind.NewBoundContract(address, HealthRecordABIParsed, backend, backend, backend)
	return &HealthRecord{Address: address, Contract: contract}, nil
}

// AddRecord calls the "AddRecord" method in the contract.
func (hr *HealthRecord) AddRecord(auth *bind.TransactOpts, patientName string, dataHash string) (*types.Transaction, error) {
	return hr.Contract.Transact(auth, "AddRecord", patientName, dataHash)
}

// CallAddRecord is a helper to make a read-only call to the contract method AddRecord.
// 注意：由于 AddRecord 是一个写入方法，通常不用于 Call 操作。此函数仅为示例。
func (hr *HealthRecord) CallAddRecord(opts *bind.CallOpts, patientName string, dataHash string) (*big.Int, error) {
	var out []interface{}
	err := hr.Contract.Call(opts, &out, "AddRecord", patientName, dataHash)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, nil
	}
	// 假设返回的是 uint256 类型
	result, ok := out[0].(*big.Int)
	if !ok {
		return nil, err
	}
	return result, nil
}

// DeployHealthRecord deploys a new HealthRecord contract.
func DeployHealthRecord(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HealthRecord, error) {
	address, tx, contract, err := bind.DeployContract(auth, HealthRecordABIParsed, common.FromHex(HealthRecordBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HealthRecord{Address: address, Contract: contract}, nil
}