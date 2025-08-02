package example

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/bogi-lyceya-44/common/pkg/utils"
	"github.com/bogi-lyceya-44/example-service/internal/app/models"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s *Storage) Get(
	ctx context.Context,
	ids []models.Id,
) ([]models.Example, error) {
	sql, args, err := sq.
		Select(allColumns...).
		From(tableName).
		Where(map[string]any{
			"id": ids,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "building sql")
	}

	rows, err := s.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "fetching rows")
	}

	fetched, err := pgx.CollectRows(rows, pgx.RowToStructByName[Example])
	if err != nil {
		return nil, errors.Wrap(err, "collecting rows")
	}

	return utils.Map(
		fetched,
		func(item Example) models.Example {
			return models.Example{
				Id:       models.NewId(item.Id),
				FormedAt: item.FormedAt,
			}
		},
	), nil
}
