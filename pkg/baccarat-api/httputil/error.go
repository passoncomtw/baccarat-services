package httputil

import "github.com/gin-gonic/gin"

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPUnauthorizedError struct {
	Message string `json:"message" example:"Unauthorized"`
}
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTP404Error struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"route not found"`
}

type HTTP500Error struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Internal Server Error"`
}
