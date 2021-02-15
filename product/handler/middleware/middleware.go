package middleware

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"time"

	"github.com/google/uuid"

	"github.com/larien/product/product/drivers/log"
	"github.com/larien/product/product/drivers/request"
	"github.com/larien/product/product/drivers/router"
	pkgErrors "github.com/pkg/errors"
)

var (
	errInvalidProductID = errors.New("invalid product ID")
)

type id string

const (
	// RequestIDKey is the key that holds the unique request ID in a request context
	RequestIDKey id = "requestID"

	// ProductIDKey is the key that holds the unique product ID in a request context
	ProductIDKey id = "productID"

	// UserIDKey is the key that holds the unique user ID in a request context
	UserIDKey id = "userID"
)

// RequestID is a middleware to generate an UUID that will be passed to the entire stack
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestID := r.Header.Get("X-REQUEST-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx = context.WithValue(ctx, RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logger is a middleware to define the logging configuration to be used in the entire stack
// It must be used after RequestID so that request ID is injected into the logging.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestID, _ := ctx.Value(RequestIDKey).(string)
		ctx = log.New(ctx, requestID)
		t1 := time.Now()

		defer func() {
			log.Debugf(ctx, "%s %s%s %s", r.Method, r.Host, r.URL, time.Since(t1).String())
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ProductID is a middleware to obtain and validate the productID received in the request
// This middleware must only be used in endpoints that need the product ID field and any
// invalid ID (non-uuid or empty) will return a Bad Request status
func ProductID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		productID := router.GetParamFromURL(r, "productID")
		if !isValid(productID) {
			request.Error(ctx, w, http.StatusBadRequest, pkgErrors.Wrap(errInvalidProductID, productID))
			return
		}

		ctx = log.Insert(ctx, "product_id", productID)
		ctx = context.WithValue(ctx, ProductIDKey, productID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserID is a middleware to obtain and validate the productID received in the request
// The user ID is only injected in the stack if it's valid, but it's not mandatory
func UserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := r.Header.Get("X-USER-ID")
		if isValid(userID) {
			ctx = log.Insert(ctx, "user_id", userID)
			ctx = context.WithValue(ctx, UserIDKey, userID)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TODO: find a better way to validate UUIDs
func isValid(uuid string) bool {
	if uuid == "" {
		return false
	}
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
