package routers

import (
	"JayHonChat/services/ServiceUser"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	if viper.GetString(`app.debug_mod`) == "false" {

	}

	router.POST("/Register", ServiceUser.Register)
	router.POST("/login", ServiceUser.Login)

	return router

}
