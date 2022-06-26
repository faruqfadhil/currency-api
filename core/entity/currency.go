package entity

import (
	"fmt"

	errutil "github.com/faruqfadhil/currency-api/pkg/error"
)

type Currency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCurrencyRequest struct {
	*Currency
	CreatedBy string `json:"-"`
}

func (c *CreateCurrencyRequest) Validate() error {
	if c.ID == 0 {
		return errutil.New(errutil.ErrGeneralBadRequest, fmt.Errorf("ID is required"), "ID is required")
	}
	if c.Name == "" {
		return errutil.New(errutil.ErrGeneralBadRequest, fmt.Errorf("name is required"), "Name is required")
	}
	return nil
}
