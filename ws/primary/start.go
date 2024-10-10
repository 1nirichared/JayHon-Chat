package primary

import (
	"JayHonChat/ws"
	"JayHonChat/ws/go_ws"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

// 定义serve的映射关系
var serveMap = map[string]ws.ServeInterface{
	"GoServe": &go_ws.GoServe{},
}

func Create() ws.ServeInterface {
	_type := viper.GetString("app.serve_type")
	serve, ok := serveMap[_type]
	if !ok || serve == nil {
		log.Printf("Error: No service found for serve_type: %s", _type)
		return nil
	}
	return serve
}

func Start(c *gin.Context) {
	Create().RunWs(c)
}

func OnlineUserCount() int {
	return Create().GetOnlineUserCount()
}

func OnlineRoomUserCount(roomId int) int {
	return Create().GetOnlineRoomUserCount(roomId)
}
