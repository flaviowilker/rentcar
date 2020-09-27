package main

import (
	"github.com/flaviowilker/rentcar/infrastructure/database"
)

func main() {
	db := database.NewConfig()
	dbSQL, ok := db.DB()
	if ok != nil {
		defer dbSQL.Close()
	}

}
