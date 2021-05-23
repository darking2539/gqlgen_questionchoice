package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func GetDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=boss123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
