package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameOrder struct {
	gorm.Model
	Id           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserId       uint      `gorm:"foreignKey:ID"`
	GameId       uint      `gorm:"foreignKey:ID"`
	GameRoundId  uint
	BetAmount    int        `gorm:"comment: '下注額度'"`
	BetType      int        `gorm:"comment: '下注類別: 0: 莊閒和, 1: 點數'"`
	WinAmount    int        `gorm:"comment: '派彩結果(贏或是輸)'"`
	Odds         int        `gorm:"comment: '派彩賠率'"`
	ConcludeTime *time.Time `gorm:"comment: '注單結算時間'"`
	// GameRound    GameRound  `gorm:"foreignKey:GameRoundId"`
	Game Game `gorm:"foreignKey:GameId"`
	User User `gorm:"foreignKey:UserId"`
}

type CreateGameOrderRequestDto struct {
	gameRoundId uint
	// userId       uint
	// gameId       uint
	betAmount    int
	betType      int
	winAmount    int
	odds         int
	concludeTime *time.Time
	User         User
	Game         Game
}

func GetGameOrdersByGameRoundId(id uint) *[]GameOrder {
	var gameOrders *[]GameOrder
	DBConnection.Find(&gameOrders, "game_round_id = ?", id)

	return gameOrders
}

func CreateGameOrder(values *CreateGameOrderRequestDto) GameOrder {
	var gameOrder GameOrder
	gameOrder.BetAmount = values.betAmount
	gameOrder.BetType = values.betType
	gameOrder.WinAmount = values.winAmount
	gameOrder.Odds = values.odds
	gameOrder.ConcludeTime = values.concludeTime

	gameOrder.GameRoundId = values.gameRoundId
	gameOrder.User = values.User
	gameOrder.Game = values.Game
	// DBConnection.Create(&gameOrder)
	return gameOrder
}
