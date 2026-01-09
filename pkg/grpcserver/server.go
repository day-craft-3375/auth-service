package grpcserver

import (
	"context"
	"errors"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

const (
	_defaultAddr = ":50051"
)

type GRPCServer struct {
	Server *grpc.Server

	ctx context.Context
	eg  *errgroup.Group

	notify  chan error
	address string
}

func New(opts ...Option) *GRPCServer {
	group, ctx := errgroup.WithContext(context.Background())
	group.SetLimit(1) // запуск только в одной горитине

	s := &GRPCServer{
		ctx:     ctx,
		eg:      group,
		Server:  grpc.NewServer(),
		notify:  make(chan error),
		address: _defaultAddr,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *GRPCServer) Run() {
	s.eg.Go(func() error {
		var lc net.ListenConfig

		ln, err := lc.Listen(s.ctx, "tcp", s.address)
		if err != nil {
			s.notify <- err

			close(s.notify)

			return err
		}

		if err := s.Server.Serve(ln); err != nil {
			s.notify <- err

			close(s.notify)

			return err
		}

		return nil
	})

	log.Println("grpc server - Started")
}

func (s *GRPCServer) Notify() <-chan error {
	return s.notify
}

func (s *GRPCServer) Shutdown() error {
	var errs []error

	s.Server.GracefulStop()

	// дожидается завершения всех горрутин
	err := s.eg.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Println("grpc server:", err)

		errs = append(errs, err)
	}

	log.Println("grpc server - Shutdown")
	return errors.Join(errs...)
}
