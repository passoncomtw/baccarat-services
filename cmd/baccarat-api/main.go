package main

import (
	"bytes"
	"log"
	"net/http"

	"baccarat-services/m/pkg/baccarat-api/conf"
	"baccarat-services/m/pkg/baccarat-api/controllers"
	"baccarat-services/m/pkg/baccarat-api/models"

	"github.com/spf13/viper"
)

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
	routersInit := controllers.InitRouter()
	port := viper.GetString(`app.port`)

	log.Println("監聽端口", "http://127.0.0.1:"+port)

	http.ListenAndServe(":"+port, routersInit)
}
