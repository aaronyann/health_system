package services

import (
	"backend/models"
	"errors"
	"fmt"

	"strconv"
	"time"
)

// GetPrescriptionList 获取处方列表
func GetPrescriptionList() ([]models.Prescription, error) {
	var prescriptions []models.Prescription
	if err := DB.Find(&prescriptions).Error; err != nil {
		return nil, err
	}
	return prescriptions, nil
}

// GetPrescriptionByID 根据 ID 获取处方详情
func GetPrescriptionByID(id string) (*models.Prescription, error) {
	pid, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("无效的处方ID")
	}
	var prescription models.Prescription
	if err := DB.First(&prescription, pid).Error; err != nil {
		return nil, fmt.Errorf("查询处方失败: %v", err)
	}
	return &prescription, nil
}

// CreatePrescription 创建处方
func CreatePrescription(appointmentID, doctorID, patientID uint, medications, instructions, issuedTimeStr string) (*models.Prescription, error) {
	issuedTime, err := time.Parse("2006-01-02 15:04:05", issuedTimeStr)
	if err != nil {
		return nil, errors.New("处方日期格式错误")
	}
	prescription := models.Prescription{
		AppointmentID: appointmentID,
		DoctorID:      doctorID,
		PatientID:     patientID,
		Medications:   medications,
		Instructions:  instructions,
		IssuedTime:    issuedTime,
		CreatedAt:     time.Now(),
	}
	if err := DB.Create(&prescription).Error; err != nil {
		return nil, err
	}
	return &prescription, nil
}
