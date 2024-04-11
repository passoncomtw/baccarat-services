package controllers

import (
	"net/http"
	"strings"

	"baccarat-services/m/pkg/baccarat-api/controllers/api"
	apiv1 "baccarat-services/m/pkg/baccarat-api/controllers/api/v1"
	"baccarat-services/m/pkg/baccarat-api/docs"

	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files
var (
	tokens []string
	i      interface{}
)

// var tokens []string

// jwt secret key
var jwtSecret = []byte("secret")

// custom claims
type Claims struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

// validate JWT
func AuthRequired(context *gin.Context) {
	auth := context.GetHeader("Authorization")
	token := strings.Split(auth, "Bearer ")[1]
	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		context.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		context.Set("account", claims.Account)
		context.Set("role", claims.Role)
		context.Next()
	} else {
		context.Abort()
		return
	}
}

type SwaggerConfigDto struct {
	BasePath    string
	Version     string
	Host        string
	Title       string
	Description string
}

func InitRouter() *gin.Engine {

	basePath := viper.GetString("swagger.basePath")
	title := viper.GetString("swagger.title")
	host := viper.GetString("swagger.host")
	version := viper.GetString("swagger.version")
	description := viper.GetString("swagger.description")

	docs.SwaggerInfo.Title = title
	docs.SwaggerInfo.Description = description
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/ping", api.GetPing)
	router.GET("/home", api.GetHome)
	router.GET("/resource", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), api.GetResource)
	router.POST("/auth", api.GetAuth)
	router.POST("/users", apiv1.RegisteUser)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := router.Group("/api/v1")
	v1.Use(AuthRequired)
	{
		v1.GET("/member/profile", apiv1.GetProfile)
		v1.GET("/users/:id", apiv1.GetUser)
		v1.PUT("/users/:id", apiv1.UpdateUser)
		v1.DELETE("/users/:id", apiv1.DeleteUser)
	}

	return router
}
