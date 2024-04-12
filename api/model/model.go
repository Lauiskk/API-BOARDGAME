package model

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name  string `json:"name"`
	Image string `json:"image"`
	Stock uint   `json:"stockTotal"`
	Price uint   `json:"pricePerDay"`
}

type Customer struct {
	gorm.Model
	Name     string    `json:"name"`
	Phone    string    `json:"phone" validate:"min=10,max=11"`
	CPF      string    `json:"cpf" validate:"len=11"`
	Birthday time.Time `json:"birthday"`
}

type Rental struct {
	gorm.Model
	CustomerID    uint `gorm:"foreignKey:CustomerID"`
	Customer      Customer
	GameID        uint `gorm:"foreignKey:GameID"`
	Game          Game
	RentDate      string `json:"rentDate"`
	DaysRented    uint   `json:"daysRented"`
	ReturnDate    string `json:"returnDate"`
	OriginalPrice uint   `json:"originalPrice"`
	DelayFee      uint   `json:"delayFee"`
}
