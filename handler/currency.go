package handler

import (
	"github.com/faruqfadhil/currency-api/core/module"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	usecase module.Usecase
}

func New(uc module.Usecase) *HTTPHandler {
	return &HTTPHandler{
		usecase: uc,
	}
}

func (h *HTTPHandler) CreateCurrency(c *gin.Context) {
}
