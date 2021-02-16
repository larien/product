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
	pkgErrors "github.com/pkg/errors"
)

var (
	errInvalidUserID = errors.New("invalid user ID")
)

type id string

const (
	// RequestIDKey is the key that holds the unique request ID in a request context
	RequestIDKey id = "requestID"

	// UserIDKey is the key that holds the unique user ID in a request context
	UserIDKey id = "userID"
)

// Logger is a middleware to define the logging configuration to be used in the entire stack
// It also creates a request ID so that it can be injected into logging
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestID := r.Header.Get("X-REQUEST-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx = context.WithValue(ctx, RequestIDKey, requestID)

		ctx = log.New(ctx, requestID)
		t1 := time.Now()

		defer func() {
			log.Infof(ctx, "%s %s%s %s", r.Method, r.Host, r.URL, time.Since(t1).String())
		}()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// UserID is a middleware to obtain and validate the productID received in the request
// The user ID is only injected in the stack if it's valid, but it's not mandatory
func UserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userID := r.Header.Get("X-USER-ID")
		if userID != "" {
			if isValid(userID) {
				ctx = log.Insert(ctx, "user_id", userID)
				ctx = context.WithValue(ctx, UserIDKey, userID)
			} else {
				request.Error(ctx, w, http.StatusBadRequest, pkgErrors.Wrap(errInvalidUserID, userID))
				return
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TODO: regular expressions usually are slow and it'd be better to find another way to validate UUIDs in the future
func isValid(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
