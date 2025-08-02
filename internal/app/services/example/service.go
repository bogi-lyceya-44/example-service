package example

import (
	"context"

	"github.com/bogi-lyceya-44/example-service/internal/app/models"
)

type storage interface {
	Get(ctx context.Context, ids []models.ID) ([]models.Example, error)
}

type Service struct {
	storage storage
}

func New(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}
