package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func LoadDB() *sql.DB {
	// db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONFIG_LOCAL")), &gorm.Config{})
	db, err := sql.Open("postgres", os.Getenv("DB_CONFIG_LOCAL"))
	if err != nil {
		log.Fatal("Error connecting database")
	}

	if err != nil {
		panic(err)
	}

	// err = db.AutoMigrate(&models.User{})
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	log.Printf("Success connecting to database")

	return db
}
