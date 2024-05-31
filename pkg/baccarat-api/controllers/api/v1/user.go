package apiv1

import (
	"baccarat-services/m/pkg/baccarat-api/controllers/api"
	"baccarat-services/m/pkg/baccarat-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponseDto struct {
	User api.UserResponseDto `json:"user"`
}

type UpdateUserRequestDto struct {
	Name string `json:"name" validate:"required" example:"test001" gorm:"type:varchar(40);"`
}

// registe User
// @Summary User
// @Description registe an user account
//
//	@Param	registe body	models.BaseUserDto	true	"Registe Body"
//
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} UserResponseDto
//
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTP404Error
//	@Failure		500	{object}	httputil.HTTP500Error
//
// @Router /users [post]
func RegisteUser(context *gin.Context) {
	// validate request body
	var body models.BaseUserDto

	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, result := models.CreateUser(body)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user": api.UserResponseDto{
			Id:        user.ID,
			CreatedAt: user.CreatedAt,
			Account:   user.Account,
			Name:      user.Name,
			Amount:    user.Amount,
		},
	})
}

// get User
// @Summary User
// @Description get an user account
// @Security BearerAuth
//
//	@Param			id	path		string	true	"User ID"
//
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} UserResponseDto
//
// @Failure		404	{object}	httputil.HTTP404Error
//
// @Router /api/v1/users/{id} [get]
func GetUser(context *gin.Context) {
	id := context.Param("id")

	userResult := models.GetUserById(id)

	if userResult == nil || userResult.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user": api.UserResponseDto{
			Id:        userResult.ID,
			CreatedAt: userResult.CreatedAt,
			Account:   userResult.Account,
			Name:      userResult.Name,
			Amount:    userResult.Amount,
		},
	})
}

// update User
// @Summary User
// @Description update an user name
// @Security BearerAuth
//
//	@Param			id	path		int	true	"User ID"	Format(int64)
//	@Param	registe body	UpdateUserRequestDto	true	"Update User Request Body"
//
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} UserResponseDto
//
// @Failure		404	{object}	httputil.HTTP404Error
//
// @Router /api/v1/users/{id} [put]
func UpdateUser(context *gin.Context) {
	id := context.Param("id")
	var body UpdateUserRequestDto

	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userResult := models.UpdateUserById(id, body.Name)

	if userResult == nil || userResult.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user": api.UserResponseDto{
			Id:        userResult.ID,
			CreatedAt: userResult.CreatedAt,
			Account:   userResult.Account,
			Name:      userResult.Name,
			Amount:    userResult.Amount,
		},
	})
}

// delete User
// @Summary User
// @Description delete an user
// @Security BearerAuth
//
//	@Param			id	path		int	true	"User ID"	Format(int64)
//
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} object
//
// @Router /api/v1/users/{id} [delete]
func DeleteUser(context *gin.Context) {
	id := context.Param("id")
	models.DeleteUserById(id)
	context.JSON(http.StatusOK, gin.H{})
}
