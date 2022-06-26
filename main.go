package main

import (
	"fmt"
	"log"

	"github.com/faruqfadhil/currency-api/config"
	"github.com/faruqfadhil/currency-api/core/module"
	"github.com/faruqfadhil/currency-api/handler"
	"github.com/faruqfadhil/currency-api/pkg/validation"
	currencyRepo "github.com/faruqfadhil/currency-api/repository/currency"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Get()
	db := conn(&cfg)
	repo := currencyRepo.New(db)
	internalValidator := validation.NewGoValidator(validator.New())
	usecase := module.New(repo, internalValidator)
	handler := handler.New(usecase)
	router := gin.Default()
	apiv1 := router.Group("/api/v1/currency")
	{
		apiv1.POST("/create", handler.CreateCurrency)
	}
	router.Run(fmt.Sprintf(":%s", cfg.Port))
}

func conn(cfg *config.Config) *gorm.DB {
	defaultParams := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.CurrencyDBUsername, cfg.CurrencyDBPassword, cfg.CurrencyDBHost, cfg.CurrencyDBPort, cfg.CurrencyDBName, defaultParams)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error when try to connect db: %v", err)
	}
	return db
}
