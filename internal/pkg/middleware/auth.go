package middleware

import (
	"context"
	"net/http"

	"wume-composer/internal/pkg/jwt"
	"wume-composer/internal/pkg/config"
	"wume-composer/internal/pkg/models"
)

func AuthChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		jwtCookie, errNoCookie := r.Cookie(config.Auth.CookieName)
		if errNoCookie != nil {
			ctx = context.WithValue(ctx, "isAuth", false)
			ctx = context.WithValue(ctx, "jwtData", models.JwtData{})
		} else {
			data, err := jwt.CheckJwt(jwtCookie.Value)
			if err != nil {
				ctx = context.WithValue(ctx, "isAuth", false)
				ctx = context.WithValue(ctx, "jwtData", models.JwtData{})
			} else {
				ctx = context.WithValue(ctx, "isAuth", true)
				ctx = context.WithValue(ctx, "jwtData", data)
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
