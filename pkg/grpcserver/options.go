package grpcserver

import "fmt"

type Option func(*GRPCServer)

func WithPort(port string) Option {
	return func(s *GRPCServer) {
		s.address = fmt.Sprintf(":%s", port)
	}
}
