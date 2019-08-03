package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"wume-composer/internal/pkg/auth"
	"wume-composer/internal/pkg/models"
)

func IsAuth(w http.ResponseWriter, r *http.Request) {
	if isAuth(r) {
		models.Send(w, http.StatusOK, models.SignedInAnswer)
	} else {
		models.Send(w, http.StatusUnauthorized, models.SignedOutAnswer)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if isAuth(r) {
		models.Send(w, http.StatusMethodNotAllowed, models.AlreadySignedInAnswer)
		return
	}

	signInData := models.SignInData{}
	err := json.NewDecoder(r.Body).Decode(&signInData)
	if err != nil {
		models.Send(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
		return
	}
	defer r.Body.Close()

	jwtData, err, fields := auth.SignIn(signInData)
	if err != nil {
		if fields != nil {
			if err == models.IncorrectDataError {
				models.Send(w, http.StatusBadRequest, models.GetIncorrectFieldsAnswer(fields))
			} else {
				models.Send(w, http.StatusBadRequest, &models.IncorrectFieldsAnswer{
					Status:  200,
					Message: err.Error(),
					Data:    fields,
				})
			}
			return
		}
		models.Send(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		fmt.Println(fmt.Sprintf("DEV ERR: %q ==> %v", r.RequestURI, err))
		return
	}

	http.SetCookie(w, auth.CreateAuthCookie(jwtData, 30*24*time.Hour))
	models.Send(w, http.StatusOK, models.SignedInAnswer)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	if !isAuth(r) {
		models.Send(w, http.StatusMethodNotAllowed, models.AlreadySignedOutAnswer)
		return
	}

	jwtCookie := auth.CreateAuthCookie(models.JwtData{}, -1)
	http.SetCookie(w, jwtCookie)
	models.Send(w, http.StatusOK, models.SignedOutAnswer)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	if !isAuth(r) {
		models.Send(w, http.StatusUnauthorized, models.NotSignedInAnswer)
		return
	}

	updateData := models.UpdatePasswordData{}
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		models.Send(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
		return
	}
	defer r.Body.Close()

	err, fields := auth.UpdatePassword(jwtData(r).Id, updateData)
	if err != nil {
		if fields != nil {
			if err == models.IncorrectDataError {
				models.Send(w, http.StatusBadRequest, models.GetIncorrectFieldsAnswer(fields))
			} else {
				models.Send(w, http.StatusBadRequest, &models.IncorrectFieldsAnswer{
					Status:  200,
					Message: err.Error(),
					Data:    fields,
				})
			}
			return
		}
		models.Send(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		fmt.Println(fmt.Sprintf("DEV ERR: %q ==> %v", r.RequestURI, err))
		return
	}

	models.Send(w, http.StatusOK, models.PasswordUpdatedAnswer)
}
