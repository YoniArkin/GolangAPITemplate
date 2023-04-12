package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func getClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://0.0.0.0:27017").
		SetMaxPoolSize(100).
		SetMinPoolSize(10).
		SetConnectTimeout(20 * time.Second)

	return mongo.Connect(context.Background(), clientOptions)
}

func getCollection(collection string) (*mongo.Collection, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	result := client.Database("test").Collection(collection)
	return result, nil
}
