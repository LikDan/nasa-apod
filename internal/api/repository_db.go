package api

import (
	"awesomeProject2/internal/shared"
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

type RepositoryDB interface {
	GetByDate(ctx context.Context, date time.Time) (shared.APODEvent, error)
	GetBetweenDateRange(ctx context.Context, fromDate time.Time, toDate time.Time) ([]shared.APODEvent, error)
}

type repositoryDB struct {
	database *pgx.Conn
}

func NewRepositoryDB(database *pgx.Conn) RepositoryDB {
	return &repositoryDB{
		database: database,
	}
}

func (r *repositoryDB) GetByDate(ctx context.Context, date time.Time) (item shared.APODEvent, err error) {
	row := r.database.QueryRow(ctx, "SELECT id, date, title, explanation, picture FROM apod_events WHERE date = $1", date)
	err = row.Scan(&item.ID, &item.Date, &item.Title, &item.Explanation, &item.Picture)
	return
}

func (r *repositoryDB) GetBetweenDateRange(ctx context.Context, fromDate time.Time, toDate time.Time) (items []shared.APODEvent, err error) {
	result, err := r.database.Query(ctx, "SELECT id, date, title, explanation, picture FROM apod_events WHERE date >= $1 AND date <= $2", fromDate, toDate)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var item shared.APODEvent
		if err = result.Scan(&item.ID, &item.Date, &item.Title, &item.Explanation, &item.Picture); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return
}
