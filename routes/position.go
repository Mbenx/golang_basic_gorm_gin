package routes

import (
	"golang_basic_gorm_gin/config"
	"golang_basic_gorm_gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosition(c *gin.Context) {
	Positions := []models.Position{}
	// config.DB.Find(&Positions)
	// config.DB.Preload(clause.Associations).Find(&Positions)
	config.DB.Preload("Department").Preload("Employees").Find(&Positions)

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Position",
		"data":    Positions,
	})
}

func GetPositionByID(c *gin.Context) {
	id := c.Param("id")

	var Position models.Position

	// data := config.DB.Preload(clause.Associations).First(&Position, "id = ?", id)
	data := config.DB.Preload("Department").First(&Position, "id = ?", id)

	// data := DB.First(&Position, id)
	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    Position,
	})
}

func InsertPosition(c *gin.Context) {
	department_id, err := strconv.ParseUint(c.PostForm("department_id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Internal Server Error",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	Position := models.Position{
		Name:         c.PostForm("name"),
		Code:         c.PostForm("code"),
		DepartmentID: uint(department_id),
	}

	config.DB.Create(&Position)

	c.JSON(http.StatusCreated, gin.H{
		"data":    Position,
		"message": "Insert Success",
	})
}

func UpdatePosition(c *gin.Context) {
	id := c.Param("id")
	department_id, err := strconv.ParseUint(c.PostForm("department_id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Internal Server Error",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	var Position models.Position

	data := config.DB.First(&Position, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	// config.DB.Model(&Position).Where("id = ?").Update(models.Position{
	// 	Name: c.PostForm("name"),
	// 	Code: c.PostForm("code"),
	// })

	config.DB.Model(&Position).Where("id = ?", id).Update("Name", c.PostForm("name"))
	config.DB.Model(&Position).Where("id = ?", id).Update("Code", c.PostForm("code"))
	config.DB.Model(&Position).Where("id = ?", id).Update("DepartmentID", department_id)

	c.JSON(200, gin.H{
		"message": "Update Success",
		"data":    Position,
	})
}

func DeletePosition(c *gin.Context) {
	id := c.Param("id")

	var Position models.Position

	data := config.DB.First(&Position, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	config.DB.Delete(&Position, id)

	c.JSON(200, gin.H{
		"message": "Delete Success",
	})
}
