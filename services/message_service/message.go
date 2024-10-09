package message_service

import "JayHonChat/models"

func GetLimitMsg(roomId string, offset int) []map[string]interface{} {
	return models.GetLimitMsg(roomId, offset)
}

func GetLimitPrivateMsg(uid, toUid string, offset int) []map[string]interface{} {
	return models.GetLimitPrivateMsg(uid, toUid, offset)
}
