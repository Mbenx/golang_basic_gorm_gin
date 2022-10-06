package models

import (
	"time"

	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	Name                string
	Description         string
	Archive             Archive
	Employees           []*Employee `gorm:"many2many:employee_inventories;"`
	EmployeeInventories []EmployeeInventory
}

type EmployeeInventory struct {
	EmployeeID  uint
	InventoryID uint
	Description string
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
