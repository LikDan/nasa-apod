package main

import (
	"awesomeProject2/internal/api"
	"awesomeProject2/internal/updater"
	"log"
	"os"
	"time"
)

func main() {
	time.Local = time.FixedZone("GMT", 0)

	updaterImpl := updater.NewUpdater(os.Getenv("NASA_API_KEY"), os.Getenv("STORAGE_PATH"), os.Getenv("API_CRON_UPDATE_PATTERN"), database)
	updaterImpl.Start()

	api.NewApi(group.Group("apod"), os.Getenv("HOST"), os.Getenv("STORAGE_PATH"), database)

	log.Fatalf("error launching webserver: %v\n", engine.Run())
}
