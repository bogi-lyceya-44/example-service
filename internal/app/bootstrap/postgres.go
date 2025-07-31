package bootstrap

import (
	"context"
	"log"

	"github.com/bogi-lyceya-44/common/pkg/closer"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

func InitPostgresPool(
	ctx context.Context,
	connectionString string,
) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "init postgres")
	}

	closer.AddCallback(
		CloserGroupConnections,
		func() error {
			log.Print("closing pool")
			pool.Close()
			return nil
		},
	)

	return pool, nil
}
