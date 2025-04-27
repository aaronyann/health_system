package models

import "time"

type InsuranceInfo struct {
	InsuranceID  int       `gorm:"primaryKey" json:"insurance_id"`
	PatientID    uint      `json:"patient_id"`
	Provider     string    `json:"provider"`
	PolicyNumber string    `json:"policy_number"`
	ValidFrom    time.Time `json:"valid_from"`
	ValidTo      time.Time `json:"valid_to"`
	Coverage     string    `json:"coverage"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
