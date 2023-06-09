package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"inventory-app-be/internal/config"
	"inventory-app-be/internal/ginhttp"
	repository "inventory-app-be/internal/repository/postgres/inventory"
	"inventory-app-be/internal/service"
	"log"
	"time"
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
	g := gin.Default()

	ginhttp.NewServer(g, svc).Run(context.Background())
}
