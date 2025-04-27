package controllers

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealthRecords 获取健康档案列表
func GetHealthRecords(c *gin.Context) {
	records, err := services.GetHealthRecords()
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取健康档案失败")
		return
	}
	utils.ResponseSuccess(c, records)
}

// CreateHealthRecord 创建健康档案并上链
func CreateHealthRecord(c *gin.Context) {
	var req struct {
		PatientName string `json:"patientName" binding:"required"`
		DataHash    string `json:"dataHash" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	txHash, err := services.CreateHealthRecord(req.PatientName, req.DataHash)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "创建健康档案失败: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, gin.H{"txHash": txHash})
}

// GetHealthRecordByID 获取单条健康档案详情
func GetHealthRecordByID(c *gin.Context) {
	id := c.Param("id")
	record, err := services.GetHealthRecordByID(id)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "查询失败")
		return
	}
	utils.ResponseSuccess(c, record)
}
