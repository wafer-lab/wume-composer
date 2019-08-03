package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"wume-composer/internal/pkg/models"
)

func isAuth(r *http.Request) bool {
	return r.Context().Value("isAuth").(bool)
}

func jwtData(r *http.Request) models.JwtData {
	return r.Context().Value("jwtData").(models.JwtData)
}

func requireAuth(w http.ResponseWriter, r *http.Request) bool {
	if !isAuth(r) {
		sendJson(w, http.StatusUnauthorized, models.NotSignedInAnswer)
		return false
	}
	return true
}

func requireNotAuth(w http.ResponseWriter, r *http.Request) bool {
	if isAuth(r) {
		sendJson(w, http.StatusForbidden, models.NotSignedOutAnswer)
		return false
	}
	return true
}

func parseJson(w http.ResponseWriter, r *http.Request, result models.InputModel) bool {
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		sendJson(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
		return false
	}

	if err := r.Body.Close(); err != nil {
		sendJson(w, http.StatusInternalServerError, models.IncorrectJsonAnswer)
		return false
	}

	if incorrectFields := result.Validate(); incorrectFields != nil {
		sendJson(w, http.StatusBadRequest, models.GetIncorrectFieldsAnswer(incorrectFields))
		return false
	}

	return true
}

func handleCommonErrors(w http.ResponseWriter, err error, incorrectFields []string) bool {
	if err != nil {
		if incorrectFields != nil {
			if err == models.IncorrectDataError {
				sendJson(w, http.StatusBadRequest, models.GetIncorrectFieldsAnswer(incorrectFields))
			} else {
				sendJson(w, http.StatusBadRequest, &models.IncorrectDataAnswer{
					Status:  200,
					Message: err.Error(),
					Data:    incorrectFields,
				})
			}
			return false
		}
		sendJson(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		log.Println("Developer error: " + err.Error())
		return false
	}
	return true
}

func sendJson(w http.ResponseWriter, statusCode int, outModel models.OutputModel) {
	answer, err := outModel.MarshalJSON()
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(statusCode)
	_, err = fmt.Fprintln(w, string(answer))
	if err != nil {
		log.Println(err)
		return
	}
}
