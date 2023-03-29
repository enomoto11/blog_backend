package controller

import (
	"blog/api/controller/request"
	"blog/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type PostController interface {
	RegisterHandlers(r gin.IRouter)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &postController{
		postService,
	}
}

func (c *postController) RegisterHandlers(r gin.IRouter) {
	r.POST("post/new", c.createPost)
}

func (c *postController) createPost(ctx *gin.Context) {
	var requestBody request.POSTPostRequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.postService.CreatePost(ctx, requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success in creating new post",
		"result":  result,
	})
}
