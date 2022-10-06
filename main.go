package main

import (
	"golang_basic_gorm_gin/config"
	"golang_basic_gorm_gin/middlewares"
	"golang_basic_gorm_gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()

	router := gin.Default()

	router.GET("/", getHome)

	user := router.Group("/user")
	{
		user.POST("get-token", routes.GenerateToken)
		user.POST("register", routes.RegisterUser)
	}

	department := router.Group("/department").Use(middlewares.Auth())
	{
		department.GET("/", routes.GetDepartment)
		department.GET("/:id", routes.GetDepartmentByID)
		department.POST("/", routes.InsertDepartment)
		department.PUT("/:id", routes.UpdateDepartment)
		department.DELETE("/:id", routes.DeleteDepartment)
	}

	position := router.Group("/position").Use(middlewares.Auth())
	{
		position.GET("/", routes.GetPosition)
		position.GET("/:id", routes.GetPositionByID)
		position.POST("/", routes.InsertPosition)
		position.PUT("/:id", routes.UpdatePosition)
		position.DELETE("/:id", routes.DeletePosition)
	}

	employee := router.Group("/employee").Use(middlewares.Auth())
	{
		employee.GET("/", routes.GetEmployee)
		employee.GET("/:id", routes.GetEmployeeByID)
		employee.POST("/", routes.InsertEmployee)
		employee.PUT("/:id", routes.UpdateEmployee)
		employee.DELETE("/:id", routes.DeleteEmployee)
	}

	inventory := router.Group("/inventory").Use(middlewares.Auth())
	{
		inventory.GET("/", routes.GetInventory)
		// inventory.GET("/:id", routes.GetInventoryByID)
		inventory.POST("/", routes.InsertInventory)
		// inventory.PUT("/:id", routes.UpdateInventory)
		// inventory.DELETE("/:id", routes.DeleteInventory)
	}

	rental := router.Group("/rental").Use(middlewares.Auth())
	{
		rental.GET("/", routes.GetDataRental)
		rental.GET("/inventory/:id", routes.GetRentalByInventoryID)
		rental.GET("/employee/:id", routes.GetRentalByEmployeeID)
		rental.POST("/employee", routes.RentalByEmployee)
	}

	router.Run()
}

func getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":     "Welcome to Home",
		"description": "ini halaman home",
	})
}
