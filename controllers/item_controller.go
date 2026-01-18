package controllers

import (
	"lab-inventaris/config"
	"lab-inventaris/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetItemsByLab(c *gin.Context) {
	labId := c.Param("lab_id")
	var items []models.Item

	if err := config.DB.Where("lab_id = ?", labId).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Data tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data" : items})
}

func CreateItem(c *gin.Context)  {
	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Status = "Good"
	input.LastChecked = time.Now()

	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data" : input})
}

func UpdateItemStatus(c *gin.Context)  {
	id := c.Param("id")
	var item models.Item

	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Barang tidak ditemukan"})
		return
	}

	var input struct {
		Status string `json:"status"`
		Note  string `json:"note"`
		Admin  string `json:"admin_name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	config.DB.Model(&item).Updates(map[string]interface{}{
		"status": input.Status,
		"condition": input.Note,
		"last_checked": time.Now(),
	})

	log := models.MaintenanceLog{
		ItemId: item.ID,
		Status: input.Status,
		Note: input.Note,
		CheckedBy: input.Admin,
		CheckedAt: time.Now(),
	}
	config.DB.Create(&log)

	c.JSON(http.StatusOK, gin.H{"message" : "Status barang berhasil diupdate", "data":item})
}

type BatchUpdateInput struct{
	ID     uint   `json:"id"`
    Status string `json:"status"`
    Note   string `json:"note"`
}

func BatchUpdateItems(c *gin.Context)	{
	var inputs []BatchUpdateInput

	if err := c.ShouldBindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	tx := config.DB.Begin()
	for _, item := range inputs{
		if err := tx.Model(&models.Item{}).Where("id = ?", item.ID).Updates(map[string]interface{}{
			"status":       item.Status,
			"condition":    item.Note,
			"last_checked": time.Now(),
		}).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update item ID " + string(rune(item.ID))})
			return
		}
		log := models.MaintenanceLog{
			ItemId:    item.ID,
			Status:    item.Status,
			Note:      item.Note,
			CheckedBy: "Admin Batch", 
			CheckedAt: time.Now(),
		}
		tx.Create(&log)
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "Semua data berhasil disimpan!"})
}