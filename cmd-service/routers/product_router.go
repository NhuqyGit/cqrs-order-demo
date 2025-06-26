package routers

import (
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/handler"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, handler *handler.ProductHandler) {
	products := r.Group("/products")
	{
		products.POST("", handler.CreateProduct)
	}
}