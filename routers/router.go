package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	if viper.GetString(`app.debug_mod`) == "false" {

	}

	return router

}
