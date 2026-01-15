package grpcserver

import (
	"fmt"
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
}

// New создает новый GRPC сервер
func New(opts ...Option) *GRPCServer {
	s := &GRPCServer{
		server:  grpc.NewServer(),
		address: _defaultAddr,
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

// Address возвращает адрес GRPC сервера
func (s *GRPCServer) Address() string {
	return s.address
}

// Run запускает GRPC сервер
func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("не удалось создать tcp-слушатель: %w", err)
	}

	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("не удалось запустить GRPC сервер: %w", err)
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
	case <-time.After(10 * time.Second):
		s.server.Stop()
	}
}
