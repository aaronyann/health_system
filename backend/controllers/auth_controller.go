package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 用户登录，返回 JWT
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	user, err := services.Authenticate(req.Username, req.Password)
	if err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "生成 token 失败")
		return
	}
	utils.ResponseSuccess(c, gin.H{"token": token, "user": user})
}

// Register 用户注册
func Register(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, "参数错误")
		return
	}
	user, err := services.RegisterUser(&req)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "注册失败: "+err.Error())
		return
	}
	utils.ResponseSuccess(c, user)
}
