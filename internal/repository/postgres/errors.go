package postgres

import (
	"errors"
)

var RecordNotFound error = errors.New("record not found")
var InvalidPassportFormat error = errors.New("invalid passport format")
var InternalServerError error = errors.New("the server encountered a problem and could not process your request")
