package models

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func InitDB() *gorm.DB {
	dsn := viper.GetString(`database.dsn`)
	DBConnection, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	DBConnection.AutoMigrate(
		&Game{},
		&Lobby{},
	)

	DBConnection.AutoMigrate(
		&User{},
		&GameOrder{},
		&GameRound{},
	)
	return DBConnection
}
