package example

import (
	"context"

	"github.com/bogi-lyceya-44/example-service/internal/app/models"
	desc "github.com/bogi-lyceya-44/example-service/internal/pb/api/example"
)

type exampleService interface {
	Get(ctx context.Context, ids []models.ID) ([]models.Example, error)
}

type Implementation struct {
	desc.UnimplementedExampleServiceServer

	exampleService exampleService
}

func New(
	exampleService exampleService,
) *Implementation {
	return &Implementation{
		exampleService: exampleService,
	}
}
