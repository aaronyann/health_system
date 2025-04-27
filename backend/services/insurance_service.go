package services

import (
	"backend/models"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// GetInsuranceInfo 根据患者ID获取保险信息
func GetInsuranceInfo(patientIDStr string) (*models.InsuranceInfo, error) {
	pid, err := strconv.Atoi(patientIDStr)
	if err != nil {
		return nil, errors.New("无效的患者ID")
	}
	var info models.InsuranceInfo
	if err := DB.Where("patient_id = ?", pid).First(&info).Error; err != nil {
		return nil, fmt.Errorf("查询保险信息失败: %v", err)
	}
	return &info, nil
}

// UpdateInsuranceInfo 更新保险信息
func UpdateInsuranceInfo(patientID uint, provider, policyNumber, validFromStr, validToStr, coverage string) (*models.InsuranceInfo, error) {
	validFrom, err := time.Parse("2006-01-02", validFromStr)
	if err != nil {
		return nil, errors.New("生效日期格式错误")
	}
	validTo, err := time.Parse("2006-01-02", validToStr)
	if err != nil {
		return nil, errors.New("失效日期格式错误")
	}
	var info models.InsuranceInfo
	if err := DB.Where("patient_id = ?", patientID).First(&info).Error; err != nil {
		// 未找到记录时创建新记录
		info = models.InsuranceInfo{
			PatientID:    patientID,
			Provider:     provider,
			PolicyNumber: policyNumber,
			ValidFrom:    validFrom,
			ValidTo:      validTo,
			Coverage:     coverage,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		if err := DB.Create(&info).Error; err != nil {
			return nil, err
		}
		return &info, nil
	}
	// 更新已有记录
	info.Provider = provider
	info.PolicyNumber = policyNumber
	info.ValidFrom = validFrom
	info.ValidTo = validTo
	info.Coverage = coverage
	info.UpdatedAt = time.Now()
	if err := DB.Save(&info).Error; err != nil {
		return nil, err
	}
	return &info, nil
}
