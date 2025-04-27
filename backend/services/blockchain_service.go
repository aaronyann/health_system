package services

import (
	"backend/config"
	"context"
	"fmt"
	"math/big"

	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// 假设你通过 abigen 生成的合约绑定代码在 backend/blockchain/healthrecord 包下
	"backend/blockchain/healthrecord"
)

// AddHealthRecordOnChain 将健康档案数据上链，返回交易哈希（完整实现）
// patientName: 患者姓名
// dataHash: 存储数据文件或IPFS哈希的字符串
func AddHealthRecordOnChain(patientName, dataHash string) (string, error) {
	// 连接到区块链节点（例如 Ganache）
	client, err := ethclient.Dial(config.AppConfig.Blockchain.RPCURL)
	if err != nil {
		return "", fmt.Errorf("连接区块链节点失败: %v", err)
	}
	defer client.Close()

	// 解析配置中的私钥（生产环境中请使用安全的密钥管理方式）
	privateKey, err := crypto.HexToECDSA(config.AppConfig.Blockchain.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("无效私钥: %v", err)
	}

	// 获取链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("获取链ID失败: %v", err)
	}

	// 创建交易签名者 (transactor)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("创建 transactor 失败: %v", err)
	}

	// 设置交易选项
	auth.Value = big.NewInt(0)              // 不发送 ETH
	auth.GasLimit = uint64(3000000)         // 设置合适的 GasLimit
	auth.GasPrice = big.NewInt(20000000000) // 例如 20 Gwei

	// 实例化智能合约对象，合约地址由配置文件提供
	contractAddress := common.HexToAddress(config.AppConfig.Blockchain.ContractAddr)
	instance, err := healthrecord.NewHealthRecord(contractAddress, client)
	if err != nil {
		return "", fmt.Errorf("实例化合约失败: %v", err)
	}

	// 调用智能合约方法 AddRecord，将患者姓名和数据哈希上链
	tx, err := instance.AddRecord(auth, patientName, dataHash)
	if err != nil {
		return "", fmt.Errorf("上链失败: %v", err)
	}

	// 可选：等待交易被挖矿确认（这里简单等待 3 秒，你也可以使用 bind.WaitMined 等待）
	time.Sleep(3 * time.Second)

	// 返回交易哈希
	return tx.Hash().Hex(), nil
}
