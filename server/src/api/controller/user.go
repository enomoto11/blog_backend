package controller

import (
	"blog/api/controller/request"
	"blog/api/controller/response"
	"blog/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserController interface {
	CreateUserController(r *gin.Engine)
	FindAllUserController(r *gin.Engine)
}

type userController struct {
	teamService service.UserService
}

func NewUserController(teamService service.UserService) UserController {
	return &userController{
		teamService,
	}
}

func (c *userController) FindAllUserController(r *gin.Engine) {
	r.GET("users", func(ctx *gin.Context) {
		result, err := c.teamService.FindAllUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := make([]response.AllUserResponse, 0, len(result))

		ctx.JSON(http.StatusOK, response)
	})
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

		result, err := c.teamService.CreateUser(ctx, requestBody)
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
