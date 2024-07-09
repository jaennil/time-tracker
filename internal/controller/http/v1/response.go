package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}

func NotFoundResponse(c *gin.Context) {
	errorResponse(c, http.StatusNotFound, "route not found")
}

func noContentResponse(c *gin.Context) {
	c.AbortWithStatus(http.StatusNoContent)
}
