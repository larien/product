package main

import (
	"log"
	"net/http"

	"github.com/larien/product/product/controller"
	"github.com/larien/product/product/drivers/config"
	"github.com/larien/product/product/drivers/database"
	"github.com/larien/product/product/drivers/grpc"
	localLog "github.com/larien/product/product/drivers/log"
	"github.com/larien/product/product/handler"
	"github.com/larien/product/product/repository/discount"
	"github.com/larien/product/product/repository/product"
	protobuf "github.com/larien/product/protos"
)

func main() {
	panic(productService())
}

func productService() error {
	log.Println("Product service has started")

	config := config.New()

	db, err := database.New(config.DB)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	cc, err := grpc.Dial(config.GRPC)
	if err != nil {
		log.Fatal("failed to connect to the gRPC server")
	}
	defer cc.Close()

	discount := discount.New(protobuf.NewDiscountServiceClient(cc))
	product := product.New(db)
	controller := controller.New(product, discount)
	handler := handler.New(controller)

	localLog.SetLevel(config.LogLevel)

	log.Println("Listening in port", config.Port)
	return http.ListenAndServe(config.Port, handler)
}
