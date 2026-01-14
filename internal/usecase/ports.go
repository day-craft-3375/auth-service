package usecase

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type AppContext interface {
	Context() context.Context
	Transaction() *sqlx.Tx
}

type AppContextCancelFunc func()

type AppContextManager interface {
	CreateContext(parent context.Context) (AppContext, AppContextCancelFunc)
}

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Error(msg string, err error, args ...any)
	Warn(msg string, args ...any)
}

type IDGenerator interface {
	NewID() string
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) (bool, error)
}
