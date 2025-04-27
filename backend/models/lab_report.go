package models

import "time"

type LabReport struct {
	ReportID      int       `gorm:"primaryKey" json:"report_id"`
	AppointmentID uint      `json:"appointment_id"`
	PatientID     uint      `json:"patient_id"`
	LabName       string    `json:"lab_name"`
	TestType      string    `json:"test_type"`
	Results       string    `json:"results"`
	ResultsHash   string    `json:"results_hash"`
	IssuedTime    time.Time `json:"issued_time"`
	CreatedAt     time.Time `json:"created_at"`
}
