package dbmodels

import (
	"errors"
	"github.com/YoniArkin/GolangAPITemplate/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// DBGreet Greet is the database model for a Greeting
type DBGreet struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	InsertedTime time.Time          `bson:"insertedTime,omitempty"`
	UpdatedTime  time.Time          `bson:"updatedTime,omitempty"`
	Message      string             `bson:"message,omitempty"`
}

// ToDomainModel converts a DBGreet to a model.Greet
func (g DBGreet) ToDomainModel() models.Greet {
	return models.Greet{
		MetaData: models.MetaData{
			ID:           g.ID.String(),
			InsertedTime: g.InsertedTime,
			UpdatedTime:  g.UpdatedTime,
		},
		Message: g.Message,
	}
}

// ToDBModel converts a model.Greet to a DBGreet
func ToDBModel(g models.Greet) (*DBGreet, error) {
	if g.ID != "" && !primitive.IsValidObjectID(g.ID) {
		return nil, errors.New("invalid ID")
	}

	id, _ := primitive.ObjectIDFromHex(g.ID)
	return &DBGreet{
		ID:           id,
		InsertedTime: g.InsertedTime,
		UpdatedTime:  g.UpdatedTime,
		Message:      g.Message,
	}, nil
}
