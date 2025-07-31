package bootstrap

import (
	"context"
	"log"

	"github.com/bogi-lyceya-44/common/pkg/closer"
)

func InitGlobalContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	closer.AddCallback(
		CloserGroupGlobalContext,
		func() error {
			log.Print("cancelling context")
			cancel()
			return nil
		},
	)

	return ctx
}
