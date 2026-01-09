package auth

type RegisterUserInput struct {
	Email    string
	Password string
}

func NewRegisterUserInput(email, password string) RegisterUserInput {
	return RegisterUserInput{
		Email:    email,
		Password: password,
	}
}

type RegisterUserOutput struct {
	UserID string
}

func NewRegisterUserOutput(userID string) RegisterUserOutput {
	return RegisterUserOutput{
		UserID: userID,
	}
}
