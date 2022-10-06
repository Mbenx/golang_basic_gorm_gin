package routes

import (
	"golang_basic_gorm_gin/config"
	"golang_basic_gorm_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestRental struct {
	EmployeeID  uint   `json:"employee_id"`
	InventoryID uint   `json:"inventory_id"`
	Description string `json:"description"`
}

func RentalByEmployee(c *gin.Context) {
	var reqRental RequestRental

	if err := c.ShouldBindJSON(&reqRental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": "Bad Request",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	rental := models.EmployeeInventory{
		EmployeeID:  reqRental.EmployeeID,
		InventoryID: reqRental.InventoryID,
		Description: reqRental.Description,
	}

	insert := config.DB.Create(&rental)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Internal Server Error",
			"error":    insert.Error.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    rental,
		"message": "Insert Success",
	})
}

func GetDataRental(c *gin.Context) {
	EmployeeInventory := []models.EmployeeInventory{}

	config.DB.Find(&EmployeeInventory)

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Data Rental",
		"data":    EmployeeInventory,
	})
}

func GetRentalByInventoryID(c *gin.Context) {
	id := c.Param("id")
	Inventory := []models.Inventory{}

	config.DB.Preload("Employees").First(&Inventory, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Data Rental",
		"data":    Inventory,
	})

}

func GetRentalByEmployeeID(c *gin.Context) {
	id := c.Param("id")
	Employee := []models.Employee{}

	config.DB.Preload("Inventories").First(&Employee, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Data Rental",
		"data":    Employee,
	})
}
