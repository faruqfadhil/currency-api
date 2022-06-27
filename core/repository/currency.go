package repository

import (
	"context"

	"github.com/faruqfadhil/currency-api/core/entity"
)

type Repository interface {
	Insert(ctx context.Context, req *entity.CreateCurrencyRequest) error
	FindByID(ctx context.Context, ID int) (*entity.Currency, error)
	InsertConversionRates(ctx context.Context, reqs []*entity.CreateCurrencyConversionRate) error
	FindConversionRateByFromTo(ctx context.Context, from, to int) (*entity.CurrencyConversionRate, error)
	FindCurrencies(ctx context.Context, pagination *entity.PaginationRequest) (*entity.CurrencyList, error)
}
