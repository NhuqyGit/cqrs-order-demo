package main

import (
	"log"

	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/db"
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/handler"
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/repository"
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/routers"
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	router := gin.Default()


	//
	database := db.GetDB()
	productRepo := repository.NewProductRepo(database)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	routers.RegisterProductRoutes(router, productHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + "8080")
}