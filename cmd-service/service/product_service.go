package service

import (
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/models"
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepo
}

func NewProductService(productRepo *repository.ProductRepo) *ProductService{
	return &ProductService{productRepo: productRepo}
}

func (r *ProductService) CreateProductService(p *models.Product) error{
	return r.productRepo.CreateProduct(p)
} 