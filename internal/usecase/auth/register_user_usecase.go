package auth

import "context"

type RegisterUserUseCase struct {
}

func NewRegisterUser() *RegisterUserUseCase {
	return &RegisterUserUseCase{}
}

func (u *RegisterUserUseCase) Execute(ctx context.Context, in RegisterUserInput) (RegisterUserOutput, error) {
	return NewRegisterUserOutput(
		"test-user-id-from-usecase",
	), nil
}
