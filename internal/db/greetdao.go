package db

import (
	"context"
	"github.com/YoniArkin/GolangAPITemplate/internal/models"
)

type GreetDAO interface {
	AddGreet(ctx context.Context, greet *models.Greet) error
	GetGreetsByPage(ctx context.Context, pageNumber int, pageSize int) ([]models.Greet, error)
	GetGreetsCount(ctx context.Context) (int, error)
}
