package controllers

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPrescriptionList 获取处方列表
func GetPrescriptionList(c *gin.Context) {
	prescriptions, err := services.GetPrescriptionList()
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取处方列表失败")
		return
	}
	utils.ResponseSuccess(c, prescriptions)
}

// GetPrescriptionDetail 获取处方详情
func GetPrescriptionDetail(c *gin.Context) {
	id := c.Param("id")
	prescription, err := services.GetPrescriptionByID(id)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取处方详情失败")
		return
	}
	utils.ResponseSuccess(c, prescription)
}

// CreatePrescription 创建处方
func CreatePrescription(c *gin.Context) {
	var req struct {
		AppointmentID uint   `json:"appointment_id" binding:"required"`
		DoctorID      uint   `json:"doctor_id" binding:"required"`
		PatientID     uint   `json:"patient_id" binding:"required"`
		Medications   string `json:"medications" binding:"required"`
		Instructions  string `json:"instructions"`
		IssuedTime    string `json:"issued_time" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	prescription, err := services.CreatePrescription(req.AppointmentID, req.DoctorID, req.PatientID, req.Medications, req.Instructions, req.IssuedTime)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "创建处方失败: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, prescription)
}
