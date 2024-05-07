package models

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConnection *gorm.DB

func PrintMapKeyValue(data interface{}) {

	value := reflect.ValueOf(data)
	types := reflect.TypeOf(data)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("%v: %v\n", types.Field(i).Name, value.Field(i))
	}
}

// 初始化預設資料, 以方便提供測試
// 新增大廳, 新增玩家
// 新增 game 同時新增 round -> 新增 order
func initData() {
	// create lobby
	// lobbyRequestDto := CreateLobbyRequestDto{name: "Vi001"}
	// CreateLobby(&lobbyRequestDto)

	// create user
	// baseUserDto := BaseUserDto{Name: "vincent", Account: "admin", Password: "123456", Amount: 123456}
	// CreateUser(baseUserDto)

	// create game
	// gameRequestDto := CreateGameRequestDto{redCardIndex: 13}
	// CreateGame(gameRequestDto)

	// create order
	// user := GetUserFirst()
	// round := GetGameRoundFirst()
	// game := GetGameFirst()
	// gameOrderRequestDto := CreateGameOrderRequestDto{gameRoundId: round.ID, User: *user, Game: *game}
	// order := CreateGameOrder(&gameOrderRequestDto)

	// DBConnection.Save(&order)

	// preload
	// var gameOrders []GameOrder
	// res := DBConnection.Preload("Game").Find(&gameOrders)
	// if res.Error != nil {
	// 	fmt.Println(res.Error)
	// }
	// for _, v := range gameOrders {
	// 	PrintMapKeyValue(v)
	// }
}
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

	initData()
	return DBConnection
}
