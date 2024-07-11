package logger

// Middleware for using slog logger in http server

import (
	"log/slog"
	"net/http"
)

// Logger middleware

func Logger(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info("Request", slog.String("method", r.Method), slog.String("url", r.URL.Path))
			next.ServeHTTP(w, r)
		})
	}
}
