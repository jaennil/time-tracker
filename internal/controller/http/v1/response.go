package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Error string `json:"error" example:"message"`
}

type InternalServerErrorResponse struct {
	Error string `json:"error" example:"the server encountered a problem and could not process your request"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, Response{msg})
}

func NotFoundResponse(c *gin.Context) {
	errorResponse(c, http.StatusNotFound, "route not found")
}

func noContentResponse(c *gin.Context) {
	c.AbortWithStatus(http.StatusNoContent)
}

func internalServerErrorResponse(c *gin.Context) {
	errorResponse(c, http.StatusInternalServerError, "the server encountered a problem and could not process your request")
}
