package services

import (
	"backend/models"
	"errors"
	"fmt"

	"strconv"
	"time"
)

// GetHealthRecords 从数据库中获取健康档案列表（示例，实际可加入分页）
func GetHealthRecords() ([]models.HealthRecord, error) {
	var records []models.HealthRecord
	if err := DB.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// CreateHealthRecord 将健康档案写入数据库，并调用区块链上链
func CreateHealthRecord(patientName, dataHash string) (string, error) {
	// 调用区块链服务上链数据
	txHash, err := AddHealthRecordOnChain(patientName, dataHash)
	if err != nil {
		return "", err
	}
	// 保存记录到数据库
	record := models.HealthRecord{
		PatientID:          0, // 此处需要根据实际情况关联患者ID
		BlockchainRecordID: 0, // 可存入链上返回的编号
		PatientName:        patientName,
		DataHash:           dataHash,
		Timestamp:          time.Now(),
	}
	if err := DB.Create(&record).Error; err != nil {
		return "", err
	}
	return txHash, nil
}

// GetHealthRecordByID 根据 ID 查询健康档案详情
func GetHealthRecordByID(id string) (*models.HealthRecord, error) {
	recordID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("无效的记录ID")
	}
	var record models.HealthRecord
	if err := DB.First(&record, recordID).Error; err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	return &record, nil
}
