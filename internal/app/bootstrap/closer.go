package bootstrap

import "github.com/bogi-lyceya-44/common/pkg/closer"

const (
	CloserGroupApp           = "app"
	CloserGroupConnections   = "connections"
	CloserGroupGlobalContext = "global context"
)

func InitCloser() {
	closer.AddGroups([]closer.Group{
		{
			Name:     CloserGroupApp,
			Priority: 0,
		},
		{
			Name:     CloserGroupConnections,
			Priority: 1,
		},
		{
			Name:     CloserGroupGlobalContext,
			Priority: 2,
		},
	}...)
}
