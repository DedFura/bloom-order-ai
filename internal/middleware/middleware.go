package middleware

import (
	"bloom-order-ai/internal/logging"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				var errMsg string
				if e, ok := err.(error); ok {
					errMsg = e.Error()
				} else {
					errMsg = fmt.Sprintf("%v", err)
				}
				log.Println(errMsg)

				jsonBody, jsonErr := json.Marshal(map[string]string{
					"error": errMsg,
				})
				if jsonErr != nil {
					log.Println("Error marshaling JSON:", jsonErr)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(`{"error": "internal server error"}`))
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, writeErr := w.Write(jsonBody)
				if writeErr != nil {
					log.Println("Error writing response:", writeErr)
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func LoggerMiddleware(logger logging.Logger) func(http.Handler) http.Handler {
	return ctxMiddleware(logger, ctxKeyLog)
}

func ctxMiddleware(value interface{}, key ctxKey) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), key, value)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
