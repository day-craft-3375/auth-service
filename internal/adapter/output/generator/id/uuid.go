package id

import (
	"github.com/day-craft-3375/auth-service/internal/usecase"
	"github.com/google/uuid"
)

type uuidGenerator struct {
}

func NewUUID() usecase.IDGenerator {
	return &uuidGenerator{}
}

// NewID генерирует UUID в виде строки
func (g *uuidGenerator) NewID() string {
	return uuid.NewString()
}
