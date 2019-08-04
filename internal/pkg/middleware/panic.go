package middleware

import (
	"net/http"
	"runtime/debug"

	"wume-composer/internal/pkg/logger"
)

func PanicCatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("Oh no, it's a panic: %v\n%v", r, string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
