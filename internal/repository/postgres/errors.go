package postgres

import (
	"errors"
)

var RecordNotFound error = errors.New("record not found")
var InternalServerError error = errors.New("the server encountered a problem and could not process your request")
