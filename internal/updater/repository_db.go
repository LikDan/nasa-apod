package updater

import (
	"awesomeProject2/internal/shared"
	"context"
	"github.com/jackc/pgx/v5"
)

type RepositoryDB interface {
	Save(ctx context.Context, item shared.APODEvent) error
}

type repositoryDB struct {
	database *pgx.Conn
}

func NewRepositoryDB(database *pgx.Conn) RepositoryDB {
	return &repositoryDB{
		database: database,
	}
}

func (r *repositoryDB) Save(ctx context.Context, item shared.APODEvent) error {
	conn, err := r.database.Query(ctx, "INSERT INTO apod_events (date, title, explanation, picture) VALUES ($1, $2, $3, $4)", item.Date, item.Title, item.Explanation, item.Picture)
	if err != nil {
		return err
	}

	conn.Close()
	return nil
}
