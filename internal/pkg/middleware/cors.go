package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		origin, hasOrigin := req.Header["Origin"]
		if hasOrigin {
			found := false
			for _, allowed := range corsData.AllowOrigins {
				if origin[0] == allowed {
					found = true
					break
				}
			}
			if found {
				res.Header().Set("Access-Control-Allow-Origin", origin[0])
				res.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(corsData.AllowCredentials))
			} else {
				fmt.Println("Origin " + origin[0] + " wasn't found!")
			}
		}

		if req.Method == "OPTIONS" {
			res.Header().Set("Access-Control-Allow-Methods", strings.Join(corsData.AllowMethods, ", "))
			res.Header().Set("Access-Control-Allow-Headers", strings.Join(corsData.AllowHeaders, ", "))
			res.Header().Set("Access-Control-Max-Age", strconv.Itoa(corsData.MaxAge))
			return
		}

		next.ServeHTTP(res, req)
	})
}
