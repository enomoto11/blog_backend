package controller

import (
	"blog/api/controller/request"
	"blog/api/model"
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
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService,
	}
}

func (c *userController) FindAllUserController(r *gin.Engine) {
	r.GET("users", func(ctx *gin.Context) {
		result, err := c.userService.FindAllUsers(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response := convertGETUserModelsToAllUserResponse(result)

		ctx.JSON(http.StatusOK, response)
	})
}

func convertGETUserModelsToAllUserResponse(models []*model.GETUserModel) GETAllUserResponse {
	var response GETAllUserResponse
	for _, model := range models {
		user := getEachUser{
			ID:        model.GetID(),
			FirstName: model.GetFirstName(),
			LastName:  model.GetLastName(),
			Email:     model.GetEmail(),
		}

		response = append(response, user)
	}

	return response
}

func (c *userController) CreateUserController(r *gin.Engine) {
	r.POST("user/new", func(ctx *gin.Context) {
		var requestBody request.POSTUserRequestBody
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()
		if err := validate.Struct(requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := c.userService.CreateUser(ctx, requestBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "success in creating new user",
			"result":  result,
		})
	})
}
