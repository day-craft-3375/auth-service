package auth

import (
	"context"

	"github.com/day-craft-3375/auth-service/internal/domain"
	"github.com/day-craft-3375/auth-service/internal/usecase"
)

// AuthUseCase представляет собой реализацию бизнес-логики аутентификации
type AuthUseCase struct {
	log               usecase.Logger
	appContextManager usecase.AppContextManager
	tsManager         usecase.TransactionManager
	idGenerator       usecase.IDGenerator
	passwordHasher    usecase.PasswordHasher
}

// NewAuth создает новый экземпляр AuthUseCase
func NewAuth(
	log usecase.Logger,
	appContextManager usecase.AppContextManager,
	tsManager usecase.TransactionManager,
	idGenerator usecase.IDGenerator,
	passwordHasher usecase.PasswordHasher,
) *AuthUseCase {
	return &AuthUseCase{
		appContextManager: appContextManager,
		tsManager:         tsManager,
		idGenerator:       idGenerator,
		passwordHasher:    passwordHasher,
	}
}

// RegisterUser регистрирует нового пользователя
func (u *AuthUseCase) RegisterUser(ctx context.Context, in RegisterUserInput) (RegisterUserOutput, error) {
	var zero RegisterUserOutput

	// открывает транзакцию
	ts := u.tsManager.CreateTransaction()
	if err := ts.Start(); err != nil {
		u.log.Error("не удалось открыть транзакцию", err)
		return zero, usecase.ErrInternalError
	}
	defer ts.Rollback()

	// создаёт контекст приложения
	_, cancel := u.appContextManager.CreateContext(ctx)
	defer cancel()

	// хеширует пароль пользователя
	hash, err := u.passwordHasher.Hash(in.Password)
	if err != nil {
		u.log.Error("не удалось хешировать пароль", err)
		return zero, usecase.ErrInternalError
	}

	// создаёт нового пользователя
	user := domain.NewUser(
		u.idGenerator.NewID(),
		in.Email,
		hash,
	)

	// фиксирует транзакцию
	if err := ts.Commit(); err != nil {
		u.log.Error("не удалось зафиксировать транзакцию", err)
		return zero, usecase.ErrInternalError
	}

	return NewRegisterUserOutput(
		user.ID,
	), nil
}
