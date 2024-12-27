package models

import (
	"backend/utils/models"
	"time"
)

type User struct {
	models.Model `gorm:"embedded"`
	Username     string `json:"username" gorm:"uniqueIndex:idx_username" `
	Email        string `json:"email" gorm:"uniqueIndex:idx_customer_email" `
	Password     []byte `json:"-"`
}

// Transaction Model
type Transaction struct {
	models.Model
	UserId          uint64    `json:"user_id"`     // Foreign key to User
	CategoryId      uint64    `json:"category_id"` // Foreign key to Category
	Amount          float64   `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
	Type            string    `json:"type"` // Enum: "income" or "expense"
	Description     string    `json:"description"`
}

// Category Model
type Category struct {
	models.Model
	UserId uint64 `gorm:"not null" json:"user_id"` // Foreign key to User
	Name   string `gorm:"not null" json:"name"`
}

// Budget Model
type Budget struct {
	models.Model
	UserId     uint64    `gorm:"not null" json:"user_id"`              // Foreign key to User
	CategoryId uint64    `gorm:"not null" json:"category_id"`          // Foreign key to Category
	Amount     float64   `gorm:"default:0;not null" json:"amount"`     // Budgeted amount
	StartDate  time.Time `gorm:"type:date;not null" json:"start_date"` // Budget start date
	EndDate    time.Time `gorm:"type:date;not null" json:"end_date"`   // Budget end date
}

// Report Model
type Report struct {
	models.Model
	UserId     uint64    `gorm:"not null" json:"user_id"`              // Foreign key to User
	CategoryId uint64    `gorm:"not null" json:"category_id"`          // Foreign key to Category
	StartDate  time.Time `gorm:"type:date;not null" json:"start_date"` // Report period start
	EndDate    time.Time `gorm:"type:date;not null" json:"end_date"`   // Report period end
	ReportData string    `json:"report_data"`                          // JSON or text data summarizing transactions, budget compliance, etc.
}
