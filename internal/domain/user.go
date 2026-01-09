package domain

type User struct {
	ID             string
	Email          string
	HashedPassword string
}

func NewUser(id, email, hashedPassword string) *User {
	return &User{
		ID:             id,
		Email:          email,
		HashedPassword: hashedPassword,
	}
}
