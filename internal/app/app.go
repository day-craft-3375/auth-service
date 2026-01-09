package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/day-craft-3375/auth-service/config"
	"github.com/day-craft-3375/auth-service/internal/adapter/input/grpc"
	auth_usecase "github.com/day-craft-3375/auth-service/internal/usecase/auth"
	"github.com/day-craft-3375/auth-service/pkg/grpcserver"
)

func Run(cfg *config.Config) {
	// Use Cases
	registerUserUseCase := auth_usecase.NewRegisterUser()

	// GRPC server
	grpcServer := grpcserver.New(grpcserver.Port(cfg.GRPC.Port))
	grpc.NewRouter(grpcServer.Server, registerUserUseCase)

	// запускает серверы
	grpcServer.Run()

	// ожидает сигнала
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("app - signal: %s\n", s.String())
	case err := <-grpcServer.Notify():
		log.Println("App -", err)
	}

	// shutdown
	if err := grpcServer.Shutdown(); err != nil {
		log.Println("App -", err)
	}
}
