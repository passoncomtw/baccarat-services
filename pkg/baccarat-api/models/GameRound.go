package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

type GameRound struct {
	gorm.Model
	GameId       uint            `gorm:"comment: 'game.id'"`
	GameResult   int             `gorm:"default:0;comment: '遊戲勝負, 0: 和局, -1: 莊贏, 1: 閒贏'"`
	BankerResult *postgres.Jsonb `gorm:"default:'{}';comment: '紀錄莊家卡牌資訊'"`
	PlayerResult *postgres.Jsonb `gorm:"default:'{}';comment: '紀錄玩家卡牌資訊'"`
	OddsResult   *postgres.Jsonb `gorm:"default:'{}';comment: '所有壓住區賠率'"`
	StartTime    time.Time       `gorm:"autoCreateTime"`
	EndTime      *time.Time
	GameOrders   []GameOrder `gorm:"foreignKey:GameRoundId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type GameRoundResultDto struct {
	amount int
	result []int
}

type CreateGameRoundRequestDto struct {
	gameId uint
	// gameResult int
}

func GetGameRoundById(id uint) *GameRound {
	var round *GameRound
	DBConnection.Where("id = ?", id).Find(&round)

	return round
}

func GetGameRoundFirst() *GameRound {
	var round *GameRound
	DBConnection.First(&round)

	return round
}

func CreateGameRound(values *CreateGameRoundRequestDto) GameRound {
	var gameRound GameRound
	gameRound.GameResult = 0
	gameRound.BankerResult = nil
	gameRound.PlayerResult = nil
	gameRound.OddsResult = nil
	gameRound.StartTime = time.Now()
	gameRound.EndTime = nil

	// Game
	// game := GetGameById(1)

	// gameRound.Game = *game
	// gameRound.GameOrders = orders

	// DBConnection.Create(&gameRound)
	return gameRound
}
