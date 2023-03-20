package controller

import (
	"blog/api/controller/request"
	"blog/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserController interface {
	CreateUserController(r *gin.Engine)
}

type userController struct {
	createTeamService service.UserService
}

func NewUserController(createTeamService service.UserService) UserController {
	return &userController{
		createTeamService,
	}
}

func (c *userController) CreateUserController(r *gin.Engine) {
	r.POST("user/new", func(ctx *gin.Context) {
		var requestBody request.CreateUserRequestBody
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()
		if err := validate.Struct(requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := c.createTeamService.CreateUser(ctx, requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "success in creating new user",
			"result":  result,
		})
	})
}
