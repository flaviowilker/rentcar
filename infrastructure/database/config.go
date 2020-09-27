package database

import (
	"log"
	"os"

	"github.com/flaviowilker/rentcar/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	host     string
	name     string
	port     string
	user     string
	password string
}

// NewConfig ...
func NewConfig() *gorm.DB {

	configDB := config{
		host:     os.Getenv("DATABASE_HOST"),
		port:     os.Getenv("DATABASE_PORT"),
		name:     os.Getenv("DATABASE_NAME"),
		user:     os.Getenv("DATABASE_USER"),
		password: os.Getenv("DATABASE_PASSWORD"),
	}

	dsn := "host=" + configDB.host + " port=" + configDB.port + " dbname=" + configDB.name + " user=" + configDB.user + " password=" + configDB.password + " sslmode=disable TimeZone=America/Fortaleza"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&domain.Role{})
	db.AutoMigrate(&domain.CarModel{})
	db.AutoMigrate(&domain.Car{})
	db.AutoMigrate(&domain.CarSchedule{})
	db.AutoMigrate(&domain.UserSchedule{})
	db.AutoMigrate(&domain.User{})

	return db
}
