package config

import (
	"fmt"
	"lab-inventaris/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=123 dbname=lab_inventaris port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database :", err)
	}

	database.AutoMigrate(&models.Lab{}, &models.Item{}, &models.MaintenanceLog{}, &models.User{})

	DB = database
	fmt.Println("Database Terkoneksi & Termigrasi")
}