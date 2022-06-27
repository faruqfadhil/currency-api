package validation

import (
	"fmt"

	errutil "github.com/faruqfadhil/currency-api/pkg/error"
	"github.com/go-playground/validator/v10"
)

type govalidator struct {
	validator *validator.Validate
}

// NewGoValidator create new validator that implements InternalValidator inteface.
func NewGoValidator(v *validator.Validate) InternalValidator {
	return &govalidator{v}
}

// ValidateWithParam use go-playground/validator to validate request.
// See https://godoc.org/gopkg.in/go-playground/validator.v10 for more details.
func (v *govalidator) ValidateParam(param interface{}) error {
	if err := v.validator.Struct(param); err != nil {
		msg := ""
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				msg += fmt.Sprintf(" %s is required", e.Field())
			case "gte":
				msg += fmt.Sprintf(" %s should be greater than or equal %s", e.Field(), e.Param())
			case "gt":
				msg += fmt.Sprintf(" %s should be greater than %s", e.Field(), e.Param())
			default:
				msg += fmt.Sprintf(" %s", err.Error())
			}
		}
		return errutil.New(errutil.ErrGeneralBadRequest, fmt.Errorf(msg), msg)
	}
	return nil
}
