package mongo

import (
	"context"
	"github.com/YoniArkin/GolangAPITemplate/internal/db/mongo/dbmodels"
	"github.com/YoniArkin/GolangAPITemplate/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const collection string = "Greeting"

type GreetingDAOImpl struct{}

func (GreetingDAOImpl) AddGreet(ctx context.Context, greet *models.Greet) error {
	collection, err := getCollection(collection)
	if err != nil {
		return err
	}

	greet.InsertedTime = time.Now().UTC()
	dbGreet, err := dbmodels.ToDBModel(*greet)
	if err != nil {
		return err
	}

	doc, err := collection.InsertOne(ctx, dbGreet)
	if err != nil {
		return err
	}
	greet.ID = doc.InsertedID.(primitive.ObjectID).String()
	return nil
}

// GetGreetsByPage returns a page of greets
func (GreetingDAOImpl) GetGreetsByPage(ctx context.Context, pageNumber int, pageSize int) ([]models.Greet, error) {
	collection, err := getCollection(collection)
	if err != nil {
		return nil, err
	}

	skip := (pageNumber - 1) * pageSize
	findOptions := options.FindOptions{}
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageSize))
	cursor, err := collection.Find(ctx, bson.M{}, &findOptions)
	if err != nil {
		return nil, err
	}

	var greets []dbmodels.DBGreet
	if err = cursor.All(ctx, &greets); err != nil {
		return nil, err
	}

	var result []models.Greet
	for g := range greets {
		result = append(result, greets[g].ToDomainModel())
	}

	return result, nil
}

func (GreetingDAOImpl) GetGreetsCount(ctx context.Context) (int, error) {
	collection, err := getCollection(collection)
	if err != nil {
		return 0, err
	}

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
