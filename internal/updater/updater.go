package updater

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewUpdater(apiKey string, storagePath string, cronPattern string, database *pgxpool.Pool) Controller {
	apiRepository := NewRepositoryApi(apiKey)
	dbRepository := NewRepositoryDB(database)
	cdnRepository := NewRepositoryCDN(storagePath)

	updaterController := NewController(apiRepository, cdnRepository, dbRepository, cronPattern)
	return updaterController
}
