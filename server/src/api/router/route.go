package router

import (
	"blog/api/controller"
	"blog/api/repository"
	"blog/api/service"
	"blog/ent"
	"blog/ent/migrate"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func InitControllers() (*gin.Engine, *ent.Client) {
	router := gin.Default()

	path := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	entClient, err := ent.Open("mysql", path)
	if err != nil {
		log.Fatalf("failed connect to mysql: %v", err)
	}

	ctx := context.Background()
	if err := entClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	controllers := setUpControllers(entClient)
	for _, controller := range controllers {
		controller.RegisterHandlers(router)
	}

	return router, entClient
}

type Controller interface {
	RegisterHandlers(r gin.IRouter)
}

func setUpControllers(entClient *ent.Client) []Controller {
	userRepo := repository.NewUserRepository(entClient)
	categoryRepo := repository.NewCategoryRepository(entClient)
	postRepo := repository.NewPostRepository(entClient)

	userService := service.NewUserService(userRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	postService := service.NewPostService(
		postRepo,
		userRepo,
		categoryRepo,
	)

	userController := controller.NewUserController(userService)
	categoryController := controller.NewCategoryController(categoryService)
	postController := controller.NewPostController(postService)

	return []Controller{
		userController,
		categoryController,
		postController,
	}
}
