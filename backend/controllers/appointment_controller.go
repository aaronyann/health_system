package controllers

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAppointments 获取预约列表
func GetAppointments(c *gin.Context) {
	appointments, err := services.GetAppointments()
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取预约列表失败")
		return
	}
	utils.ResponseSuccess(c, appointments)
}

// CreateAppointment 创建预约
func CreateAppointment(c *gin.Context) {
	var req struct {
		PatientID       uint   `json:"patient_id" binding:"required"`
		DoctorID        uint   `json:"doctor_id" binding:"required"`
		DepartmentID    uint   `json:"department_id" binding:"required"`
		AppointmentTime string `json:"appointment_time" binding:"required"`
		Notes           string `json:"notes"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	appointment, err := services.CreateAppointment(req.PatientID, req.DoctorID, req.DepartmentID, req.AppointmentTime, req.Notes)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "创建预约失败: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, appointment)
}

// GetAppointmentByID 获取预约详情
func GetAppointmentByID(c *gin.Context) {
	id := c.Param("id")
	appointment, err := services.GetAppointmentByID(id)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取预约详情失败")
		return
	}
	utils.ResponseSuccess(c, appointment)
}
