package module

import (
	"context"
	"errors"
	"fmt"

	"github.com/faruqfadhil/currency-api/core/entity"
	"github.com/faruqfadhil/currency-api/core/repository"
	errutil "github.com/faruqfadhil/currency-api/pkg/error"
	"github.com/faruqfadhil/currency-api/pkg/validation"
)

type Usecase interface {
	CreateCurrency(ctx context.Context, req *entity.CreateCurrencyRequest) error
	CreateConversionRate(ctx context.Context, req *entity.CreateCurrencyConversionRate) error
}

type usecase struct {
	repo      repository.Repository
	validator validation.InternalValidator
}

func New(repo repository.Repository, validator validation.InternalValidator) Usecase {
	return &usecase{
		repo:      repo,
		validator: validator,
	}
}

func (u *usecase) CreateCurrency(ctx context.Context, req *entity.CreateCurrencyRequest) error {
	if err := u.validator.ValidateParam(req); err != nil {
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

func (u *usecase) CreateConversionRate(ctx context.Context, req *entity.CreateCurrencyConversionRate) error {
	if err := u.validator.ValidateParam(req); err != nil {
		return err
	}
	existingConversionRate, err := u.repo.FindConversionRateByFromTo(ctx, req.From, req.To)
	if err != nil && !errors.Is(errutil.ErrGeneralNotFound, errutil.GetTypeErr(err)) {
		return err
	}
	if existingConversionRate != nil {
		return errutil.New(errutil.ErrGeneralBadRequest, fmt.Errorf("conversion rate already exist"), "conversion rate already exist")
	}

	var payloads []*entity.CreateCurrencyConversionRate
	payloads = append(payloads, req, req.MakeOppositeConversion())
	return u.repo.InsertConversionRates(ctx, payloads)
}
