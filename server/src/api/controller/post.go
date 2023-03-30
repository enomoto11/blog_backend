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
	r.POST("/post/new", c.createPost)
	r.GET("/posts", c.findAllPosts)
	r.GET("/posts/category/:id", c.findByCategoryID)
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

	res := createdPostResponse{
		ID:         result.GetID(),
		Title:      result.GetTitle(),
		Body:       result.GetBody(),
		CategoryID: result.GetCategoryID(),
		UserID:     result.GetUserID(),
	}

	ctx.JSON(http.StatusCreated, res)
}

func (c *postController) findAllPosts(ctx *gin.Context) {
	results, err := c.postService.FindAllPosts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res allpostsResponse
	for _, result := range results {
		res = append(res, post{
			ID:         result.GetID(),
			Title:      result.GetTitle(),
			Body:       result.GetBody(),
			CategoryID: result.GetCategoryID(),
			UserID:     result.GetUserID(),
		})
	}

	ctx.JSON(http.StatusCreated, res)
}

func (c *postController) findByCategoryID(ctx *gin.Context) {
	results, err := c.postService.FindByCategoryID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res allpostsResponse
	for _, result := range results {
		res = append(res, post{
			ID:         result.GetID(),
			Title:      result.GetTitle(),
			Body:       result.GetBody(),
			CategoryID: result.GetCategoryID(),
			UserID:     result.GetUserID(),
		})
	}

	ctx.JSON(http.StatusCreated, res)
}
