package example

import (
	"context"

	"github.com/bogi-lyceya-44/example-service/internal/app/models"
)

type storage interface {
	Get(ctx context.Context, ids []models.Id) ([]models.Example, error)
}

type ExampleService struct {
	storage storage
}

func New(storage storage) *ExampleService {
	return &ExampleService{
		storage: storage,
	}
}
