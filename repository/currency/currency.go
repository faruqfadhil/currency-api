package currency

import (
	"context"

	"github.com/faruqfadhil/currency-api/core/entity"
	repoInterface "github.com/faruqfadhil/currency-api/core/repository"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) repoInterface.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Insert(ctx context.Context, req *entity.CreateCurrencyRequest) error {
	return nil
}

func (r *repository) FindByID(ctx context.Context, ID string) (*entity.Currency, error) {
	return nil, nil
}
