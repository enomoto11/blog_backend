package main

import (
	"blog/ent"
	"blog/ent/migrate"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
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

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"gin": "root",
		})
	})

	router.Run()

	defer entClient.Close()
}
