package models

import (
	"JayHonChat/result"
	"JayHonChat/services/dto"
	"JayHonChat/services/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func IsExited(email string) bool {
	db := GetChatDB()
	var user User
	db.Where("email=?", email).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}
func AddUser(user dto.UserDTO, c *gin.Context) {
	db := GetChatDB()
	PassWord := user.Password
	EncryptedPassword, err := helper.Bcrypt(PassWord)
	if err != nil {
		result.Failture(result.APIcode.EncryptionFailed, result.APIcode.GetMessage(result.APIcode.EncryptionFailed), c, nil)
	}
	User := &User{
		Username: user.Username,
		Password: EncryptedPassword,
		Email:    user.Email,
	}
	db.Create(&User)
}

func CheckUser(user dto.UserDTO, c *gin.Context) bool {
	db := GetChatDB()
	var User User
	db.Where("email=?", user.Email).First(&User)
	if User.ID == 0 {
		result.Failture(result.APIcode.UserNotExist, result.APIcode.GetMessage(result.APIcode.UserExist), c, nil)
	}
	if !helper.CheckPasswordHash(user.Password, User.Password) {

		return false
	}
	return true
}
