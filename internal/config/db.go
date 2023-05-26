package config

import (
	"inventory-app-be/internal/models"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONFIG_LOCAL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database")
	}

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(models.User{})
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Success connecting to database")

	return db
}
