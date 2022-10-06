package config

import (
	"golang_basic_gorm_gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:H3ru@mysql@tcp(127.0.0.1:3306)/golang_basic_gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Department{}, &models.Position{}, &models.User{}, &models.Employee{},
		&models.Inventory{}, &models.Archive{}, &models.EmployeeInventory{})

	// DB.Create(&models.Department{
	// 	Name: "Human Resource",
	// 	Code: "HRD",
	// 	Positions: []models.Position{
	// 		{Name: "General Manager", Code: "GM HRD"},
	// 		{Name: "Manager", Code: "M HRD"},
	// 	},
	// })
	// DB.Create(&models.Department{
	// 	Name: "Finance",
	// 	Code: "FIN",
	// 	Positions: []models.Position{
	// 		{Name: "General Manager", Code: "GM HRD"},
	// 		{Name: "Manager", Code: "M HRD"},
	// 	},
	// })

	// DB.Create(&models.Employee{
	// 	Name:       "Heru",
	// 	Address:    "Cawang",
	// 	Email:      "heru@mail.com",
	// 	PositionID: 1,
	// })

	// DB.Create(&models.Inventory{
	// 	Name:        "Honda Supra X - Hitam",
	// 	Description: "B 1231 AC",
	// 	Archive: models.Archive{
	// 		Name:        "STNK & BPKB B 1231 AC",
	// 		Description: "Honda Supra X - Hitam",
	// 	},
	// })

	// DB.Create(&models.Inventory{
	// 	Name:        "Honda Vario - Hitam",
	// 	Description: "B 1232 AC",
	// 	Archive: models.Archive{
	// 		Name:        "STNK & BPKB B 1232 AC",
	// 		Description: "Honda Vario - Hitam",
	// 	},
	// })
}
