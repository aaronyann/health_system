package services

import (
	"backend/models"
	"backend/utils"
	"gorm.io/gorm"
)

var DB *gorm.DB // 在 main.go 初始化后赋值

// Authenticate 用户登录验证
func Authenticate(username, password string) (*models.User, error) {
	var user models.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	// 此处使用 utils.CheckPasswordHash 检查密码（建议使用 bcrypt）
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

// RegisterUser 注册用户
func RegisterUser(user *models.User) (*models.User, error) {
	// 密码加密
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed
	if err := DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserProfile 获取用户个人信息
func GetUserProfile(userID uint) (*models.User, error) {
	var user models.User
	if err := DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserProfile 更新用户个人信息
func UpdateUserProfile(userID uint, email, phone string) (*models.User, error) {
	var user models.User
	if err := DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	user.Email = email
	user.Phone = phone
	if err := DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
