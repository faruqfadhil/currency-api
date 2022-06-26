package repository

import (
	"context"

	"github.com/faruqfadhil/currency-api/core/entity"
)

type Repository interface {
	Insert(ctx context.Context, req *entity.CreateCurrencyRequest) error
	FindByID(ctx context.Context, ID string) (*entity.Currency, error)
}
