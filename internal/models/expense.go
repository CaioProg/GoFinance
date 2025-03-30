package models

import (
	"time"
)

type Expense struct {
	Id          uint      `gorm:"primaryKey"`
	UserId      uint32    `gorm:"not null"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"null"`
	CategoryId  uint32    `gorm:"null"`
	Date        string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

func (Expense) TableName() string {
	return "expenses"
}
