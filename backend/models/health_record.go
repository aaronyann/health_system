package models

import "time"

type HealthRecord struct {
	RecordID           uint      `gorm:"primaryKey" json:"record_id"`
	PatientID          uint      `json:"patient_id"`
	BlockchainRecordID int64     `json:"blockchain_record_id"` // 上链编号
	PatientName        string    `json:"patient_name"`
	DataHash           string    `json:"data_hash"`
	Timestamp          time.Time `json:"timestamp"` // 区块链记录时间
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
