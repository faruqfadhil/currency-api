package module

import (
	"context"
	"errors"
	"fmt"

	"github.com/faruqfadhil/currency-api/core/entity"
	"github.com/faruqfadhil/currency-api/core/repository"
	errutil "github.com/faruqfadhil/currency-api/pkg/error"
)

type Usecase interface {
	CreateCurrency(ctx context.Context, req *entity.CreateCurrencyRequest) error
}

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) CreateCurrency(ctx context.Context, req *entity.CreateCurrencyRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	existingCurrency, err := u.repo.FindByID(ctx, req.ID)
	if err != nil && !errors.Is(errutil.ErrGeneralNotFound, errutil.GetTypeErr(err)) {
		return err
	}
	if existingCurrency != nil {
		return errutil.New(errutil.ErrGeneralBadRequest, fmt.Errorf("currency ID already exist"), "currency ID already exist")
	}

	return u.repo.Insert(ctx, req)
}
