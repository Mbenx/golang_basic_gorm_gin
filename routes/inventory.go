package routes

import (
	"golang_basic_gorm_gin/config"
	"golang_basic_gorm_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestInsertInventory struct {
	InventoryName        string `json:"inventory_name"`
	InventoryDescription string `json:"inventory_description"`
	ArchiveName          string `json:"archive_name"`
	ArchiveDescription   string `json:"archive_description"`
}

func GetInventory(c *gin.Context) {
	Inventory := []models.Inventory{}
	// config.DB.Find(&Positions)
	// config.DB.Preload(clause.Associations).Find(&Positions)
	config.DB.Preload("Archive").Find(&Inventory)

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Position",
		"data":    Inventory,
	})
}

func InsertInventory(c *gin.Context) {
	var reqInventory RequestInsertInventory

	if err := c.ShouldBindJSON(&reqInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": "Bad Request",
			"error":    err.Error(),
		})

		c.Abort()
		return
	}

	Inventory := models.Inventory{
		Name:        reqInventory.InventoryName,
		Description: reqInventory.InventoryDescription,
		Archive: models.Archive{
			Name:        reqInventory.ArchiveName,
			Description: reqInventory.InventoryName,
		}}

	insert := config.DB.Create(&Inventory)
	if insert.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": "Internal Server Error",
			"error":    insert.Error.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    Inventory,
		"message": "Insert Success",
	})
}
