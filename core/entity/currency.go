package entity

import (
	errutil "github.com/faruqfadhil/currency-api/pkg/error"
)

type Currency struct {
	ID   int    `json:"id" validate:"gt=0,required"`
	Name string `json:"name" validate:"required"`
}

type CreateCurrencyRequest struct {
	*Currency `validate:"required"`
	CreatedBy string `json:"-"`
}

type CurrencyConversionRate struct {
	From int     `json:"from" validate:"required"`
	To   int     `json:"to" validate:"required"`
	Rate float64 `json:"rate" validate:"gte=0"`
}

type CreateCurrencyConversionRate struct {
	*CurrencyConversionRate `validate:"required"`
	CreatedBy               string `json:"-"`
}

func (c *CreateCurrencyConversionRate) MakeOppositeConversion() *CreateCurrencyConversionRate {
	return &CreateCurrencyConversionRate{
		CurrencyConversionRate: &CurrencyConversionRate{
			From: c.To,
			To:   c.From,
			Rate: 1 / c.Rate,
		},
		CreatedBy: c.CreatedBy,
	}
}

type ConvertRequest struct {
	From   int     `json:"from" validate:"required"`
	To     int     `json:"to" validate:"required"`
	Amount float64 `json:"amount" validate:"gte=0"`
}

type Pagination struct {
	TotalItems int `json:"totalItems"`
	TotalPage  int `json:"totalPage"`
}

// PaginationRequest represent the pagination request, we use seek method for paginate the data.
// Detail about seek methos: https://www.slideshare.net/MarkusWinand/p2d2-pagination-done-the-postgresql-way?ref=https://use-the-index-luke.com/no-offset
type PaginationRequest struct {
	// StartingAfter used to fetch next page by given the last ID in current page.
	StartingAfter int `json:"startingAfter"`

	// StartingBefore used to fetch previous page by given the first ID in current page.
	StartingBefore int  `json:"startingBefore"`
	Limit          int  `json:"limit"`
	All            bool `json:"all"`
}

func (p *PaginationRequest) Validate() error {
	if !p.All && p.Limit < 1 {
		return errutil.New(errutil.ErrGeneralBadRequest, errutil.ErrGeneralBadRequest, "limit should be greater than 0")
	}
	if !p.All && p.StartingAfter < 0 {
		return errutil.New(errutil.ErrGeneralBadRequest, errutil.ErrGeneralBadRequest, "startingAfter should be greater than equal 0")
	}
	if !p.All && p.StartingBefore < 0 {
		return errutil.New(errutil.ErrGeneralBadRequest, errutil.ErrGeneralBadRequest, "startingBefore should be greater than equal 0")
	}
	return nil
}

type CurrencyList struct {
	Currencies []*Currency `json:"currencies"`
	Pagination *Pagination `json:"pagination"`
}
