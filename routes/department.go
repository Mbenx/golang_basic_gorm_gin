package routes

import (
	"golang_basic_gorm_gin/config"
	"golang_basic_gorm_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type DepartmentsResponse struct {
// 	ID        uint   `json:"id"`
// 	Name      string `json:"name"`
// 	Code      string `json:"code"`
// 	Positions []PositionResponse
// }

func GetDepartment(c *gin.Context) {
	departments := []models.Department{}
	config.DB.Find(&departments)

	config.DB.Preload("Positions").Find(&departments)

	// respDept := []DepartmentsResponse{}
	// position := []PositionResponse{}

	// for _, val := range departments {
	// 	post := PositionResponse{}
	// 	for _, val1 := range val.Positions {
	// 		post.DepartmentID = val1.DepartmentID
	// 		post.ID = val1.ID
	// 		post.Name = val1.Name
	// 		post.Code = val1.Code

	// 	}
	// 	position = append(position, post)

	// 	dept := DepartmentsResponse{
	// 		ID:        val.ID,
	// 		Name:      val.Name,
	// 		Code:      val.Code,
	// 		Positions: position,
	// 	}
	// 	respDept = append(respDept, dept)
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to department",
		"data":    departments,
	})
}

func GetDepartmentByID(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	data := config.DB.Preload("Positions").First(&department, "id = ?", id)

	// data := DB.First(&department, id)

	// respDept := DepartmentsResponse{
	// 	ID:   department.ID,
	// 	Name: department.Name,
	// 	Code: department.Code,
	// }

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    department,
	})
}

func InsertDepartment(c *gin.Context) {
	department := models.Department{
		Name: c.PostForm("name"),
		Code: c.PostForm("code"),
	}

	config.DB.Create(&department)

	// respDept := DepartmentsResponse{
	// 	ID:   department.ID,
	// 	Name: department.Name,
	// 	Code: department.Code,
	// }

	c.JSON(http.StatusCreated, gin.H{
		"data":    department,
		"message": "Insert Success",
	})
}

func UpdateDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	data := config.DB.First(&department, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	// config.DB.Model(&department).Where("id = ?").Update(models.Department{
	// 	Name: c.PostForm("name"),
	// 	Code: c.PostForm("code"),
	// })

	config.DB.Model(&department).Where("id = ?", id).Update("Name", c.PostForm("name"))
	config.DB.Model(&department).Where("id = ?", id).Update("Code", c.PostForm("code"))

	// respDept := DepartmentsResponse{
	// 	ID:   department.ID,
	// 	Name: department.Name,
	// 	Code: department.Code,
	// }

	c.JSON(200, gin.H{
		"message": "Update Success",
		"data":    department,
	})
}

func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")

	var department models.Department

	data := config.DB.First(&department, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": "Data Not Found",
		})

		return
	}

	config.DB.Delete(&department, id)

	c.JSON(200, gin.H{
		"message": "Delete Success",
	})
}
