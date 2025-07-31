package example

import (
	"context"

	"github.com/bogi-lyceya-44/example-service/internal/app/models"
)

func (s *ExampleService) Get(
	ctx context.Context,
	ids []models.Id,
) ([]models.Example, error) {
	return s.storage.Get(ctx, ids)
}
