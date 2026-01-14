package logger

import (
	"log/slog"
	"os"

	"github.com/day-craft-3375/auth-service/internal/usecase"
	"github.com/day-craft-3375/auth-service/pkg/slogconsole"
)

type slogLogger struct {
	log *slog.Logger
}

func NewSlog() usecase.Logger {
	// TODO: настроить уровень логирования из конфигурации

	handler := slogconsole.New(
		os.Stderr,
		slogconsole.WithLevel(slog.LevelDebug),
		slogconsole.WithSource(true),
	)

	log := slog.New(handler)

	return &slogLogger{
		log: log,
	}
}

func (sl slogLogger) Debug(msg string, args ...any) {
	sl.log.Debug(msg, args...)
}

func (sl slogLogger) Info(msg string, args ...any) {
	sl.log.Info(msg, args...)
}

func (sl slogLogger) Warn(msg string, args ...any) {
	sl.log.Warn(msg, args...)
}

func (sl slogLogger) Error(msg string, err error, args ...any) {
	args = append(args, "err", err.Error())
	sl.log.Error(msg, args...)
}
