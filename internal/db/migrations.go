package db

import (
	"log"

	"github.com/CaioProg/GoFinance/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatalf("Error while running migrations: %v", err)
	} else {
		log.Println("Migrations ran successfully!")
	}
}
