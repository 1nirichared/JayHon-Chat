package routers

import (
	"JayHonChat/controller"
	"JayHonChat/services/midware"
	"JayHonChat/static"
	"JayHonChat/ws/primary"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	if viper.GetString(`app.debug_mod`) == "false" {
		//live模式打包用
		router.StaticFS("/static", http.FS(static.EmbedStatic))
	} else {
		router.StaticFS("/static", http.Dir("static"))
	}

	sr := router.Group("/", midware.EnableCookieSession())
	{
		sr.GET("/", controller.Index)
		sr.POST("/login", controller.Login)
		sr.GET("/RegisterPage", controller.RegisterPage)
		sr.POST("/register", controller.Register)
		sr.GET("/logout", controller.Logout)
		sr.GET("/ws", primary.Start)
		authoriezd := sr.Group("/", midware.AuthSessionMiddle())
		{
			authoriezd.GET("/home", controller.Home)
			authoriezd.GET("/room/:room_id", controller.Room)
			authoriezd.GET("/private-chat", controller.PrivateChat)
			authoriezd.POST("/img-kr-upload", controller.ImgKrUpload)
			authoriezd.GET("/pagination", controller.Pagination)
		}
	}

	return router

}
