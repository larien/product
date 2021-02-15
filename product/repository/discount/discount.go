package discount

import (
	"context"

	"github.com/larien/product/product/handler/middleware"
	protobuf "github.com/larien/product/protos"

	grpc "google.golang.org/grpc"
)

// Discount represents the methods available in this repository layer
type Discount interface {
	Get(ctx context.Context, productID, userID string) (int64, error)
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

// Get obtains a product by its ID from the database
func (r *repository) Get(ctx context.Context, productID, userID string) (int64, error) {
	if userID == "" {
		return 0, nil
	}
	// TODO - is there a way to keep the ctx across services written in different technologies?
	requestID, _ := ctx.Value(middleware.RequestIDKey).(string)
	req := &protobuf.DiscountRequest{
		RequestID: requestID,
		ProductID: productID,
		UserID:    userID,
	}
	res, err := r.Client.Discount(ctx, req)
	if err != nil {
		return 0, err
	}
	return res.Percentage, nil
}
