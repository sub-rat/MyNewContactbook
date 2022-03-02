package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=postgres password=password dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&Contact{},
		&Address{},
		&Phone{},
	)
	DB = db
	return db
}
