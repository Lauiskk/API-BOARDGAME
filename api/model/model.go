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
	Price uint   `json:"totalPrice"`
}

type Customer struct {
	gorm.Model
	Name     string    `json:"name"`
	Phone    uint      `json:"phone"`
	CPF      uint      `json:"cpf"`
	Birthday time.Time `json:"birthday"`
}

type Rental struct {
	gorm.Model
	CustomerId    []Customer `gorm:"foreignKey:UserID"`
	GameId        []Game     `gorm:"foreignKey:UserID"`
	RentDate      time.Time  `json:"rentDate"`
	DaysRented    uint       `json:"daysRented"`
	ReturnDate    bool       `json:"returnDate"`
	OriginalPrice uint       `json:"originalPrice"`
	DelayFee      bool       `json:"delayFee"`
}
