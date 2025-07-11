package service

import (
	"log"

	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/event/publisher"
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/models"
	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepo
	eventPublisher *publisher.EventPublisher
}

func NewProductService(productRepo *repository.ProductRepo, eventPublisher *publisher.EventPublisher) *ProductService{
	return &ProductService{productRepo: productRepo, eventPublisher: eventPublisher}
}

func (r *ProductService) CreateProductService(p *models.Product) error{
	err := r.productRepo.CreateProduct(p)
	if err != nil{
		return err
	}

	// Create message and send into RabbitMQ
	productEvent := publisher.ProductCreatedEvent{
		ID:       p.ID,
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
	}

	err = r.eventPublisher.PublishProductCreated(productEvent)
	if err != nil {
		return err
	}
	log.Println("Publisher product event success")

	return nil
} 