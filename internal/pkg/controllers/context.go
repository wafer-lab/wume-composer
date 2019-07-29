package controllers

import (
	"net/http"

	"wume-composer/internal/pkg/models"
)

func isAuth(r *http.Request) bool {
	return r.Context().Value("isAuth").(bool)
}

func jwtData(r *http.Request) models.JwtData {
	return r.Context().Value("jwtData").(models.JwtData)
}
