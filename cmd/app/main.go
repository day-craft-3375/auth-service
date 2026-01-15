package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/day-craft-3375/auth-service/internal/adapter/input/grpc"
	"github.com/day-craft-3375/auth-service/internal/adapter/output/appcontext"
	"github.com/day-craft-3375/auth-service/internal/adapter/output/generator/id"
	"github.com/day-craft-3375/auth-service/internal/adapter/output/security/hasher"
	"github.com/day-craft-3375/auth-service/internal/infrastructure/config"
	"github.com/day-craft-3375/auth-service/internal/infrastructure/logger"
	"github.com/day-craft-3375/auth-service/internal/infrastructure/transaction"
	"github.com/day-craft-3375/auth-service/internal/usecase/auth"
	"github.com/day-craft-3375/auth-service/pkg/grpcserver"
	"github.com/day-craft-3375/auth-service/pkg/postgres"
)

func main() {
	// инициализация конфигурации
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// инфраструктура
	log := logger.NewSlog()

	pg, err := postgres.New(
		cfg.Postgres.URL,
		postgres.WithMaxPoolSize(cfg.Postgres.PoolMax),
	)
	if err != nil {
		log.Error("ошибка подключения к Postgres", err)
		os.Exit(1)
	}

	transactionManager := transaction.NewManager(pg.DB())
	defer transactionManager.Close()

	// адаптеры
	appContextManager := appcontext.NewAppContextManager(cfg.Context.Timeout)
	uuidGenerator := id.NewUUID()
	bcryptPasswordHasher := hasher.NewBcrypt(cfg.Security.BcryptCost)

	// бизнес-логика
	authUseCase := auth.NewAuth(
		log,
		appContextManager,
		transactionManager,
		uuidGenerator,
		bcryptPasswordHasher,
	)

	// GRPC сервер
	grpcServer := grpcserver.New(
		grpcserver.WithPort(cfg.GRPC.Port),
	)
	grpc.NewRouter(grpcServer.Server(), authUseCase)

	// обработка сигналов завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		log.Info("получен сигнал завершения", "signal", sig.String())

		grpcServer.Shutdown()
		os.Exit(0)
	}()

	// запуск GRPC сервера
	log.Info("запуск GRPC сервера", "address", grpcServer.Address())
	if err := grpcServer.Run(); err != nil {
		log.Error("ошибка запуска GRPC сервера", err, "address", grpcServer.Address())
		os.Exit(1)
	}
}
