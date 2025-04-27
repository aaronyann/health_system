package controllers

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLabReports 获取检验报告列表
func GetLabReports(c *gin.Context) {
	reports, err := services.GetLabReports()
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取检验报告列表失败")
		return
	}
	utils.ResponseSuccess(c, reports)
}

// GetLabReportByID 获取检验报告详情
func GetLabReportByID(c *gin.Context) {
	id := c.Param("id")
	report, err := services.GetLabReportByID(id)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取检验报告详情失败")
		return
	}
	utils.ResponseSuccess(c, report)
}

// CreateLabReport 创建检验报告
func CreateLabReport(c *gin.Context) {
	var req struct {
		AppointmentID uint   `json:"appointment_id" binding:"required"`
		PatientID     uint   `json:"patient_id" binding:"required"`
		LabName       string `json:"lab_name" binding:"required"`
		TestType      string `json:"test_type" binding:"required"`
		Results       string `json:"results" binding:"required"`
		ResultsHash   string `json:"results_hash" binding:"required"`
		IssuedTime    string `json:"issued_time" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	report, err := services.CreateLabReport(req.AppointmentID, req.PatientID, req.LabName, req.TestType, req.Results, req.ResultsHash, req.IssuedTime)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "创建检验报告失败: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, report)
}
