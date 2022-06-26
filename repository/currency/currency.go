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
	if err := r.db.Debug().Table(currency.String()).Create(&model).Error; err != nil {
		return errutil.New(errutil.ErrGeneralNotFound, fmt.Errorf("[Insert] err: %v", err))
	}
	return nil
}

func (r *repository) FindByID(ctx context.Context, ID int) (*entity.Currency, error) {
	var out Currency
	err := r.db.Debug().Table(currency.String()).
		Where("id = ?", ID).
		First(&out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errutil.New(errutil.ErrGeneralNotFound, fmt.Errorf("[FindByID] err: %v", err))
		}
		return nil, errutil.New(errutil.ErrGeneralDB, fmt.Errorf("[FindByID] err: %v", err))
	}
	return out.ToEntity(), nil
}
