package controllers

import (
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile 获取用户个人信息
func GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ResponseError(c, http.StatusUnauthorized, "未登录")
		return
	}
	user, err := services.GetUserProfile(userID.(uint))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "获取用户信息失败")
		return
	}
	utils.ResponseSuccess(c, user)
}

// UpdateProfile 更新用户个人信息
func UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.ResponseError(c, http.StatusUnauthorized, "未登录")
		return
	}
	var req struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	user, err := services.UpdateUserProfile(userID.(uint), req.Email, req.Phone)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "更新失败: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, user)
}
