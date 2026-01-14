package v1

import (
	"context"

	"github.com/day-craft-3375/auth-service/internal/usecase/auth"
	authv1 "github.com/day-craft-3375/protos/gen/go/auth/v1"
	"google.golang.org/grpc"
)

type V1 struct {
	authv1.UnimplementedAuthServiceServer

	authUseCase *auth.AuthUseCase
}

func NewAuthRoutes(
	s *grpc.Server,
	authUseCase *auth.AuthUseCase,
) {
	r := &V1{
		authUseCase: authUseCase,
	}

	authv1.RegisterAuthServiceServer(s, r)
}

func (r *V1) RegisterUser(ctx context.Context, req *authv1.RegisterUserRequest) (*authv1.RegisterUserResponse, error) {
	out, _ := r.authUseCase.RegisterUser(ctx, auth.NewRegisterUserInput(
		req.GetEmail(),
		req.GetPassword(),
	))

	return &authv1.RegisterUserResponse{
		UserId: out.UserID,
	}, nil
}
