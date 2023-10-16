package updater

import (
	"awesomeProject2/internal/shared"
	"awesomeProject2/internal/utils"
	"context"
	"github.com/robfig/cron"
	"log"
	"strings"
	"time"
)

type Controller interface {
	Start()
}

type controller struct {
	api RepositoryApi
	cdn RepositoryCDN
	db  RepositoryDB

	cron *cron.Cron
}

func NewController(api RepositoryApi, cdn RepositoryCDN, db RepositoryDB, pattern string) Controller {
	controllerImpl := &controller{
		api: api,
		cdn: cdn,
		db:  db,
	}

	c := cron.New()
	err := c.AddFunc(pattern, controllerImpl.update)
	if err != nil {
		log.Fatalf("Eroor creating cron %v\n", err)
	}

	controllerImpl.cron = c

	return controllerImpl
}

func (c *controller) Start() {
	c.update()
	c.cron.Start()
}

func (c *controller) update() {
	log.Println("Updating information from API")
	ctx := context.Background()

	info, err := c.api.GetTodayInfo(ctx)
	if err != nil {
		log.Printf("error getting todays info from api: %v\n", err)
		return
	}

	date, err := time.Parse("2006-01-02", info.Date)
	if err != nil {
		log.Printf("error parsing date: %v\n", err)
		return
	}

	extension := info.Hdurl[strings.LastIndex(info.Hdurl, "."):]
	path := utils.RandomString(10) + "_" + info.Date + extension
	err = c.cdn.SaveFromURL(info.Hdurl, path)

	item := shared.APODEvent{
		Date:        date,
		Explanation: info.Explanation,
		Title:       info.Title,
		Picture:     path,
	}

	if err = c.db.Save(ctx, item); err != nil {
		log.Printf("error saving data to db: %v\n", err)
		return
	}

	log.Println("Updating information from API finished")
}
