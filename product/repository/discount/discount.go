package discount

import (
	"context"

	protobuf "github.com/larien/product-service/protos"

	grpc "google.golang.org/grpc"
)

// Discount represents the methods available in this repository layer
type Discount interface {
}

// Client represents the methods used by this repository layer
type Client interface {
	// Discount obtains the discount percentage via gRPC communication
	Discount(
		ctx context.Context,
		in *protobuf.DiscountRequest,
		opts ...grpc.CallOption) (*protobuf.DiscountResponse, error)
}

type repository struct {
	Client Client
}

// New creates a new instance of Product repository to manipulate the database
func New(client Client) Discount {
	return &repository{client}
}

