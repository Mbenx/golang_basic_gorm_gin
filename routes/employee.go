package routes

import (
	"golang_basic_gorm_gin/config"
	"golang_basic_gorm_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetEmployee(c *gin.Context) {
	Employee := []models.Employee{}
	config.DB.Preload(clause.Associations).Find(&Employee)
	// config.DB.Preload("Positions").Find(&Employee)

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Employees",
		"data":    Employee,
	})
}

func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")
	var Employee models.Employee
	data := config.DB.Preload("Position").First(&Employee, "id = ?", id)
	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    Employee,
	})
}

func InsertEmployee(c *gin.Context) {
	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": "Bad Request",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	insert := config.DB.Create(&employee)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Internal Server Error",
			"error":    insert.Error.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    employee,
		"message": "Insert Success",
	})
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var RequestEmployee models.Employee

	if err := c.ShouldBindJSON(&RequestEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": "Bad Request",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	var DataEmployee models.Employee
	data := config.DB.First(&DataEmployee, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}
	config.DB.Model(&DataEmployee).Where("id = ?", id).Update("Name", RequestEmployee.Name)
	config.DB.Model(&DataEmployee).Where("id = ?", id).Update("Address", RequestEmployee.Address)
	config.DB.Model(&DataEmployee).Where("id = ?", id).Update("Email", RequestEmployee.Email)
	config.DB.Model(&DataEmployee).Where("id = ?", id).Update("PositionID", RequestEmployee.PositionID)

	c.JSON(200, gin.H{
		"message": "Update Success",
		"data":    DataEmployee,
	})
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	var Employee models.Employee

	data := config.DB.First(&Employee, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	config.DB.Delete(&Employee, id)

	c.JSON(200, gin.H{
		"message": "Delete Success",
	})
}
