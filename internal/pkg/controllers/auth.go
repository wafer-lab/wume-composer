package controllers

import (
	"log"
	"net/http"

	"wume-composer/internal/pkg/auth"
	"wume-composer/internal/pkg/models"
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

	jwtData, err, incorrectFields := auth.SignIn(signInData)
	if !handleCommonErrors(w, err, incorrectFields) {
		return
	}

	if err := auth.SetAuthCookie(w, jwtData); err != nil {
		log.Println("Impossible to set auth cookie! Error: " + err.Error())
	}
	sendJson(w, http.StatusOK, models.SignedInAnswer)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}
	auth.RemoveAuthCookie(w)
	sendJson(w, http.StatusOK, models.SignedOutAnswer)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}

	updateData := models.UpdatePasswordData{}
	if !parseJson(w, r, &updateData) {
		return
	}

	err, incorrectFields := auth.UpdatePassword(jwtData(r).Id, updateData)
	if !handleCommonErrors(w, err, incorrectFields) {
		return
	}

	sendJson(w, http.StatusOK, models.PasswordUpdatedAnswer)
}
