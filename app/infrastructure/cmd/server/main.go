package main

import (
	"log"

	"github.com/flaviowilker/rentcar/app/infrastructure/database"
	"github.com/flaviowilker/rentcar/app/infrastructure/web"
)

func main() {
	db := database.NewConfig()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(1)
	defer sqlDB.Close()

	web.CreateConfig(db)
}
