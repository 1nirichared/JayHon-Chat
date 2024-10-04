package models

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Message struct {
	gorm.Model
	ID        uint
	UserID    int
	TOUsrId   int
	RoomId    int
	Content   string
	ImageUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func SaveContent(value interface{}) (*Message, error) {
	var m Message
	m.UserID = value.(map[string]interface{})["user_id"].(int)
	m.TOUsrId = value.(map[string]interface{})["to_user_id"].(int)
	m.Content = value.(map[string]interface{})["content"].(string)
	roomIdStr := value.(map[string]interface{})["room_id"].(string)
	roomIdInt, _ := strconv.Atoi(roomIdStr)
	m.RoomId = roomIdInt
	if _, ok := value.(map[string]interface{})["image_url"].(time.Time); ok {
		m.ImageUrl = value.(map[string]interface{})["image_url"].(string)
	}
	db := GetChatDB()
	db.Create(&m)
	return &m, nil
}
