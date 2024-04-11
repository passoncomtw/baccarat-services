package models

import (
	"time"

	"gorm.io/gorm"
)

type GameOrder struct {
	gorm.Model
	Id           string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserId       uint       `gorm:"comment: 'refrence: user.id'"`
	GameRoundId  uint       `gorm:"comment: 'refrence: game_round.id'"`
	BetAmount    int        `gorm:"comment: '下注額度'"`
	BetType      int        `gorm:"comment: '下注類別: 0: 莊閒和, 1: 點數'"`
	WinAmount    int        `gorm:"comment: '派彩結果(贏或是輸)'"`
	Odds         int        `gorm:"comment: '派彩賠率'"`
	ConcludeTime *time.Time `gorm:"comment: '注單結算時間'"`
	// GameRound    GameRound  `gorm:"foreignKey:Id"`
	// User User `gorm:"foreignKey:Id"`
}

type CreateGameOrderRequestDto struct {
	gameRoundId  uint
	userId       uint
	betAmount    int
	betType      int
	winAmount    int
	odds         int
	concludeTime *time.Time
}

func createGameOrder(values *CreateGameOrderRequestDto) GameOrder {
	var gameOrder GameOrder
	gameOrder.BetAmount = values.betAmount
	gameOrder.BetType = values.betType
	gameOrder.WinAmount = values.winAmount
	gameOrder.Odds = values.odds
	gameOrder.ConcludeTime = values.concludeTime
	DBConnection.Create(&gameOrder)
	return gameOrder
}
