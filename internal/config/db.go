package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func LoadDB() *sql.DB {
	// db, err := sql.Open("postgres", os.Getenv("DB_CONFIG_LOCAL"))
	db, err := sql.Open("postgres", "postgres://postgres:postgres@127.0.0.1/kreasitex-jakarta?sslmode=disable")
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
