package currency

import (
	"time"

	"github.com/faruqfadhil/currency-api/core/entity"
)

type Currency struct {
	ID        int
	Name      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

func (c Currency) FromCreateCurrencyRequestEntity(currency *entity.CreateCurrencyRequest) *Currency {
	if currency == nil {
		return nil
	}
	now := time.Now().UTC()
	return &Currency{
		ID:        currency.ID,
		Name:      currency.Name,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: currency.CreatedBy,
		UpdatedBy: currency.CreatedBy,
	}
}

func (c *Currency) ToEntity() *entity.Currency {
	return &entity.Currency{
		ID:   c.ID,
		Name: c.Name,
	}
}

type CurrencyConversionRate struct {
	FromCurrencyID int
	ToCurrencyID   int
	Rate           float64
	CreatedAt      time.Time
	CreatedBy      string
	UpdatedAt      time.Time
	UpdatedBy      string
}

func (c CurrencyConversionRate) FromCreateCurrencyConversionRateRequestEntity(req *entity.CreateCurrencyConversionRate) *CurrencyConversionRate {
	if req == nil {
		return nil
	}
	now := time.Now().UTC()
	return &CurrencyConversionRate{
		FromCurrencyID: req.From,
		ToCurrencyID:   req.To,
		Rate:           req.Rate,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      req.CreatedBy,
		UpdatedBy:      req.CreatedBy,
	}
}

func (c *CurrencyConversionRate) ToEntity() *entity.CurrencyConversionRate {
	return &entity.CurrencyConversionRate{
		From: c.FromCurrencyID,
		To:   c.ToCurrencyID,
		Rate: c.Rate,
	}
}
