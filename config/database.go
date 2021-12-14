package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
	return db
}
