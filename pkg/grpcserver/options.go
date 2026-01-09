package grpcserver

import "fmt"

type Option func(*GRPCServer)

func Port(port string) Option {
	return func(s *GRPCServer) {
		s.address = fmt.Sprintf(":%s", port)
	}
}
