package module

import (
	"context"
	"fmt"
	"testing"

	"github.com/faruqfadhil/currency-api/core/entity"
	errutil "github.com/faruqfadhil/currency-api/pkg/error"
	"github.com/faruqfadhil/currency-api/pkg/validation"
	currencyRepo "github.com/faruqfadhil/currency-api/repository/currency"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCurrency(t *testing.T) {
	tests := map[string]struct {
		req *entity.CreateCurrencyRequest
		err error
	}{
		"success": {
			req: &entity.CreateCurrencyRequest{
				Currency: &entity.Currency{
					ID:   1,
					Name: "test1",
				},
				CreatedBy: "t1",
			},
			err: nil,
		},
		"already exist": {
			req: &entity.CreateCurrencyRequest{
				Currency: &entity.Currency{
					ID:   2,
					Name: "test2",
				},
				CreatedBy: "t2",
			},
			err: errutil.New(errutil.ErrGeneralBadRequest, fmt.Errorf("currency ID already exist"), "currency ID already exist"),
		},
		"bad request empty ID": {
			req: &entity.CreateCurrencyRequest{
				Currency: &entity.Currency{
					Name: "test3",
				},
				CreatedBy: "t3",
			},
			err: errutil.New(errutil.ErrGeneralBadRequest, errutil.ErrGeneralBadRequest),
		},
	}

	repo := &currencyRepo.RepositoryMock{Mock: mock.Mock{}}
	validator := &validation.ValidatorMock{Mock: mock.Mock{}}
	svc := New(repo, validator)
	repo.Mock.On("FindByID", context.Background(), 1).Return(nil, errutil.New(errutil.ErrGeneralNotFound, errutil.ErrGeneralNotFound))
	repo.Mock.On("FindByID", context.Background(), 2).Return(&entity.Currency{
		ID:   2,
		Name: "test2",
	}, nil)
	validator.Mock.On("ValidateParam", &entity.CreateCurrencyRequest{
		Currency: &entity.Currency{
			ID:   1,
			Name: "test1",
		},
		CreatedBy: "t1",
	}).Return(nil)
	validator.Mock.On("ValidateParam", &entity.CreateCurrencyRequest{
		Currency: &entity.Currency{
			ID:   2,
			Name: "test2",
		},
		CreatedBy: "t2",
	}).Return(nil)
	validator.Mock.On("ValidateParam", &entity.CreateCurrencyRequest{
		Currency: &entity.Currency{
			Name: "test3",
		},
		CreatedBy: "t3",
	}).Return(errutil.New(errutil.ErrGeneralBadRequest, errutil.ErrGeneralBadRequest))

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := svc.CreateCurrency(context.Background(), test.req)
			assert.Equal(t, test.err, err)
		})
	}
}
