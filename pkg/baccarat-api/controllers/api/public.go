package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary ping
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} resource data
// @Router /resource [get]
func GetPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Home godoc
// @Summary home
// @Schemes
// @Description home
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} resource data
// @Router /home [get]
func GetHome(context *gin.Context) {
	bearerToken := context.Request.Header.Get("Authorization")
	reqToken := strings.Split(bearerToken, " ")[1]
	for _, token := range tokens {
		if token == reqToken {
			context.JSON(http.StatusOK, gin.H{
				"data": "resource data",
			})
			return
		}
	}
	context.JSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
}

func GetResource(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"data": "resource data",
	})
}
