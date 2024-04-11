package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Id           uint        `gorm:"primaryKey;comment: '主鍵'"`
	RedCardIndex int         `gorm:"default:0;comment: '紅卡位置'"`
	GameOrders   []GameOrder `gorm:"foreignKey:GameRoundId"`
}

type CreateGameRequestDto struct {
	redCardIndex int
}

func CreateGame(values CreateGameRequestDto) Game {
	var game Game
	game.RedCardIndex = values.redCardIndex
	DBConnection.Create(&game)
	return game
}
