package ping

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStore contains information required to establish a mondodb connection and read/write data
type MongoStore struct {
	MongoDBURL string
	Client     *mongo.Client
}

// NewMongoStore is a factory for MongoDB backed Document Storage
func NewMongoStore(mongoDBURL string) (*MongoStore, error) {
	mongoURI := fmt.Sprintf("mongodb://%v:27017", mongoDBURL)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &MongoStore{
		MongoDBURL: mongoDBURL,
		Client:     client,
	}, nil
}

// Update stores a document in a Mongo Document Store
func (ms *MongoStore) Update(url string, rc int, latency time.Duration, document string) (err error) {
	log.Printf("creating connection")
	collection := ms.Client.Database("openping").Collection("documents")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"url": url, "rc": rc, "document": document})
	if err != nil {
		return err
	}
	return nil
}
