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
	RegisterHandlers(r gin.IRouter)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService,
	}
}

func (c *userController) RegisterHandlers(r gin.IRouter) {
	r.GET("users", c.findAllUsers)
	r.POST("user/new", c.createUser)
}

func (c *userController) findAllUsers(ctx *gin.Context) {
	result, err := c.userService.FindAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := convertGETUserModelsToAllUserResponse(result)

	ctx.JSON(http.StatusOK, res)
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

func (c *userController) createUser(ctx *gin.Context) {
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

	res := createdUserResponse{
		ID:        result.GetID(),
		FirstName: result.GetFirstName(),
		LastName:  result.GetLastName(),
		Email:     result.GetEmail(),
	}

	ctx.JSON(http.StatusCreated, res)
}
