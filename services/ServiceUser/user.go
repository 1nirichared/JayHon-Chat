package ServiceUser

import (
	"JayHonChat/models"
	"JayHonChat/result"
	"JayHonChat/services/dto"
	"JayHonChat/services/midware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	var User dto.UserDTO
	log.Printf("Raw Request Data: %v", c.Request.PostForm)
	if err := c.ShouldBind(&User); err != nil {
		result.Failture(result.APIcode.ShouldBindError, result.APIcode.GetMessage(result.APIcode.ShouldBindError), c, &err)
		return
	}
	if models.IsExited(User.Email) == false {
		models.AddUser(User, c)
		result.Success("RegisterSuccess", c)
		c.Redirect(http.StatusFound, "/")
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
	if status, user := models.CheckUser(User, c); status {
		models.SaveAvatarId(User.AvatarId, User)
		midware.SaveAutuSession(c, string(strconv.Itoa(int(user.ID))))
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})

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
