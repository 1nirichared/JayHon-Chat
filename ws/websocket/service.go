package websocket

import (
	"JayHonChat/ws"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"net/http"
	"sync"
)

type wsClient struct {
	Conn       *websocket.Conn `json:"conn"`
	RemoteAddr string          `json:"remote_addr"`
	Uid        string          `json:"uid"`
	Username   string          `json:"username"`
	Roomid     string          `json:"roomid"`
	AvatarId   string          `json:"avatar_id"`
}

type msgData struct {
	Uid      string        `json:"uid"`
	Username string        `json:"username"`
	AvatarId string        `json:"avatar_id"`
	ToUid    string        `json:"to_uid"`
	Content  string        `json:"content"`
	ImageUrl string        `json:"image_url"`
	Roomid   string        `json:"roomid"`
	Count    int           `json:"count"`
	List     []interface{} `json:"list"`
	Time     int64         `json:"time"`
}

type msg struct {
	Stattus int             `json:"status"`
	Data    msgData         `json:"data"`
	Conn    *websocket.Conn `json:"conn"`
}

// 心跳检测
type pingStorage struct {
	Conn       *websocket.Conn `json:"conn"`
	RemoteAddr string          `json:"remote_addr"`
	Time       int64           `json:"time"`
}

var (
	wsUpgrader = websocket.Upgrader{}

	clientMsg = msg{}

	mutex = sync.Mutex{}

	//rooms = [roomCount + 1][]wsClients{}
	//值是存储在该房间内的 WebSocket 客户端集合（[]interface{}）。
	//这里的 interface{} 代表的是一个可以存储任何类型的空接口，
	//但实际存储的是 wsClients 结构体的实例。
	rooms = make(map[int][]interface{})
	//用于传递新加入房间的客户端连接。
	//它接收的是 wsClients 结构体，
	//这个结构体包含了客户端连接的信息（如 Conn、RemoteAddr、Uid 等
	enterRooms = make(chan wsClients)
	//用于传递服务器和客户端之间的消息（类型为 msg 结构体）。
	//当服务器或客户端发消息时，消息会通过该通道在不同的 goroutine 之间传递。
	sMsg = make(chan msg)
	//这是一个 chan 通道，用于传递需要断开连接的客户端（类型为 *websocket.Conn）。
	//当客户端连接断开或超时时，连接会通过该通道传递，并触发相关的断开处理。
	offline = make(chan *websocket.Conn)
	//这是一个带有缓冲区大小为 1 的 chan 通道，
	//通常用于在并发操作中起到同步或通知的作用。
	//它在一些操作（如消息通知、写入）之间起到阻塞或控制的作用，避免多个 goroutine 并发修改同一资源。
	//在写消息时，它会阻塞其他 goroutine，
	//确保同一时间只有一个消息被发送，
	//避免 concurrent write to websocket connection 错误。
	chNotify = make(chan int, 1)
	//pingMap 是一个 []interface{} 切片，用于存储心跳检测的客户端连接信息。
	//每个连接的信息被存储为 pingStorage 结构体，记录了连接的 WebSocket 对象、远程地址和最近一次心跳的时间。
	//它用于处理 WebSocket 心跳机制，定期检查并清理超时的客户端连接。
	pingMap []interface{}

	clientMsgLock = sync.Mutex{}
	clientMsgData = clientMsg // 临时存储 clientMsg 数据
)

type GoServe struct {
	ws.ServeInterface
}

func Run(c *gin.Context) {
	wsUpgrader.che := func(r *http.Request) bool { return true }
}
