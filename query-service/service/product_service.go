package service

import (
	"context"

	"github.com/NhuqyGit/cqrs-order-demo/query-service/models"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/repository"
)

type ProductService interface{
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	CreateProduct(ctx context.Context, user models.Product) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *productService{
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts(ctx context.Context) ([]models.Product, error){
	return s.repo.GetAll(ctx)
}

func (s *productService) CreateProduct(ctx context.Context, product models.Product) error{
	return s.repo.Create(ctx, product)
}
