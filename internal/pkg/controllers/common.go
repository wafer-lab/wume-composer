package controllers

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"wume-composer/internal/pkg/common/logger"
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

func getStrUrlParam(r *http.Request, name string) (string, error) {
	vars := mux.Vars(r)
	result, ok := vars[name]
	if ok {
		return result, nil
	} else {
		return "", models.IncorrectDataError
	}
}

func getIntUrlParam(r *http.Request, name string) (int64, error) {
	str, err := getStrUrlParam(r, name)
	if err != nil {
		return 0, err
	}
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, models.IncorrectDataError
	}
	return result, nil
}

func getStrQueryParam(r *http.Request, name string) (value string, err error) {
	value = r.URL.Query().Get(name)
	if value == "" {
		err = models.IncorrectDataError
	}
	return
}

func getIntQueryParam(r *http.Request, name string) (int64, error) {
	str, err := getStrQueryParam(r, name)
	if err != nil {
		return 0, err
	}
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, models.IncorrectDataError
	}
	return result, nil
}

func getFormFile(w http.ResponseWriter, r *http.Request, maxSize int64, name string) (multipart.File, *multipart.FileHeader, error) {
	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		sendJson(w, http.StatusBadRequest, models.FormTooBigAnswer)
		return nil, nil, err
	}

	file, header, err := r.FormFile(name)
	if err != nil {
		sendJson(w, http.StatusBadRequest, models.GetIncorrectFieldsAnswer([]string{name}))
		return nil, nil, err
	}

	return file, header, nil
}

func parseJson(w http.ResponseWriter, r *http.Request, result models.InputModel) bool {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendJson(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		logger.Error("Developer error: " + err.Error())
		return false
	}

	if err = r.Body.Close(); err != nil {
		sendJson(w, http.StatusInternalServerError, models.GetDeveloperErrorAnswer(err.Error()))
		logger.Error("Developer error: " + err.Error())
		return false
	}

	if err = result.UnmarshalJSON(data); err != nil {
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
		logger.Error("Developer error: " + err.Error())
		return false
	}
	return true
}

func sendJson(w http.ResponseWriter, statusCode int, outModel models.OutputModel) {
	answer, err := outModel.MarshalJSON()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	w.WriteHeader(statusCode)
	_, err = fmt.Fprintln(w, string(answer))
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
