package main

import (
	"github.com/flaviowilker/rentcar/app/infrastructure/database"
)

func main() {
	db := database.NewConfig()
	dbSQL, err := db.DB()
	if err != nil {
		defer dbSQL.Close()
	}

}
