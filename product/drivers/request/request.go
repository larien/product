package request

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/larien/product/product/drivers/log"
)

// Write encodes the response for a request with its status code
func Write(ctx context.Context, w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	log.Debugf(ctx, "status code: %d", statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

type SuccessResponse struct {
	Message string `json:"message"`
}

// Success returns a success message for a request
func Success(ctx context.Context, w http.ResponseWriter, statusCode int, message string) {
	Write(ctx, w, statusCode, SuccessResponse{message})
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// Error returns the message with reason for failure for a request
func Error(ctx context.Context, w http.ResponseWriter, statusCode int, err error) {
	Write(ctx, w, statusCode, ErrorResponse{err.Error()})
}
