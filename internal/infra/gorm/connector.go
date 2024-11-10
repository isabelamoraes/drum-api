package gormpostgres

import (
	"drum-api/internal/adapter/logger"
	"drum-api/internal/error/errorcontext"
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	once  sync.Once
	dbErr error
)

func GetConnection() (*gorm.DB, error) {
	once.Do(func() {
		databaseConnection := "host=localhost user=admin password=123456 dbname=drum-api port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

		db, dbErr = gorm.Open(postgres.Open(databaseConnection))

		if dbErr != nil {
			errMess := fmt.Sprintf("Unable to connect to database: %v\n", dbErr)
			logger.LogInfo(logger.ErrorLog{Message: errMess, Context: errorcontext.PostgresAdapter})

			return
		}

		sqlDB, err := db.DB()
		if err != nil {
			dbErr = fmt.Errorf("unable to get sql.DB from gorm.DB: %v", err)
			return
		}

		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(3)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})

	return db, dbErr
}
