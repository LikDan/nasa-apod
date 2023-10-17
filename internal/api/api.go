package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewApi(group *gin.RouterGroup, host string, storagePath string, database *pgxpool.Pool) Handler {
	dbRepository := NewRepositoryDB(database)
	cdnRepository := NewRepositoryCDN(storagePath)

	basePath := host + group.BasePath()
	controllerImpl := NewController(cdnRepository, dbRepository, basePath)
	handlers := NewHandler(group, controllerImpl)

	return handlers
}
