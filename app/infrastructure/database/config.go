package database

import (
	"errors"
	"log"
	"os"

	"github.com/flaviowilker/rentcar/app/infrastructure/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	messageErrorConnection = "Error connecting to database"
	errHostIsNull          = errors.New("Database Host is null")
	errPortIsNull          = errors.New("Database Port is null")
	errNameIsNull          = errors.New("Database Name is null")
	errUserIsNull          = errors.New("Database User is null")
	errPasswordIsNull      = errors.New("Database Password is null")
)

type config struct {
	host     string
	port     string
	name     string
	user     string
	password string
}

func (c config) configValidate() error {

	switch {
	case c.host == "":
		return errHostIsNull
	case c.port == "":
		return errPortIsNull
	case c.name == "":
		return errNameIsNull
	case c.user == "":
		return errUserIsNull
	case c.password == "":
		return errPasswordIsNull
	}

	return nil
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

	if err := configDB.configValidate(); err != nil {
		log.Fatalf("%v: %v", messageErrorConnection, err)
		panic(err)
	}

	db := createPostgresConnection(configDB)
	migrations.ExecuteMigrations(db)

	return db
}

func createPostgresConnection(configDB config) *gorm.DB {
	dsn := "host=" + configDB.host + " port=" + configDB.port + " dbname=" + configDB.name + " user=" + configDB.user + " password=" + configDB.password + " sslmode=disable TimeZone=America/Fortaleza"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("%v: %v", messageErrorConnection, err)
		panic(err)
	}

	return db
}
