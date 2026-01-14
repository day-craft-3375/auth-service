package grpcserver

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, err error, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
}
