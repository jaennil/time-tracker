package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func readIDParam(c *gin.Context) (int64, error) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id")
	}

	return id, nil
}
