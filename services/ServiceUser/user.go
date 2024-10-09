package ServiceUser

import (
	"JayHonChat/models"
	"JayHonChat/result"
	"JayHonChat/services/dto"
	"JayHonChat/services/midware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var User dto.UserDTO
	if err := c.ShouldBind(&User); err != nil {
		result.Failture(result.APIcode.ShouldBindError, result.APIcode.GetMessage(result.APIcode.ShouldBindError), c, &err)
		return
	}
	if models.IsExited(User.Email) == false {
		models.AddUser(User, c)
		result.Success("RegisterSuccess", c)
	} else {
		result.Failture(result.APIcode.UserExist, result.APIcode.GetMessage(result.APIcode.UserExist), c, nil)
		return
	}
}

func Login(c *gin.Context) {
	Password := c.PostForm("password")
	Email := c.PostForm("email")
	var User dto.UserDTO
	User.Email = Email
	User.Password = Password
	if models.CheckUser(User, c) {
		result.Success("LoginSuccess", c)
	} else {
		result.Failture(result.APIcode.PasswordError, result.APIcode.GetMessage(result.APIcode.PasswordError), c, nil)
		return
	}
}
func Logout(c *gin.Context) {
	midware.ClearAuthSession(c)
	c.Redirect(http.StatusFound, "/")
	return
}
func GetUserInfo(c *gin.Context) map[string]interface{} {
	return midware.GetSessionUserInfo(c)
}
