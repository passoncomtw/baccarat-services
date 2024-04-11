package main

import (
	"baccarat-services/m/pkg/baccarat-game-service/conf"
	"baccarat-services/m/pkg/baccarat-game-service/models"
	"baccarat-services/m/pkg/baccarat-game-service/services"
	"bytes"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("json") // 設置配置文件的類型
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")

	if err := viper.ReadConfig(bytes.NewBuffer(conf.AppJsonConfig)); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 讀取配置文件失敗致命錯誤
	}

	models.InitDB()
}

func main() {
	services.LoopPokerRound()
}
