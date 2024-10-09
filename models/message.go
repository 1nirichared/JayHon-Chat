package models

import (
	"gorm.io/gorm"
	"sort"
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

func GetLimitMsg(roomId string, offsert int) []map[string]interface{} {
	var results []map[string]interface{}
	db := GetChatDB()
	db.Model(&Message{}).
		Select("message.*,user.username,user.avatar_id").
		Joins("INNER Join users on users.id=message.user_id").
		Where("message.room_id= " + roomId).
		Where("message.to_user_id=0").
		Order("message.id.desc").
		Offset(offsert).Limit(100).Find(&results)

	if offsert == 0 {
		sort.Slice(results, func(i, j int) bool { return results[i]["id"].(uint32) <= results[j]["id"].(uint32) })
	}
	return results
}

func GetLimitPrivateMsg(uid, toUId string, offset int) []map[string]interface{} {
	var results []map[string]interface{}
	db := GetChatDB()
	db.Model(&Message{}).
		Select("messages.*, users.username ,users.avatar_id").
		Joins("INNER Join users on users.id = messages.user_id").
		Where("(" +
			"(" + "messages.user_id = " + uid + " and messages.to_user_id=" + toUId + ")" +
			" or " +
			"(" + "messages.user_id = " + toUId + " and messages.to_user_id=" + uid + ")" +
			")").
		Order("messages.id desc").
		Offset(offset).
		Limit(100).
		Scan(&results)
	if offset == 0 {
		sort.Slice(results, func(i, j int) bool {
			return results[i]["id"].(uint32) < results[j]["id"].(uint32)
		})
	}
	return results
}
