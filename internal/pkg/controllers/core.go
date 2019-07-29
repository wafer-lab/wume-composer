package controllers

import (
	"net/http"

	"wume-composer/internal/pkg/models"
)

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	models.Send(w, http.StatusOK, models.GetSuccessAnswer("Backend of project 'wume-composer'!"))
}

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	if isAuth(r) {
		models.Send(w, http.StatusOK, models.GetSuccessAnswer("Hello, "+jwtData(r).Username+"!"))
	} else {
		models.Send(w, http.StatusOK, models.GetSuccessAnswer("I don't know about you, but hello!"))
	}
}
