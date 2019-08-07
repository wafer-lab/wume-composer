package controllers

import (
	"net/http"

	"wume-composer/internal/pkg/common/logger"
	"wume-composer/internal/pkg/jwt"
	"wume-composer/internal/pkg/models"
	"wume-composer/internal/pkg/user"
)

func IsAuth(w http.ResponseWriter, r *http.Request) {
	if isAuth(r) {
		sendJson(w, http.StatusOK, models.SignedInAnswer)
	} else {
		sendJson(w, http.StatusUnauthorized, models.SignedOutAnswer)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if !requireNotAuth(w, r) {
		return
	}

	signInData := models.SignInData{}
	if !parseJson(w, r, &signInData) {
		return
	}

	jwtData, err, incorrectFields := user.SignIn(signInData)
	if !handleCommonErrors(w, err, incorrectFields) {
		return
	}

	if err := jwt.SetAuthCookie(w, jwtData); err != nil {
		logger.Error("Impossible to set auth cookie! Error: " + err.Error())
	}
	sendJson(w, http.StatusOK, models.SignedInAnswer)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}
	jwt.RemoveAuthCookie(w)
	sendJson(w, http.StatusOK, models.SignedOutAnswer)
}
