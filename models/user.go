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
	AvatarId string `json:"avatar_id"`
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
	db = db.Debug()
	PassWord := user.Password
	EncryptedPassword, err := helper.Bcrypt(PassWord)
	if err != nil {
		result.Failture(result.APIcode.EncryptionFailed, result.APIcode.GetMessage(result.APIcode.EncryptionFailed), c, nil)
		return
	}
	User := &User{
		Username: user.Username,
		Password: EncryptedPassword,
		Email:    user.Email,
		AvatarId: user.AvatarId,
	}
	db.Create(&User)
}

func CheckUser(user dto.UserDTO, c *gin.Context) (bool, *User) {
	db := GetChatDB()
	var User User
	db.Where("email=?", user.Email).First(&User)
	if User.ID == 0 {
		result.Failture(result.APIcode.UserNotExist, result.APIcode.GetMessage(result.APIcode.UserNotExist), c, nil)
		return false, nil
	}
	if !helper.CheckPasswordHash(user.Password, User.Password) {

		return false, nil
	}
	return true, &User
}

func SaveAvatarId(AvatarId string, u User) User {

	u.AvatarId = AvatarId
	db := GetChatDB()
	db.Save(&u)
	return u
}

func FindUserByField(field, value string) User {
	var u User
	db := GetChatDB()
	if field == "id" || field == "email" {
		db.Where(field+"=?", value).First(&u)
	}
	return u
}
