package grpc

import (
	"net"

	"google.golang.org/grpc"
)

type ServerConfig struct {
	Listener   net.Listener
	GRPCServer *grpc.Server
}
