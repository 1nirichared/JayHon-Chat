package main

import (
	"JayHonChat/conf"
	"JayHonChat/models"
	"JayHonChat/routers"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func init() {

	viper.SetConfigType("json") // 设置配置文件的类型

	if err := viper.ReadConfig(bytes.NewBuffer(conf.AppJsonConfig)); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}

	models.GetChatDB()
}
func main() {
	//关闭debug模式
	gin.SetMode(gin.ReleaseMode)
	port := viper.GetString(`app.port`)
	router := routers.InitRoute()
	log.Println("监听端口", "http://127.0.0.1:"+port)

	router.Run(":" + port)
}
