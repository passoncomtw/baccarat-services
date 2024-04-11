package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	RedCardIndex int `gorm:"default:0;comment: '紅卡位置'"`
}

type CreateGameRequestDto struct {
	RedCardIndex int `json:"redCardIndex"`
}

func CreateGame(values CreateGameRequestDto) Game {
	var game Game
	game.RedCardIndex = values.RedCardIndex
	DBConnection.Create(&game)
	return game
}
