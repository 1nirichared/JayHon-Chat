package ServiceUser

import (
	"JayHonChat/models"
	"JayHonChat/result"
	"JayHonChat/services/dto"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var User dto.UserDTO
	if err := c.ShouldBind(&User); err != nil {
		result.Failture(result.APIcode.ShouldBindError, result.APIcode.GetMessage(result.APIcode.ShouldBindError), c, &err)
	}
	if models.IsExited(User.Email) == false {
		models.AddUser(User, c)
		result.Success("RegisterSuccess", c)
	} else {
		result.Failture(result.APIcode.UserExist, result.APIcode.GetMessage(result.APIcode.UserExist), c, nil)
	}
}

func Login(c *gin.Context) {
	var User dto.UserDTO
	if err := c.ShouldBind(&User); err != nil {
		result.Failture(result.APIcode.ShouldBindError, result.APIcode.GetMessage(result.APIcode.ShouldBindError), c, &err)
	}
	if models.CheckUser(User, c) {
		result.Success("LoginSuccess", c)
	} else {
		result.Failture(result.APIcode.PasswordError, result.APIcode.GetMessage(result.APIcode.PasswordError), c, nil)
	}
}
