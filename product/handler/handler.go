package handler

import (
	"context"
	"net/http"

	"github.com/larien/product/product/drivers/request"
	"github.com/larien/product/product/drivers/router"
	"github.com/larien/product/product/entity"
	"github.com/larien/product/product/handler/middleware"
)

// New creates a new instance of Product handler with a router to make endpoints available
func New(c Product) router.Router {
	r := router.New()
	r.Use(middleware.Logger)
	r.Get("/status", healthcheck()) // GET /status
	r.Route("/v1", func(r router.Router) {
		r.Route("/product", func(r router.Router) {
			r.Use(middleware.UserID)
			r.Get("/", list(c)) // GET /product
		})
	})
	return r
}

// Product represents the methods used by handlers in this layer
type Product interface {
	List(ctx context.Context, userID string) (products []*entity.Product, err error)
}

// list is the handler for Product's list and handles GET /product
func list(controller Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userID, _ := ctx.Value(middleware.UserIDKey).(string)
		products, err := controller.List(ctx, userID)
		if err != nil {
			request.Error(ctx, w, http.StatusInternalServerError, err)
			return
		}

		if products == nil {
			request.Write(ctx, w, http.StatusNoContent, nil)
			return
		}

		request.Write(ctx, w, http.StatusOK, products)
	}
}

// healthcheck is the handler to be used by instrumentation to make sure the system is up and running
func healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request.Write(r.Context(), w, http.StatusOK, nil)
	}
}
