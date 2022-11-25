package user

import (
	"context"

	"gitlab.com/tclaudel_ateme/hexagonal_architecture_golang/internal/user/core"
)

type Repository interface {
	Retriever
	Appender
	Shutdown
}

type Retriever interface {
	Get() ([]core.User, error)
}

type Appender interface {
	Create(user *core.User) error
}

type Shutdown interface {
	Shutdown(ctx context.Context) error
}
