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
	currency       table = "currency"
	conversionRate table = "currency_conversion_rate"
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
		return errutil.New(errutil.ErrGeneralDB, fmt.Errorf("[Insert] err: %v", err))
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

func (r *repository) InsertConversionRates(ctx context.Context, reqs []*entity.CreateCurrencyConversionRate) error {
	var models []*CurrencyConversionRate
	for _, req := range reqs {
		models = append(models, CurrencyConversionRate{}.FromCreateCurrencyConversionRateRequestEntity(req))
	}
	if err := r.db.Debug().Table(conversionRate.String()).Create(&models).Error; err != nil {
		return errutil.New(errutil.ErrGeneralDB, fmt.Errorf("[InsertConversionRates] err: %v", err))
	}
	return nil
}

func (r *repository) FindConversionRateByFromTo(ctx context.Context, from, to int) (*entity.CurrencyConversionRate, error) {
	var out CurrencyConversionRate
	err := r.db.Debug().Table(conversionRate.String()).
		Where("from_currency_id = ?", from).
		Where("to_currency_id = ?", to).
		Where("is_deleted = ?", false).
		First(&out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errutil.New(errutil.ErrGeneralNotFound, fmt.Errorf("[FindConversionRateByFromTo] err: %v", err))
		}
		return nil, errutil.New(errutil.ErrGeneralDB, fmt.Errorf("[FindConversionRateByFromTo] err: %v", err))
	}
	return out.ToEntity(), nil
}
