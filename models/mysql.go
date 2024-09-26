package models

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := viper.GetString(`mysql.dsn`)
	ChatDB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return ChatDB
}

func GetChatDB() *gorm.DB {
	return initDB()
}
