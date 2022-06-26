package currency

import (
	"time"

	"github.com/faruqfadhil/currency-api/core/entity"
)

type Currency struct {
	ID        string
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
