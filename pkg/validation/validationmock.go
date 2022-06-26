package validation

import "github.com/stretchr/testify/mock"

type ValidatorMock struct {
	Mock mock.Mock
}

func (v *ValidatorMock) ValidateParam(param interface{}) error {
	args := v.Mock.Called(param)
	if args.Error(0) != nil {
		return args.Error(0)
	}
	return nil
}
