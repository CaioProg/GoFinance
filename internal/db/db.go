package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=finance_user password=finance_password dbname=finance_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	} else {
		fmt.Println("Database connection successful!")
	}
}
