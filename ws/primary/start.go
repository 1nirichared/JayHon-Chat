package primary

import (
	"JayHonChat/ws"
	"JayHonChat/ws/go_ws"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 定义serve的映射关系
var serveMap = map[string]ws.ServeInterface{
	"GoServe": &go_ws.GoServe{},
}

func Create() ws.ServeInterface {
	_type := viper.GetString("app.serve_type")
	return serveMap[_type]
}

func Start(c *gin.Context) {
	Create().RunWs(c)
}

func OnlineUserCount(c *gin.Context) int {
	return Create().GetOnlineUserCount()
}

func OnlineRoomUserCount(roomId int) int {
	return Create().GetOnlineRoomUserCount(roomId)
}
