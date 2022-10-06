package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name                string       `json:"name"`
	Address             string       `json:"address"`
	Email               string       `json:"email"`
	PositionID          uint         `json:"position_id"`
	Position            Position     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Inventories         []*Inventory `gorm:"many2many:employee_inventories;"`
	EmployeeInventories []EmployeeInventory
}
