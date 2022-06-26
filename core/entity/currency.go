package entity

type Currency struct {
	ID   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type CreateCurrencyRequest struct {
	*Currency `validate:"required"`
	CreatedBy string `json:"-"`
}
