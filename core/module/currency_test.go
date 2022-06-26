package module

import (
	"context"
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
			err: errutil.New(errutil.ErrGeneralBadRequest, errutil.ErrGeneralBadRequest),
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
			assert.Equal(t, errutil.GetTypeErr(test.err), errutil.GetTypeErr(err))
		})
	}
}

func TestCreateConversionRate(t *testing.T) {
	tests := map[string]struct {
		req *entity.CreateCurrencyConversionRate
		err error
	}{
		"success": {
			req: &entity.CreateCurrencyConversionRate{
				CurrencyConversionRate: &entity.CurrencyConversionRate{
					From: 1,
					To:   2,
					Rate: 20,
				},
				CreatedBy: "t1",
			},
			err: nil,
		},
		"already exist": {
			req: &entity.CreateCurrencyConversionRate{
				CurrencyConversionRate: &entity.CurrencyConversionRate{
					From: 2,
					To:   1,
					Rate: 20,
				},
				CreatedBy: "t2",
			},
			err: errutil.New(errutil.ErrGeneralBadRequest, errutil.ErrGeneralBadRequest),
		},
	}
	repo := &currencyRepo.RepositoryMock{Mock: mock.Mock{}}
	validator := &validation.ValidatorMock{Mock: mock.Mock{}}
	svc := New(repo, validator)

	repo.Mock.On("FindConversionRateByFromTo", context.Background(), 1, 2).Return(nil, errutil.New(errutil.ErrGeneralNotFound, errutil.ErrGeneralNotFound))
	repo.Mock.On("FindConversionRateByFromTo", context.Background(), 2, 1).Return(&entity.CurrencyConversionRate{
		From: 2,
		To:   1,
		Rate: 20,
	}, nil)
	validator.Mock.On("ValidateParam", &entity.CreateCurrencyConversionRate{
		CurrencyConversionRate: &entity.CurrencyConversionRate{
			From: 1,
			To:   2,
			Rate: 20,
		},
		CreatedBy: "t1",
	}).Return(nil)
	validator.Mock.On("ValidateParam", &entity.CreateCurrencyConversionRate{
		CurrencyConversionRate: &entity.CurrencyConversionRate{
			From: 2,
			To:   1,
			Rate: 20,
		},
		CreatedBy: "t2",
	}).Return(nil)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := svc.CreateConversionRate(context.Background(), test.req)
			assert.Equal(t, errutil.GetTypeErr(test.err), errutil.GetTypeErr(err))
		})
	}
}

func TestConvert(t *testing.T) {
	tests := map[string]struct {
		req *entity.ConvertRequest
		err error
		out float64
	}{
		"success": {
			req: &entity.ConvertRequest{
				From:   1,
				To:     2,
				Amount: 10,
			},
			err: nil,
			out: 20,
		},
		"not found": {
			req: &entity.ConvertRequest{
				From:   1,
				To:     3,
				Amount: 10,
			},
			err: errutil.New(errutil.ErrGeneralNotFound, errutil.ErrGeneralNotFound),
			out: 0,
		},
	}
	repo := &currencyRepo.RepositoryMock{Mock: mock.Mock{}}
	validator := &validation.ValidatorMock{Mock: mock.Mock{}}
	svc := New(repo, validator)
	repo.Mock.On("FindConversionRateByFromTo", context.Background(), 1, 2).Return(&entity.CurrencyConversionRate{
		From: 1,
		To:   2,
		Rate: 2,
	}, nil)
	repo.Mock.On("FindConversionRateByFromTo", context.Background(), 1, 3).Return(nil, errutil.New(errutil.ErrGeneralNotFound, errutil.ErrGeneralNotFound))
	validator.Mock.On("ValidateParam", &entity.ConvertRequest{
		From:   1,
		To:     2,
		Amount: 10,
	}).Return(nil)
	validator.Mock.On("ValidateParam", &entity.ConvertRequest{
		From:   1,
		To:     3,
		Amount: 10,
	}).Return(nil)
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := svc.Convert(context.Background(), test.req)
			assert.Equal(t, errutil.GetTypeErr(test.err), errutil.GetTypeErr(err))
			assert.Equal(t, test.out, out)
		})
	}
}
