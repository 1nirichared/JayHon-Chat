package models

import (
	"JayHonChat/result"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Room struct {
	gorm.Model
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	RoomName  string
	CreatedAt time.Time
	Creater   string
	Avatar    string
}

func CreateRoom(c *gin.Context) {

	roomName := c.PostForm("room_name")
	CreaterID := c.PostForm("ID")
	db := GetChatDB()
	var Creater User
	if err := db.Where("ID=?", CreaterID).First(&Creater).Error; err != nil {
		result.Failture(result.APIcode.SelectError, result.APIcode.GetMessage(result.APIcode.SelectError), c, &err)
		return
	}
	CreaterName := Creater.Username
	db.Create(&Room{
		CreatedAt: time.Now(),
		Creater:   CreaterName,
		RoomName:  roomName,
	})
}
