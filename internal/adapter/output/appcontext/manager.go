package appcontext

import (
	"context"
	"time"

	"github.com/day-craft-3375/auth-service/internal/usecase"
)

type appContextManager struct {
	timeout time.Duration
}

// NewAppContextManager создает менеджер контекстов приложения
func NewAppContextManager(timeout time.Duration) usecase.AppContextManager {
	return &appContextManager{
		timeout: timeout,
	}
}

// CreateContext создает контекст приложения
func (m *appContextManager) CreateContext(parent context.Context) (usecase.AppContext, usecase.AppContextCancelFunc) {
	ctx, cancel := newAppContext(parent, m.timeout)
	return ctx, cancel
}
