package api

import (
	"baccarat-services/m/pkg/baccarat-api/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var tokens []string

// jwt secret key
var jwtSecret = []byte("secret")

// custom claims
type Claims struct {
	Id      uint   `json:"id"`
	Account string `json:"account"`
	jwt.StandardClaims
}

type UserResponseDto struct {
	Id        uint      `json:"id" binding:"required" validate:"required" example:1`
	CreatedAt time.Time `json:"createdAt" validate:"lte" example:"2024-03-30 18:00:00"`
	Account   string    `json:"account" example:"test001"`
	Name      string    `json:"name" example:"test001"`
	Amount    int       `json:"amount" example:0`
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
		context.Next()
	} else {
		context.Abort()
		return
	}
}

type AuthResponseDto struct {
	Token string          `json:"token" binding:"required" validate:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTMsImFjY291bnQiOiJ0ZXN0MDAxIiwiYXVkIjoidGVzdDAwMSIsImV4cCI6MTcxMjA2NTEzMiwianRpIjoiMTcxMjA2MzEzMiIsImlhdCI6MTcxMjA2MzEzMiwiaXNzIjoiZ2luSldUIiwibmJmIjoxNzEyMDYzMTMzLCJzdWIiOiJ0ZXN0MDAxIn0.RnjzI0kGPu1klhoyjkvXXxlf8gI_7469qf52b9XBajM"`
	User  UserResponseDto `json:"account" binding:"required"`
}

type AuthRequestDto struct {
	Account  string `json:"account" binding:"required" validate:"required" example:"test001"`
	Password string `json:"password" binding:"required" validate:"required" example:"a12345678"`
}

// Get Auth
// @Summary auth
// @Description get auth token
//
//	@Param			login body		AuthRequestDto	true	"Login Body"
//
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} AuthResponseDto
//
//	@Failure		401	{object}	httputil.HTTPUnauthorizedError
//
// @Router /auth [post]
func GetAuth(context *gin.Context) {
	// validate request body
	var body struct {
		Account  string
		Password string
	}
	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userResult := models.GetUserByAccountAndPassword(body.Account, body.Password)

	if userResult == nil || userResult.ID == 0 {
		// incorrect account or password
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	now := time.Now()
	jwtId := strconv.FormatInt(now.Unix(), 10)
	// set claims and sign
	claims := Claims{
		Id:      userResult.ID,
		Account: body.Account,
		StandardClaims: jwt.StandardClaims{
			Audience:  body.Account,
			ExpiresAt: now.Add(2000 * time.Second).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(1 * time.Second).Unix(),
			Subject:   body.Account,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": UserResponseDto{
			Id:        userResult.ID,
			Name:      userResult.Name,
			Account:   userResult.Account,
			CreatedAt: userResult.CreatedAt,
		},
	})
}
