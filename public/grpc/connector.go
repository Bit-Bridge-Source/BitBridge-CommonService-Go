package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IGrpcConnector interface {
	Connect(serviceURL string) (*grpc.ClientConn, error)
}

type GrpcConnector struct{}

// Connect implements IGrpcConnector.
func (*GrpcConnector) Connect(serviceURL string) (*grpc.ClientConn, error) {
	return grpc.Dial(serviceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// Ensure that GrpcConnector implements IGrpcConnector
var _ IGrpcConnector = (*GrpcConnector)(nil)
