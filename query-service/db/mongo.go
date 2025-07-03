package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	once sync.Once
)

func InitMongoDB() {
	once.Do(func(){
		uri := os.Getenv("MONGO_URI") // mongodb://localhost:27017
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var err error
		mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal("Cannot connect MongoDB:", err)
		}

		log.Println("Connected to MongoDB!")	
	})
}

func GetMongoClient() *mongo.Client{
	if mongoClient == nil{
		InitMongoDB()
	}
	return mongoClient
}

func GetMongoCollection(client *mongo.Client, db, coll string) *mongo.Collection{
	return client.Database(db).Collection(coll)
}

func GetCollection(db, coll string) *mongo.Collection {
	return GetMongoCollection(GetMongoClient(), db, coll)
}