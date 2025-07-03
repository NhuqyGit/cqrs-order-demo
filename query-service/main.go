package main

import (
	"log"

	"github.com/NhuqyGit/cqrs-order-demo/query-service/db"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/handler"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/repository"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/routers"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	router := gin.Default()

	mongoClient := db.GetMongoClient()
	productRepository := repository.NewProductRepository(mongoClient)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	routers.RegisterProductRoutes(router, productHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong query service",
		})
	})

	router.Run(":" + "8080")
}