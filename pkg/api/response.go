package api

import (
	"errors"
	"net/http"

	errutil "github.com/faruqfadhil/currency-api/pkg/error"
	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func internalServerErr(err error) *ResponseError {
	return &ResponseError{
		Status:  "error",
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func badRequestErr(err error) *ResponseError {
	return &ResponseError{
		Status:  "error",
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}

func notFoundErr(err error) *ResponseError {
	return &ResponseError{
		Status:  "error",
		Code:    http.StatusNotFound,
		Message: err.Error(),
	}
}

func ResponseFailed(c *gin.Context, err error) {
	resp := internalServerErr(err)
	typeErr := errutil.GetTypeErr(err)
	if errors.Is(typeErr, errutil.ErrGeneralBadRequest) {
		resp = badRequestErr(err)
	}
	if errors.Is(typeErr, errutil.ErrGeneralNotFound) {
		resp = notFoundErr(err)
	}
	c.JSON(resp.Code, resp)
}

func ResponseSuccess(c *gin.Context, out interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Code:    http.StatusOK,
		Data:    out,
		Message: message,
	})
}
