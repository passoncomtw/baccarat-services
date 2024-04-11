package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type GameRound struct {
	gorm.Model
	Id           uint            `gorm:"primaryKey;comment: '主鍵'"`
	GameId       uint            `gorm:"comment: 'game.id'"`
	GameResult   int             `gorm:"default:0;comment: '遊戲勝負, 0: 和局, -1: 莊贏, 1: 閒贏'"`
	BankerResult *postgres.Jsonb `gorm:"default:'{}';comment: '紀錄莊家卡牌資訊'"`
	PlayerResult *postgres.Jsonb `gorm:"default:'{}';comment: '紀錄玩家卡牌資訊'"`
	OddsResult   *postgres.Jsonb `gorm:"default:'{}';comment: '所有壓住區賠率'"`
	StartTime    time.Time       `gorm:"autoCreateTime"`
	EndTime      *time.Time
	Game         Game        `gorm:"foreignKey:Id"`
	GameOrders   []GameOrder `gorm:"foreignKey:GameRoundId"`
}

type GameRoundResultDto struct {
	amount int
	result []int
}

type CreateGameRoundRequestDto struct {
	gameId uint
	// gameResult int
}

func createGameRound(values *CreateGameRoundRequestDto) GameRound {
	var gameRound GameRound
	gameRound.GameResult = 0
	gameRound.BankerResult = nil
	gameRound.PlayerResult = nil
	gameRound.OddsResult = nil
	gameRound.StartTime = time.Now()
	gameRound.EndTime = nil

	DBConnection.Create(&gameRound)
	return gameRound
}
