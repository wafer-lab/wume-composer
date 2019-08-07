package controllers

import (
	"net/http"

	"wume-composer/internal/pkg/common/logger"
	"wume-composer/internal/pkg/jwt"
	"wume-composer/internal/pkg/models"
	"wume-composer/internal/pkg/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if !requireNotAuth(w, r) {
		return
	}

	signUpData := models.SignUpData{}
	if !parseJson(w, r, &signUpData) {
		return
	}

	jwtData, err, incorrectFields := user.CreateUser(signUpData)
	if err == models.AlreadyExistsError {
		sendJson(w, http.StatusConflict, models.GetUserExistsAnswer(incorrectFields))
		return
	} else if !handleCommonErrors(w, err, incorrectFields) {
		return
	}

	if err := jwt.SetAuthCookie(w, jwtData); err != nil {
		logger.Error("Impossible to set auth cookie! Error: " + err.Error())
	}
	sendJson(w, http.StatusOK, models.UserCreatedAnswer)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}

	username, err := getStrQueryParam(r, "username")
	if err != nil {
		username = jwtData(r).Username
	}

	userData, err := user.GetUser(username)
	if err == models.NotFoundError {
		sendJson(w, http.StatusNotFound, models.UserNotFoundAnswer)
		return
	} else if !handleCommonErrors(w, err, nil) {
		return
	}

	sendJson(w, http.StatusOK, models.GetUserDataAnswer(userData))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}

	updateData := models.UpdateUserData{}
	if !parseJson(w, r, &updateData) {
		return
	}

	jwtData, err, incorrectFields := user.UpdateUser(jwtData(r).Id, updateData)
	if err == models.AlreadyExistsError {
		sendJson(w, http.StatusConflict, models.GetUserExistsAnswer(incorrectFields))
		return
	} else if !handleCommonErrors(w, err, incorrectFields) {
		return
	}

	if err := jwt.SetAuthCookie(w, jwtData); err != nil {
		logger.Error("Impossible to set auth cookie! Error: " + err.Error())
	}
	sendJson(w, http.StatusOK, models.UserUpdatedAnswer)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}

	updateData := models.UpdatePasswordData{}
	if !parseJson(w, r, &updateData) {
		return
	}

	err, incorrectFields := user.UpdatePassword(jwtData(r).Id, updateData)
	if !handleCommonErrors(w, err, incorrectFields) {
		return
	}

	sendJson(w, http.StatusOK, models.PasswordUpdatedAnswer)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}

	var removeData models.RemoveUserData
	if !parseJson(w, r, &removeData) {
		return
	}

	err, incorrectFields := user.RemoveUser(jwtData(r).Id, removeData)
	if !handleCommonErrors(w, err, incorrectFields) {
		return
	}

	jwt.RemoveAuthCookie(w)
	sendJson(w, http.StatusOK, models.UserRemovedAnswer)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if !requireAuth(w, r) {
		return
	}

	page, err := getIntQueryParam(r, "page")
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := getIntQueryParam(r, "limit")
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	usersData, err := user.GetUsers(uint(page), uint(limit))
	if !handleCommonErrors(w, err, nil) {
		return
	}

	sendJson(w, http.StatusOK, models.GetUsersDataAnswer(usersData))
}
