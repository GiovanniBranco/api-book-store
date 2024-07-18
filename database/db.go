package database

import (
	"github.com/GiovanniBranco/classroom-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connectionString := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&models.Student{})

}
