package currency

import (
	"context"

	"github.com/faruqfadhil/currency-api/core/entity"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (r *RepositoryMock) Insert(ctx context.Context, req *entity.CreateCurrencyRequest) error {
	return nil
}

func (r *RepositoryMock) FindByID(ctx context.Context, ID int) (*entity.Currency, error) {
	args := r.Mock.Called(ctx, ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Currency), nil
}

func (r *RepositoryMock) InsertConversionRates(ctx context.Context, reqs []*entity.CreateCurrencyConversionRate) error {
	return nil
}

func (r *RepositoryMock) FindConversionRateByFromTo(ctx context.Context, from, to int) (*entity.CurrencyConversionRate, error) {
	args := r.Mock.Called(ctx, from, to)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.CurrencyConversionRate), nil
}

func (r *RepositoryMock) FindCurrencies(ctx context.Context, pagination *entity.PaginationRequest) (*entity.CurrencyList, error) {
	args := r.Mock.Called(ctx, pagination)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.CurrencyList), nil
}
