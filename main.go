package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faruqfadhil/currency-api/core/module"
	"github.com/faruqfadhil/currency-api/handler"
	"github.com/faruqfadhil/currency-api/pkg/validation"
	currencyRepo "github.com/faruqfadhil/currency-api/repository/currency"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load env, err: %v", err)
	}

	db := conn()
	repo := currencyRepo.New(db)
	internalValidator := validation.NewGoValidator(validator.New())
	usecase := module.New(repo, internalValidator)
	handler := handler.New(usecase)
	router := gin.Default()
	apiv1 := router.Group("/api/v1/currency")
	{
		apiv1.POST("/create", handler.CreateCurrency)
		apiv1.POST("/conversion/create", handler.CreateConversionRate)
		apiv1.POST("/conversion/convert", handler.Convert)
		apiv1.GET("/list", handler.GetCurrencies)
	}
	router.Run(fmt.Sprintf(":%s", os.Getenv("GIN_PORT")))
}

func conn() *gorm.DB {
	defaultParams := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"), defaultParams)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error when try to connect db: %v", err)
	}
	return db
}
