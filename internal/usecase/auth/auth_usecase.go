package auth

import (
	"context"

	"github.com/day-craft-3375/auth-service/internal/domain"
	"github.com/day-craft-3375/auth-service/internal/usecase"
)

// AuthUseCase представляет собой реализацию бизнес-логики аутентификации
type AuthUseCase struct {
	appContextManager usecase.AppContextManager
	idGenerator       usecase.IDGenerator
	passwordHasher    usecase.PasswordHasher
}

// NewAuth создает новый экземпляр AuthUseCase
func NewAuth(
	appContextManager usecase.AppContextManager,
	idGenerator usecase.IDGenerator,
	passwordHasher usecase.PasswordHasher,
) *AuthUseCase {
	return &AuthUseCase{
		appContextManager: appContextManager,
		idGenerator:       idGenerator,
		passwordHasher:    passwordHasher,
	}
}

// RegisterUser регистрирует нового пользователя
func (u *AuthUseCase) RegisterUser(ctx context.Context, in RegisterUserInput) (RegisterUserOutput, error) {
	_, cancel := u.appContextManager.CreateContext(ctx)
	defer cancel()

	// хеширует пароль пользователя
	hash, err := u.passwordHasher.Hash(in.Password)
	if err != nil {
		return RegisterUserOutput{}, err
	}

	// создаёт нового пользователя
	user := domain.NewUser(
		u.idGenerator.NewID(),
		in.Email,
		hash,
	)

	return NewRegisterUserOutput(
		user.ID,
	), nil
}
