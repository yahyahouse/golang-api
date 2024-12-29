package provider

import (
	"exam/config"
	"exam/internal/delivery/database"
	"exam/internal/repository"
)

var (
	dbConnectionPostgres *database.DB
	repositoryPostgres   repository.DatabaseRepository
)

func DBConnectionPostgres(database *database.DB) {
	if database != nil {
		dbConnectionPostgres = database
	}
}

func getDBConnectionPostgres() *database.DB {
	if dbConnectionPostgres != nil {
		return dbConnectionPostgres
	}

	dbConfig := getConfig().DBConfig
	dbConnection := database.InitDB(database.TypePostgres, dbConfig)
	return dbConnection
}

func getRepositoryPostgres() repository.DatabaseRepository {
	if repositoryPostgres == nil {
		cfg := config.GetConfig().DBConfig
		repositoryPostgres = repository.NewDatabaseRepository(getDBConnectionPostgres(), &cfg)
	}

	return repositoryPostgres
}
