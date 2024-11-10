package main

import (
	gormpostgres "drum-api/internal/infra/gorm"
	"drum-api/internal/infra/gorm/model"
	"fmt"
)

func main() {
	db, err := gormpostgres.GetConnection()

	if err != nil {
		panic(fmt.Sprintf("Error connecting to db: %v", err))
	}

	err = db.Debug().
		AutoMigrate(
			&model.Drum{},
			&model.Review{},
		)

	if err != nil {
		panic(fmt.Sprintf("Error migrating: %v", err))
	}
}
