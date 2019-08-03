package middleware

import (
	"context"
	"net/http"

	"wume-composer/internal/pkg/auth"
	"wume-composer/internal/pkg/config"
	"wume-composer/internal/pkg/models"
)

func AuthChecker(h http.Handler) http.Handler {
	var mw http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		jwtCookie, errNoCookie := req.Cookie(config.Auth.CookieName)
		if errNoCookie != nil {
			ctx = context.WithValue(ctx, "isAuth", false)
			ctx = context.WithValue(ctx, "jwtData", models.JwtData{})
		} else {
			data, err := auth.CheckJwt(jwtCookie.Value)
			if err != nil {
				ctx = context.WithValue(ctx, "isAuth", false)
				ctx = context.WithValue(ctx, "jwtData", models.JwtData{})
			} else {
				ctx = context.WithValue(ctx, "isAuth", true)
				ctx = context.WithValue(ctx, "jwtData", data)
			}
		}
		h.ServeHTTP(res, req.WithContext(ctx))
	}

	return mw
}
