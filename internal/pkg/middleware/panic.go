package middleware

import (
	"fmt"
	"net/http"
)

func PanicCatcher(h http.Handler) http.Handler {
	var mw http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Oh no, it's a panic: %v", r)
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	}

	return mw
}
