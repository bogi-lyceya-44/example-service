package main

import (
	"log"

	api "github.com/bogi-lyceya-44/example-service/internal/app/api/example"
	"github.com/bogi-lyceya-44/example-service/internal/app/bootstrap"
	service "github.com/bogi-lyceya-44/example-service/internal/app/services/example"
	storage "github.com/bogi-lyceya-44/example-service/internal/app/storages/example"
	"github.com/pkg/errors"
)

func main() {
	bootstrap.InitCloser()
	ctx := bootstrap.InitGlobalContext()

	// add working with config
	pool, err := bootstrap.InitPostgresPool(
		ctx,
		"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	)
	if err != nil {
		log.Fatal(errors.Wrap(err, "init pool"))
	}

	storage := storage.New(pool)
	service := service.New(storage)
	impl := api.New(service)

	if err := bootstrap.RunApp(ctx, impl); err != nil {
		log.Fatal(errors.Wrap(err, "running app"))
	}
}
