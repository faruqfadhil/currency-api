package handler

import (
	"context"
	"strconv"

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

func (h *HTTPHandler) CreateConversionRate(c *gin.Context) {
	var payload *entity.CreateCurrencyConversionRate
	if err := c.ShouldBindJSON(&payload); err != nil {
		api.ResponseFailed(c, errutil.ErrGeneralBadRequest)
		return
	}
	err := h.usecase.CreateConversionRate(context.Background(), payload)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	api.ResponseSuccess(c, "", "success")
}

func (h *HTTPHandler) Convert(c *gin.Context) {
	var payload *entity.ConvertRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		api.ResponseFailed(c, errutil.ErrGeneralBadRequest)
		return
	}
	resp, err := h.usecase.Convert(context.Background(), payload)
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	api.ResponseSuccess(c, resp, "success")
}

func (h *HTTPHandler) GetCurrencies(c *gin.Context) {
	var (
		loadAll        bool
		startingAfter  int
		startingBefore int
		limit          int
		err            error
	)

	startingAfterQuery := c.Query("startingAfter")
	if startingAfterQuery == "" {
		startingAfterQuery = "0"
	}
	startingAfter, err = strconv.Atoi(startingAfterQuery)
	if err != nil {
		api.ResponseFailed(c, errutil.New(errutil.ErrGeneralBadRequest, err, "startingAfter should integer"))
		return
	}

	startingBeforeQuery := c.Query("startingBefore")
	if startingBeforeQuery == "" {
		startingBeforeQuery = "0"
	}

	startingBefore, err = strconv.Atoi(startingBeforeQuery)
	if err != nil {
		api.ResponseFailed(c, errutil.New(errutil.ErrGeneralBadRequest, err, "startingBefore should integer"))
		return
	}

	limitQuery := c.Query("limit")
	if limitQuery == "" {
		limitQuery = "0"
	}
	limit, err = strconv.Atoi(limitQuery)
	if err != nil {
		api.ResponseFailed(c, errutil.New(errutil.ErrGeneralBadRequest, err, "limit should integer"))
		return
	}

	all, ok := c.GetQuery("all")
	if ok && all == "true" {
		loadAll = true
	}

	resp, err := h.usecase.GetCurrencies(context.Background(), &entity.PaginationRequest{
		StartingAfter:  startingAfter,
		StartingBefore: startingBefore,
		Limit:          limit,
		All:            loadAll,
	})
	if err != nil {
		api.ResponseFailed(c, err)
		return
	}

	api.ResponseSuccess(c, resp, "success")
}
