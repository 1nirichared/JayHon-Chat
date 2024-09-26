package main

import (
	"JayHonChat/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	//关闭debug模式
	gin.SetMode(gin.ReleaseMode)
	port := viper.GetString(`app.port`)
	router := routers.InitRoute()
	log.Println("监听端口", "http://127.0.0.1:"+port)

	http.ListenAndServe(":"+port, router)
}
