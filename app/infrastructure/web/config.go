package web

import (
	"errors"
	"log"
	"os"

	"github.com/flaviowilker/rentcar/app/infrastructure/web/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

var (
	errPortIsNull = errors.New("Error web port is null")
)

// CreateConfig ...
func CreateConfig(db *gorm.DB) {

	configWeb := &config{os.Getenv("WEB_PORT")}
	if err := configWeb.configValidate(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	router.CreateAppRouter(db, app)

	if err := app.Listen(":" + configWeb.port); err != nil {
		log.Fatal(err)
		panic(err)
	}
}

type config struct {
	port string
}

func (c *config) configValidate() error {
	if c.port == "" {
		return errPortIsNull
	}

	return nil
}
