package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"wume-composer/internal/pkg/common/logger"
)

type CorsData struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	MaxAge           int
	AllowCredentials bool
}

var corsData = CorsData{
	AllowOrigins: []string{
		"http://localhost:6001",
	},
	AllowMethods: []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
	},
	AllowHeaders: []string{
		"Content-Type",
	},
	MaxAge:           88500,
	AllowCredentials: true,
}

func ApplyCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin, hasOrigin := r.Header["Origin"]
		if hasOrigin {
			found := false
			for _, allowed := range corsData.AllowOrigins {
				if origin[0] == allowed {
					found = true
					break
				}
			}
			if found {
				w.Header().Set("Access-Control-Allow-Origin", origin[0])
				w.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(corsData.AllowCredentials))
			} else {
				logger.Warn("Origin " + origin[0] + " wasn't found!")
			}
		}

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(corsData.AllowMethods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(corsData.AllowHeaders, ", "))
			w.Header().Set("Access-Control-Max-Age", strconv.Itoa(corsData.MaxAge))
			return
		}

		next.ServeHTTP(w, r)
	})
}
