package entity

type Currency struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type CreateCurrencyRequest struct {
	*Currency `validate:"required"`
	CreatedBy string `json:"-"`
}

type CurrencyConversionRate struct {
	From int     `json:"from" validate:"required"`
	To   int     `json:"to" validate:"required"`
	Rate float64 `json:"rate" validate:"required"`
}

type CreateCurrencyConversionRate struct {
	*CurrencyConversionRate `validate:"required"`
	CreatedBy               string `json:"-"`
}

func (c *CreateCurrencyConversionRate) MakeOppositeConversion() *CreateCurrencyConversionRate {
	return &CreateCurrencyConversionRate{
		CurrencyConversionRate: &CurrencyConversionRate{
			From: c.To,
			To:   c.From,
			Rate: 1 / c.Rate,
		},
		CreatedBy: c.CreatedBy,
	}
}
