package handler

import (
	"net/http"

	"github.com/NhuqyGit/cqrs-order-demo/query-service/models"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler{
	return &ProductHandler{service: service}
}
func (h *ProductHandler) GetProducts(c *gin.Context){
	products, error := h.service.GetAllProducts(c)
	if error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := h.service.CreateProduct(c, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}