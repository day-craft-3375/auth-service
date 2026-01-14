package hasher

import (
	"github.com/day-craft-3375/auth-service/internal/usecase"
	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct {
	cost int
}

// NewBcrypt создает новый экземпляр bcrypt-хешера с заданной стоимостью.
func NewBcrypt(cost int) usecase.PasswordHasher {
	return &bcryptHasher{cost: cost}
}

// Hash генерирует bcrypt-хеш для заданного пароля.
func (b *bcryptHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare сравнивает заданный пароль с bcrypt-хешем.
func (b *bcryptHasher) Compare(password, hash string) (bool, error) {
	switch err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err {
	case nil:
		return true, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return false, nil
	default:
		return false, err
	}
}
