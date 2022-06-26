package entity

import (
	"fmt"

	errutil "github.com/faruqfadhil/currency-api/pkg/error"
)

type Currency struct {
}

type CreateCurrencyRequest struct {
	ID        string
	Name      string
	CreatedBy string
}

func (c *CreateCurrencyRequest) Validate() error {
	if c.ID == "" {
		return fmt.Errorf("%w:ID is required", errutil.ErrGeneralBadRequest)
	}
	if c.Name == "" {
		return fmt.Errorf("%w:Name is required", errutil.ErrGeneralBadRequest)
	}
	return nil
}
