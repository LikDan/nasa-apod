package updater

import "github.com/jackc/pgx/v5"

func NewUpdater(apiKey string, storagePath string, cronPattern string, database *pgx.Conn) Controller {
	apiRepository := NewRepositoryApi(apiKey)
	dbRepository := NewRepositoryDB(database)
	cdnRepository := NewRepositoryCDN(storagePath)

	updaterController := NewController(apiRepository, cdnRepository, dbRepository, cronPattern)
	return updaterController
}
