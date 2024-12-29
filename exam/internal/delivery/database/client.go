package database

import (
	"exam/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
)

const (
	TypePostgres = "postgres"
)

type DB struct {
	DBCon *gorm.DB
}

func InitDB(dbType string, dbConfig config.DBConfig) *DB {
	var conString string
	var database *gorm.DB
	var err error

	if dbType == TypePostgres {
		// postgresql://user:password@host:port/dbName?sslmode=disable\
		conString = "postgresql://" + url.QueryEscape(dbConfig.User) + ":" + url.QueryEscape(dbConfig.Password) + dbConfig.ConString
		//DB conString hard code
		//	conString = "host=localhost user=postgres password=postgres dbname=postgres port=5434 sslmode=disable TimeZone=Asia/Shanghai"
		database, err = gorm.Open(postgres.Open(conString), &gorm.Config{})
		//database, err := gorm.Open(postgres.Open(urlDB), &gorm.Config{})
	}
	if err != nil {
		panic("connection string: " + conString + "  " + err.Error())
	}

	// Set up connection pool
	sqlDB, _ := database.DB()
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenCon)
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleCon)
	sqlDB.SetConnMaxIdleTime(dbConfig.ConnIdleTimeoutDuration)

	db := &DB{
		DBCon: database,
	}

	return db
}
func (s *DB) GetDBCon() *gorm.DB {
	return s.DBCon
}
