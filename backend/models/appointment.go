package models

import "time"

type Appointment struct {
	AppointmentID   uint      `gorm:"primaryKey" json:"appointment_id"`
	PatientID       uint      `json:"patient_id"`
	DoctorID        uint      `json:"doctor_id"`
	DepartmentID    uint      `json:"department_id"`
	AppointmentTime time.Time `json:"appointment_time"`
	Status          string    `json:"status"`
	Notes           string    `json:"notes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
