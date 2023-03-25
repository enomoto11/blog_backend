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

	controllers := setUpController(entClient)

	controllers.userController.CreateUserController(router)
	controllers.userController.FindAllUserController(router)
	controllers.categoryController.CreateCategoryController(router)
	controllers.postController.CreatePostController(router)

	return router, entClient
}

type initializedContrllers struct {
	userController     controller.UserController
	categoryController controller.CategoryController
	postController     controller.PostController
}

func setUpController(entClient *ent.Client) initializedContrllers {
	userRepo := repository.NewUserRepository(entClient)
	categoryRepo := repository.NewCategoryRepository(entClient)
	postRepo := repository.NewPostRepository(entClient)

	userService := service.NewUserService(userRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	postService := service.NewPostService(postRepo)

	userController := controller.NewUserController(userService)
	categoryController := controller.NewCategoryController(categoryService)
	postController := controller.NewPostController(postService)

	return initializedContrllers{
		userController,
		categoryController,
		postController,
	}
}
