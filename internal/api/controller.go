package api

import (
	"awesomeProject2/internal/shared"
	"context"
	"errors"
	"time"
)

var (
	ValidationErr = errors.New("validation error")
)

type Controller interface {
	GetByDate(ctx context.Context, date string) (shared.APODEvent, error)
	GetBetweenDateRange(ctx context.Context, fromDate string, toDate string) ([]shared.APODEvent, error)
	GetFileFromStorage(ctx context.Context, path string) ([]byte, error)
}

type controller struct {
	cdn RepositoryCDN
	db  RepositoryDB

	basePath string
}

func NewController(cdn RepositoryCDN, db RepositoryDB, basePath string) Controller {
	return &controller{
		cdn:      cdn,
		db:       db,
		basePath: basePath,
	}
}

func (c *controller) GetByDate(ctx context.Context, dateStr string) (shared.APODEvent, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return shared.APODEvent{}, ValidationErr
	}

	item, err := c.db.GetByDate(ctx, date)
	if err != nil {
		return shared.APODEvent{}, err
	}

	item.Picture = c.applyBasePath(item.Picture)
	return item, nil
}

func (c *controller) GetBetweenDateRange(ctx context.Context, fromDateStr string, toDateStr string) ([]shared.APODEvent, error) {
	fromDate, err := time.Parse("2006-01-02", fromDateStr)
	if err != nil {
		return nil, ValidationErr
	}

	toDate, err := time.Parse("2006-01-02", toDateStr)
	if err != nil {
		return nil, ValidationErr
	}

	if fromDate.Day() >= toDate.Day() {
		return nil, ValidationErr
	}

	items, err := c.db.GetBetweenDateRange(ctx, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	for i := range items {
		items[i].Picture = c.applyBasePath(items[i].Picture)
	}

	return items, nil
}

func (c *controller) GetFileFromStorage(ctx context.Context, path string) ([]byte, error) {
	return c.cdn.GetByPath(ctx, path)
}

func (c *controller) applyBasePath(path string) string {
	return c.basePath + "/storage/" + path
}
