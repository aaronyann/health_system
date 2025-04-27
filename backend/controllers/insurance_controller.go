package controllers

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetInsuranceInfo 获取保险信息，根据 query 中的 patient_id（或从上下文获取）
func GetInsuranceInfo(c *gin.Context) {
	patientID := c.Query("patient_id")
	info, err := services.GetInsuranceInfo(patientID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取保险信息失败")
		return
	}
	utils.ResponseSuccess(c, info)
}

// UpdateInsuranceInfo 更新保险信息
func UpdateInsuranceInfo(c *gin.Context) {
	var req struct {
		PatientID    uint   `json:"patient_id" binding:"required"`
		Provider     string `json:"provider" binding:"required"`
		PolicyNumber string `json:"policy_number" binding:"required"`
		ValidFrom    string `json:"valid_from" binding:"required"`
		ValidTo      string `json:"valid_to" binding:"required"`
		Coverage     string `json:"coverage"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	info, err := services.UpdateInsuranceInfo(req.PatientID, req.Provider, req.PolicyNumber, req.ValidFrom, req.ValidTo, req.Coverage)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "更新保险信息失败: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, info)
}
