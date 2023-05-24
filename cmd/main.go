package main

import (
	"context"
	"inventory-app-be/internal/config"
	"inventory-app-be/internal/echohttp"
	repository "inventory-app-be/internal/repository/postgres/inventory"
	service "inventory-app-be/internal/service"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	locJakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalf("config: failed to load Asia/Jakarta location error=%s", err)
	}
	config.LoadEnvFile()

	db := config.LoadDB()
	// inject repo
	repo := repository.NewPostgresInventoryRepository(db)

	// inject service
	svc := service.NewService(repo)

	time.Local = locJakarta
	e := echo.New()

	echohttp.NewServer(e, svc).Run(context.Background())
}
