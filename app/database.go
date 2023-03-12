package app

import (
	"fmt"
	"github.com/xvbnm48/go-learn-api-scalable/helper"
	"github.com/xvbnm48/go-learn-api-scalable/model/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = "localhost"
	user   = "postgres"
	pass   = "fariz"
	port   = "5432"
	dbName = "go-api-scalable"
)

func NewDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, port)
	db, err := gorm.Open(postgres.Open(config))
	helper.PanicIfError(err)

	err = db.AutoMigrate(&domain.Order{}, &domain.Item{})
	helper.PanicIfError(err)

	return db
}
