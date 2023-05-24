package cmd

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func main() {
	locJakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalf("config: failed to load Asia/Jakarta location error=%s", err)
	}

	//inject repo

	//inject service

	time.Local = locJakarta
	_ = echo.New()
	//echohttp.NewServer(e, studentSvc).Run(context.Background())
}
