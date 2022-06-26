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
