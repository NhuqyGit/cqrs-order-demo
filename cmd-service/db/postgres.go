package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	once sync.Once
)

func InitPostgres() {
	once.Do(func(){
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		var err error
	
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil{
			log.Fatal("Failed to connect to Postgres: ", err)
		}
		
		Migrate(db)
		log.Println("Connected to Postgres")
	})
}

func GetDB() *gorm.DB{
	if db == nil{
		InitPostgres()
	}
	return db
}


