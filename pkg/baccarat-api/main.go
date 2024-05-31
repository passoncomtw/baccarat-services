package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"baccarat-services/m/pkg/baccarat-api/conf"
	"baccarat-services/m/pkg/baccarat-api/controllers"
	"baccarat-services/m/pkg/baccarat-api/models"

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

// @termsOfService  http://swagger.io/terms/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	url := viper.GetString(`sccserver.url`)
	client := scclient.New(url)
	client.SetBasicListener(onConnect, onConnectError, onDisconnect)
	client.SetAuthenticationListener(onSetAuthentication, onAuthentication)
	go client.Connect()

	routersInit := controllers.InitRouter()
	port := viper.GetString(`app.port`)

	log.Println("監聽端口", "http://127.0.0.1:"+port)

	http.ListenAndServe(":"+port, routersInit)
}
