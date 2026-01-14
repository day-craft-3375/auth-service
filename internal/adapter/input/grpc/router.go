package grpc

import (
	v1 "github.com/day-craft-3375/auth-service/internal/adapter/input/grpc/v1"
	"github.com/day-craft-3375/auth-service/internal/usecase/auth"
	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewRouter(
	s *pbgrpc.Server,
	authUseCase *auth.AuthUseCase,
) {
	v1.NewAuthRoutes(s, authUseCase)

	reflection.Register(s)
}
