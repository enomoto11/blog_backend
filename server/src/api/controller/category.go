package controller

import (
	"blog/api/controller/request"
	"blog/api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CategoryController interface {
	RegisterHandlers(r gin.IRouter)
}

type categoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &categoryController{
		categoryService,
	}
}

func (c *categoryController) RegisterHandlers(r gin.IRouter) {
	r.POST("category/new", func(ctx *gin.Context) {
		var requestBody request.POSTCategoryRequestBody
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()
		if err := validate.Struct(requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := c.categoryService.CreateCategory(ctx, requestBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "success in creating new category",
			"result":  result,
		})
	})
}
