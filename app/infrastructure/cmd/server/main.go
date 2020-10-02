package main

import (
	"github.com/flaviowilker/rentcar/app/infrastructure/database"
	"github.com/flaviowilker/rentcar/app/infrastructure/web"
)

func main() {
	db := database.NewConfig()
	dbSQL, err := db.DB()
	if err != nil {
		defer dbSQL.Close()
	}

	web.CreateConfig(db)
}
