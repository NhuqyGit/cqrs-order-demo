package repository

import (
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/models"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
    return &ProductRepo{db: db}
}

func (r *ProductRepo) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}