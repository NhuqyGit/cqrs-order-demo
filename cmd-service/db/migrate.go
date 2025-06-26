package db

import (
	"log"

	"github.com/NhuqyGit/cqrs-order-demo/cmd-service/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration succeed")
}

