package currency

import (
	"context"
	"errors"
	"fmt"
	"math"

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

func (r *repository) FindCurrencies(ctx context.Context, pagination *entity.PaginationRequest) (*entity.CurrencyList, error) {
	if err := pagination.Validate(); err != nil {
		return nil, err
	}

	var currencies []*Currency
	var totalItems int64
	var totalPage float64
	baseQuery := r.db.Debug().Table(currency.String())
	err := baseQuery.Count(&totalItems).Error
	if err != nil {
		return nil, errutil.New(errutil.ErrGeneralDB, fmt.Errorf("[FindCurrencies]: %v", err))
	}
	if !pagination.All {
		totalPage = math.Ceil(float64(totalItems) / float64(pagination.Limit))
	}
	paginationData := &entity.Pagination{
		TotalItems: int(totalItems),
		TotalPage:  int(totalPage),
	}

	if pagination.All {
		err = baseQuery.Order("id asc").Find(&currencies).Error
		if err != nil {
			return nil, errutil.New(errutil.ErrGeneralDB, fmt.Errorf("[FindCurrencies]: %v", err))
		}
	}

	if !pagination.All && pagination.StartingAfter == 0 && pagination.StartingBefore == 0 {
		// Fetch first page.
		err = baseQuery.Order("id asc").Limit(pagination.Limit).Find(&currencies).Error
		if err != nil {
			return nil, errutil.New(errutil.ErrGeneralDB, fmt.Errorf("[FindCurrencies]: %v", err))
		}
	} else if !pagination.All && pagination.StartingAfter > 0 {
		// Fetch next page.
		err = baseQuery.
			Where("id > ?", pagination.StartingAfter).
			Order("id asc").
			Limit(pagination.Limit).
			Find(&currencies).Error
		if err != nil {
			return nil, errutil.New(errutil.ErrGeneralDB, fmt.Errorf("[FindCurrencies]: %v", err))
		}
	} else if !pagination.All && pagination.StartingBefore > 0 {
		// Fetch previous page.
		query := `
		SELECT * from (
			SELECT * from currency c
			WHERE c.id < ?
			ORDER BY c.id DESC
			limit ?
		) as t
		order by id
		`
		err = r.db.Raw(query, pagination.StartingBefore, pagination.Limit).Find(&currencies).Error
		if err != nil {
			return nil, errutil.New(errutil.ErrGeneralDB, fmt.Errorf("[FindCurrencies]: %v", err))
		}
	}

	var entityCurrencies []*entity.Currency
	for _, c := range currencies {
		entityCurrencies = append(entityCurrencies, c.ToEntity())
	}

	return &entity.CurrencyList{
		Currencies: entityCurrencies,
		Pagination: paginationData,
	}, nil
}
