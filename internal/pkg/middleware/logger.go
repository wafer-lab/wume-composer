package middleware

import (
	"net/http"

	"wume-composer/internal/pkg/logger"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("%s (%s)", r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}

