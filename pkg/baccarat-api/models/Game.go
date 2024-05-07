package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	RedCardIndex int       `gorm:"default:0;comment: '紅卡位置'"`
	GameRound    GameRound `gorm:"foreignKey:GameId"`
}

type CreateGameRequestDto struct {
	redCardIndex int
}

func CreateGame(values CreateGameRequestDto) Game {
	var game Game
	game.RedCardIndex = values.redCardIndex

	// rounds
	gameRoundRequestDto := CreateGameRoundRequestDto{}
	round := CreateGameRound(&gameRoundRequestDto)

	// orders
	// orders := []GameOrder{{UserId: 1, GameRoundId: 1}}
	// game.GameOrders = orders

	game.GameRound = round

	DBConnection.Create(&game)
	return game
}

func GetGameById(id uint) *Game {
	var game *Game
	DBConnection.Where("id = ?", id).Find(&game)

	return game
}

func GetGameFirst() *Game {
	var game *Game
	DBConnection.First(&game)

	return game
}
