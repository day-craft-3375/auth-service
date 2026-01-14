package appcontext

import (
	"context"
	"time"

	"github.com/day-craft-3375/auth-service/internal/usecase"
	"github.com/jmoiron/sqlx"
)

type appContext struct {
	ctx context.Context
}

// newAppContext создаёт контекст приложения
func newAppContext(parent context.Context, timeout time.Duration) (usecase.AppContext, usecase.AppContextCancelFunc) {
	ctx, cancel := context.WithTimeout(parent, timeout)

	return &appContext{
		ctx: ctx,
	}, usecase.AppContextCancelFunc(cancel)
}

// Context возвращает контекст
func (a *appContext) Context() context.Context {
	return a.ctx
}

// Transaction возвращает транзакцию базы данных
func (a *appContext) Transaction() *sqlx.Tx {
	ts, ok := a.ctx.Value(TransactionKey{}).(*sqlx.Tx)
	if !ok {
		panic(ErrNoTransaction)
	}
	return ts
}
