package handler

import (
	"context"

	"github.com/faruqfadhil/currency-api/core/entity"
	"github.com/faruqfadhil/currency-api/core/module"
	"github.com/faruqfadhil/currency-api/pkg/api"
	errutil "github.com/faruqfadhil/currency-api/pkg/error"
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
	var payload *entity.CreateCurrencyRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		api.ResponseFailed(c, errutil.ErrGeneralBadRequest)
		return
	}
	err := h.usecase.CreateCurrency(context.Background(), payload)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	api.ResponseSuccess(c, "", "success")
}
