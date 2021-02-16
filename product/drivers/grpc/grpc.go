package grpc

import (
	"github.com/larien/product/product/drivers/config"
	"google.golang.org/grpc"
)

// Connection is an alias for the gRPC connection that represents the connection from the gRPC client to the server
type Connection = grpc.ClientConn

// Dial is a wrapper for the Dial's gRPC function with a standard transport security option
func Dial(config *config.GRPC) (*Connection, error) {
	return grpc.Dial(config.Host+config.Port, grpc.WithInsecure()) // should be secure in a real application
}
