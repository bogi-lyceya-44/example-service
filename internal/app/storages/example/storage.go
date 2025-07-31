package example

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExampleStorage struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *ExampleStorage {
	return &ExampleStorage{
		pool: pool,
	}
}
