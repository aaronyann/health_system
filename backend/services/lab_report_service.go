package services

import (
	"backend/models"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// GetLabReports 获取检验报告列表
func GetLabReports() ([]models.LabReport, error) {
	var reports []models.LabReport
	if err := DB.Find(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

// GetLabReportByID 根据 ID 获取检验报告详情
func GetLabReportByID(id string) (*models.LabReport, error) {
	rid, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("无效的报告ID")
	}
	var report models.LabReport
	if err := DB.First(&report, rid).Error; err != nil {
		return nil, fmt.Errorf("查询检验报告失败: %v", err)
	}
	return &report, nil
}

// CreateLabReport 创建检验报告
func CreateLabReport(appointmentID, patientID uint, labName, testType, results, resultsHash, issuedTimeStr string) (*models.LabReport, error) {
	issuedTime, err := time.Parse("2006-01-02 15:04:05", issuedTimeStr)
	if err != nil {
		return nil, errors.New("报告日期格式错误")
	}
	report := models.LabReport{
		AppointmentID: appointmentID,
		PatientID:     patientID,
		LabName:       labName,
		TestType:      testType,
		Results:       results,
		ResultsHash:   resultsHash,
		IssuedTime:    issuedTime,
		CreatedAt:     time.Now(),
	}
	if err := DB.Create(&report).Error; err != nil {
		return nil, err
	}
	return &report, nil
}
