package currency

import (
	"context"
	"errors"
	"fmt"

	"github.com/faruqfadhil/currency-api/core/entity"
	repoInterface "github.com/faruqfadhil/currency-api/core/repository"
	errutil "github.com/faruqfadhil/currency-api/pkg/error"
	"gorm.io/gorm"
)

type table string

func (t table) String() string {
	return string(t)
}

const (
	currency table = "currency"
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
	model := Currency{}.FromCreateCurrencyRequestEntity(req)
	if err := r.db.Table(currency.String()).Create(&model).Error; err != nil {
		return fmt.Errorf("%w:[Insert] err: %v", errutil.ErrGeneralDB, err)
	}
	return nil
}

func (r *repository) FindByID(ctx context.Context, ID string) (*entity.Currency, error) {
	var out Currency
	err := r.db.Table(currency.String()).
		Where("id = ?", ID).
		Find(&out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w:[FindByID] err: %v", errutil.ErrGeneralNotFound, err)
		}
		return nil, fmt.Errorf("%w:[FindByID] err: %v", errutil.ErrGeneralDB, err)
	}
	return out.ToEntity(), nil
}
