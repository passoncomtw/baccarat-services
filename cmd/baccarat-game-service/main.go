package main

import (
	"baccarat-services/m/pkg/baccarat-game-service/conf"
	"baccarat-services/m/pkg/baccarat-game-service/models"
	"baccarat-services/m/pkg/baccarat-game-service/services"
	"bytes"
	"fmt"
	"log"

	"github.com/sacOO7/socketcluster-client-go/scclient"
	"github.com/spf13/viper"
)

func startCode(client scclient.Client) {
	// start writing your code from here
	// All emit, receive and publish events
	fmt.Println("start coding")
	client.EmitAck("login", "test123", func(eventName string, error, data interface{}) {
		fmt.Println("eventName: %s", eventName)
		fmt.Println("error: %v", error)
		fmt.Println("data: %v", data)
	})
}

func onConnect(client scclient.Client) {
	fmt.Println("Connected to server")
}

func onDisconnect(client scclient.Client, err error) {
	fmt.Printf("Error: %s\n", err.Error())
}

func onConnectError(client scclient.Client, err error) {
	fmt.Printf("Error: %s\n", err.Error())
}

func onSetAuthentication(client scclient.Client, token string) {
	fmt.Println("Auth token received :", token)
	// setAuthToken
}

func onAuthentication(client scclient.Client, isAuthenticated bool) {
	fmt.Println("Client authenticated :", isAuthenticated)
	go startCode(client)
}

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
	url := viper.GetString(`sccserver.url`)
	client := scclient.New(url)
	client.SetBasicListener(onConnect, onConnectError, onDisconnect)
	client.SetAuthenticationListener(onSetAuthentication, onAuthentication)
	go client.Connect()

	services.LoopPokerRound()
}
