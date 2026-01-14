package grpcserver

import (
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	_defaultAddr = ":50051"
)

// GRPCServer GRPC сервер
type GRPCServer struct {
	server *grpc.Server

	address string
	log     Logger
}

// New создает новый GRPC сервер
func New(log Logger, opts ...Option) *GRPCServer {
	s := &GRPCServer{
		server:  grpc.NewServer(),
		address: _defaultAddr,
		log:     log,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Server возвращает grpc.Server
func (s *GRPCServer) Server() *grpc.Server {
	return s.server
}

// Run запускает GRPC сервер
func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		s.log.Error("не удалось запустить GRPC сервер", err, "address", s.address)
		return err
	}

	s.log.Info("GRPC сервер запущен", "address", s.address)
	if err := s.server.Serve(lis); err != nil {
		s.log.Error("ошибка GRPC сервера", err, "address", s.address)
		return err
	}
	return nil
}

// Shutdown останавливает GRPC сервер
func (s *GRPCServer) Shutdown() {
	done := make(chan struct{})

	go func() {
		s.server.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		s.log.Info("GRPC сервер остановлен", "address", s.address)
	case <-time.After(10 * time.Second):
		s.server.Stop()
		s.log.Info("GRPC сервер принудительно остановлен", "address", s.address)
	}
}
