package pkg

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func MongoClient(mongo_uri string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	return client, nil
}
