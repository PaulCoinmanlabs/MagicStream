package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ApiResponse{Code: 0, Message: "ok", Data: data})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, ApiResponse{Code: 0, Message: "created", Data: data})
}

func Fail(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, ApiResponse{Code: httpStatus, Message: message, Data: nil})
}

func BadRequest(c *gin.Context, message string) {
	Fail(c, http.StatusBadRequest, message)
}

func Unauthorized(c *gin.Context, message string) {
	Fail(c, http.StatusUnauthorized, message)
}

func Forbidden(c *gin.Context, message string) {
	Fail(c, http.StatusForbidden, message)
}

func NotFound(c *gin.Context, message string) {
	Fail(c, http.StatusNotFound, message)
}

func InternalError(c *gin.Context, message string) {
	Fail(c, http.StatusInternalServerError, message)
}
