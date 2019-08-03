package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	NotFoundError      = errors.New("not found")
	AlreadyExistsError = errors.New("already exists")
	IncorrectDataError = errors.New("incorrect data")
)

type InputModel interface {
	Validate() []string
}

type OutputModel interface {
	MarshalJSON() ([]byte, error)
}

func Send(w http.ResponseWriter, statusCode int, outModel interface{}) {
	w.WriteHeader(statusCode)
	jsonModel, _ := json.Marshal(outModel)
	_, _ = fmt.Fprintln(w, string(jsonModel))
}
