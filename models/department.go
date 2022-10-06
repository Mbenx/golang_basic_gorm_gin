package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name      string
	Code      string `gorm:"unique_index"`
	Positions []Position
}
