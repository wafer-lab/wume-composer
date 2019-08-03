package models

import (
	"errors"
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
