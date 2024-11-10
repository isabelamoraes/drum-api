package main

import (
	"drum-api/internal/adapter/http/router"
	gormpostgres "drum-api/internal/infra/gorm"
	"os"
)

func main() {
	conn, err := gormpostgres.GetConnection()
	if err != nil {
		os.Exit(1)
	}

	r := router.Router(*conn)
	_ = r.Run()
}
