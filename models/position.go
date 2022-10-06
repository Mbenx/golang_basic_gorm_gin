package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name         string
	Code         string `gorm:"unique_index"`
	Employees    []Employee
	DepartmentID uint
	Department   Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
