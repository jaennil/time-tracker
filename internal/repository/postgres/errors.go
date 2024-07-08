package postgres

import (
	"errors"
)

var RecordNotFound error = errors.New("record not found")
