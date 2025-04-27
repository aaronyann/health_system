package services

import (
	"backend/models"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// GetAppointments 获取预约列表
func GetAppointments() ([]models.Appointment, error) {
	var appointments []models.Appointment
	if err := DB.Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

// CreateAppointment 创建预约
func CreateAppointment(patientID, doctorID, departmentID uint, appointmentTimeStr, notes string) (*models.Appointment, error) {
	appointmentTime, err := time.Parse("2006-01-02 15:04:05", appointmentTimeStr)
	if err != nil {
		return nil, errors.New("预约时间格式错误")
	}
	app := models.Appointment{
		PatientID:       patientID,
		DoctorID:        doctorID,
		DepartmentID:    departmentID,
		AppointmentTime: appointmentTime,
		Status:          "预约中",
		Notes:           notes,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	if err := DB.Create(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

// GetAppointmentByID 根据 ID 获取预约详情
func GetAppointmentByID(id string) (*models.Appointment, error) {
	aid, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("无效的预约ID")
	}
	var app models.Appointment
	if err := DB.First(&app, aid).Error; err != nil {
		return nil, fmt.Errorf("查询预约失败: %v", err)
	}
	return &app, nil
}
