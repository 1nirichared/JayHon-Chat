package websocket

import (
	"JayHonChat/result"
	"JayHonChat/services/safe"
	"JayHonChat/ws"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jianfengye/collection"
	"net/http"
	"strconv"
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
	Status int             `json:"status"`
	Data   msgData         `json:"data"`
	Conn   *websocket.Conn `json:"conn"`
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
	enterRooms = make(chan wsClient)
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

const msgTypeOnline = 1        //上线
const msgTypeOffline = 2       //离线
const msgTypeSend = 3          //发送消息
const msgTypeGetOnlineUser = 4 //获取再线用户列表
const msgTypePrivateChat = 5   //私聊

type GoServe struct {
	ws.ServeInterface
}

func Run(c *gin.Context) {
	//解决跨域问题
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
}

func Read(conn *websocket.Conn, done chan<- struct{}) {

	defer func() {

	}()
	for {

	}
}

func Write(done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		case r := <-enterRooms:
			HandleConnClient(r.Conn)
		case cl := <-sMsg:
			serverMsgStr, _ := json.Marshal(cl)
			switch cl.Status {
			case msgTypeOnline, msgTypeSend:
				notify(cl.Conn, string(serverMsgStr))
			}
		}
	}
}

func HandleConnClient(conn *websocket.Conn) {
	roomid, roomIdInt := getRoomId()
	objColl := collection.NewObjCollection(rooms[roomIdInt])
	//移除已存在相同UID的用户
	retColl := safe.Safety.Lock(func() interface{} {
		return objColl.Reject(func(item interface{}, key int) bool {
			if item.(wsClient).Uid == clientMsg.Data.Uid {
				chNotify <- 1
				item.(wsClient).Conn.WriteMessage(websocket.TextMessage, []byte(`{"status":-1,"data":[]}`))
				<-chNotify
				return true
			}
			return false
		})
	}).(collection.ICollection)
	//将新用户信息添加到房间
	retColl = safe.Safety.Lock(func() interface{} {
		return retColl.Append(wsClient{
			Conn:       conn,
			RemoteAddr: conn.RemoteAddr().String(),
			Uid:        clientMsg.Data.Uid,
			Username:   clientMsg.Data.Username,
			Roomid:     roomid,
			AvatarId:   clientMsg.Data.AvatarId,
		})
	}).(collection.ICollection)
	//更新房间用户集合
	interfaces, _ := retColl.ToInterfaces()
	rooms[roomIdInt] = interfaces
}

func getRoomId() (string, int) {
	roomid := clientMsg.Data.Roomid
	roomIdInt, err := strconv.Atoi(roomid)
	if err != nil {
		result.Failture(result.APIcode.AtoiError, result.APIcode.GetMessage(result.APIcode.AtoiError), nil, err)
	}
	return roomid, roomIdInt
}

// 统一消息发送
func notify(conn *websocket.Conn, msg string) {
	chNotify <- 1
	_, roomIdInt := getRoomId()
	connect := rooms[roomIdInt]
	for _, con := range connect {
		if con.(wsClient).RemoteAddr != conn.RemoteAddr().String() {
			con.(wsClient).Conn.WriteMessage(websocket.TextMessage, []byte(msg))
		}

	}
	<-chNotify
}
