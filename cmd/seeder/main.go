package main

import (
	"awesomeProject2/internal/shared"
	"awesomeProject2/internal/updater"
	"awesomeProject2/internal/utils"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	items := load()

	cdn := updater.NewRepositoryCDN(os.Getenv("STORAGE_PATH"))
	db := updater.NewRepositoryDB(database)

	for _, item := range items {
		save(item, cdn, db)
	}

	log.Println("Seeder completed successfully")
}

func load() []updater.APODApiResponse {
	startDate := os.Getenv("SEEDER_DATE_START")
	endDate := os.Getenv("SEEDER_DATE_END")

	response, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=" + os.Getenv("NASA_API_KEY") + "&start_date=" + startDate + "&end_date=" + endDate)
	if err != nil {
		log.Fatalf("error sending request to nasa api: %v\n", err)
		return nil
	}

	str, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error reading response from nasa api: %v\n", err)
		return nil
	}

	var responseData []updater.APODApiResponse
	if err = json.Unmarshal(str, &responseData); err != nil {
		log.Fatalf("error unmarshaling response from nasa api: %v\n", err)
		return nil
	}

	return responseData
}

func save(item updater.APODApiResponse, cdn updater.RepositoryCDN, db updater.RepositoryDB) {
	date, err := time.Parse("2006-01-02", item.Date)
	if err != nil {
		log.Printf("error parsing date: %v\n", err)
		return
	}

	extension := item.Hdurl[strings.LastIndex(item.Hdurl, "."):]
	path := utils.RandomString(10) + "_" + item.Date + extension
	err = cdn.SaveFromURL(item.Hdurl, path)

	event := shared.APODEvent{
		Date:        date,
		Explanation: item.Explanation,
		Title:       item.Title,
		Picture:     path,
	}

	if err = db.Save(context.Background(), event); err != nil {
		log.Fatalf("error saving data to db: %v\n", err)
	}
}
