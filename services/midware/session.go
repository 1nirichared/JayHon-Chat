package midware

import (
	"JayHonChat/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

// 启用cookie会话管理
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(viper.GetString(`app.cookie_key`)))
	return sessions.Sessions("JayHon-Chat", store)
}

// 保存session信息（注册和登录时）
func SaveAutuSession(c *gin.Context, info interface{}) {
	session := sessions.Default(c)
	session.Set("uid", info)
	session.Save()
}

func GetSessionUserInfo(c *gin.Context) map[string]interface{} {
	session := sessions.Default(c)
	uid := session.Get("uid")
	data := make(map[string]interface{})
	if uid != nil {
		user := models.FindUserByField("id", uid.(string))
		data["uid"] = user.ID
		data["username"] = user.Username

	}
	return data
}

// 退出时清除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("uid"); sessionValue != nil {
		return false
	}
	return true
}

func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("uid")
		if sessionValue == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		uidInt, _ := strconv.Atoi(sessionValue.(string))
		if uidInt <= 0 {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.Set("uid", sessionValue)
		c.Next()
		return
	}
}
