package module

import (
	"context"
	"fmt"
	"testing"

	"github.com/faruqfadhil/currency-api/core/entity"
	errutil "github.com/faruqfadhil/currency-api/pkg/error"
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
			err: fmt.Errorf("%w:%s", errutil.ErrGeneralBadRequest, "currency ID already exist"),
		},
		"bad request empty ID": {
			req: &entity.CreateCurrencyRequest{
				Currency: &entity.Currency{
					Name: "test3",
				},
				CreatedBy: "t3",
			},
			err: fmt.Errorf("%w:ID is required", errutil.ErrGeneralBadRequest),
		},
		"bad request empty Name": {
			req: &entity.CreateCurrencyRequest{
				Currency: &entity.Currency{
					ID: 4,
				},
				CreatedBy: "t4",
			},
			err: fmt.Errorf("%w:Name is required", errutil.ErrGeneralBadRequest),
		},
	}

	repo := &currencyRepo.RepositoryMock{Mock: mock.Mock{}}
	svc := New(repo)
	repo.Mock.On("FindByID", context.Background(), 1).Return(nil, errutil.ErrGeneralNotFound)
	repo.Mock.On("FindByID", context.Background(), 2).Return(&entity.Currency{
		ID:   2,
		Name: "test2",
	}, nil)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := svc.CreateCurrency(context.Background(), test.req)
			assert.Equal(t, test.err, err)
		})
	}
}
