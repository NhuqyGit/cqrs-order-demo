package repository

import (
	"context"
	"time"

	"github.com/NhuqyGit/cqrs-order-demo/query-service/db"
	"github.com/NhuqyGit/cqrs-order-demo/query-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]models.Product, error)
	Create(ctx context.Context, Product models.Product) error
}

type productRepo struct {
	col *mongo.Collection
}

func NewProductRepository(client *mongo.Client) ProductRepository {
	return &productRepo{
		col: db.GetMongoCollection(client, "testdb", "Products"),
	}
}

func (r *productRepo) GetAll(ctx context.Context) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var Products []models.Product
	for cursor.Next(ctx) {
		var u models.Product
		if err := cursor.Decode(&u); err != nil {
			return nil, err
		}
		Products = append(Products, u)
	}
	return Products, nil
}

func (r *productRepo) Create(ctx context.Context, Product models.Product) error {
	_, err := r.col.InsertOne(ctx, Product)
	return err
}