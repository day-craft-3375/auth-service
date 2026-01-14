package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/day-craft-3375/auth-service/config"
	"github.com/day-craft-3375/auth-service/internal/adapter/input/grpc"
	"github.com/day-craft-3375/auth-service/internal/adapter/output/appcontext"
	"github.com/day-craft-3375/auth-service/internal/adapter/output/generator/id"
	"github.com/day-craft-3375/auth-service/internal/adapter/output/logger"
	"github.com/day-craft-3375/auth-service/internal/adapter/output/security/hasher"
	"github.com/day-craft-3375/auth-service/internal/usecase/auth"
	"github.com/day-craft-3375/auth-service/pkg/grpcserver"
	"github.com/day-craft-3375/auth-service/pkg/postgres"
)

func Run(cfg *config.Config) {
	// Adapters
	log := logger.NewSlog()
	appContextManager := appcontext.NewAppContextManager(cfg.Context.Timeout)
	uuidGenerator := id.NewUUID()
	bcryptPasswordHasher := hasher.NewBcrypt(cfg.Security.BcryptCost)

	// инфраструктура
	pg, err := postgres.New(
		cfg.Postgres.URL,
		postgres.WithMaxPoolSize(cfg.Postgres.PoolMax),
	)
	if err != nil {
		log.Error("ошибка подключения к Postgres", err)
		os.Exit(1)
	}
	defer pg.Close()

	// Use Cases
	authUseCase := auth.NewAuth(
		appContextManager,
		uuidGenerator,
		bcryptPasswordHasher,
	)

	// GRPC server
	grpcServer := grpcserver.New(
		log,
		grpcserver.WithPort(cfg.GRPC.Port),
	)
	grpc.NewRouter(grpcServer.Server(), authUseCase)

	// Обработка сигналов завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.Info("получен сигнал завершения", "signal", sig.String())

		grpcServer.Shutdown()
		pg.Close()
		os.Exit(0)
	}()

	// запуск GRPC сервера
	if err := grpcServer.Run(); err != nil {
		os.Exit(1)
	}
}
