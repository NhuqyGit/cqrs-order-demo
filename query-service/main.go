package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/NhuqyGit/cqrs-order-demo/query-service/db"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/event/consumer"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/handler"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/repository"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/routers"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Init Mongo
	mongoClient := db.GetMongoClient()
	productRepository := repository.NewProductRepository(mongoClient)
	productService := service.NewProductService(productRepository)

	// Init RabbitMQ consumer
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@rabbitmq:5672/"
	}

	eventConsumer, err := consumer.NewEventConsumer(rabbitURL)
	if err != nil {
		log.Fatal("Failed to connect RabbitMQ:", err)
	}
	defer eventConsumer.Close()

	// Listen for ProductCreatedEvent
	err = eventConsumer.StartProductCreatedConsumer(productService)
	if err != nil {
		log.Fatal("Failed to start RabbitMQ consumer:", err)
	}

	// Gin routes
	router := gin.Default()
	productHandler := handler.NewProductHandler(productService)
	routers.RegisterProductRoutes(router, productHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong query service",
		})
	})

	// Run server and consumer parallel
	log.Println("query-service is running on :8080 and listening for events...")
	go router.Run(":8080")

	// Keep main alive (consumer uses goroutine)
	select {}
}
