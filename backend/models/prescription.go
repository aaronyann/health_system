package models

import "time"

type Prescription struct {
	PrescriptionID int       `gorm:"primaryKey" json:"prescription_id"`
	AppointmentID  uint      `json:"appointment_id"`
	DoctorID       uint      `json:"doctor_id"`
	PatientID      uint      `json:"patient_id"`
	Medications    string    `json:"medications"` // 可存储 JSON 字符串
	Instructions   string    `json:"instructions"`
	IssuedTime     time.Time `json:"issued_time"`
	CreatedAt      time.Time `json:"created_at"`
}
