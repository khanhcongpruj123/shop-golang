package model

import (
	"gorm.io/gorm"
)

type Production struct {
	gorm.Model
	Name   string `json:"name"`
	Price  uint   `json:"price"`
	Amount uint   `json:"amount"`
}
