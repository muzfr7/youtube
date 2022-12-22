package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/muzfr7/echoblog/config"
)

func main() {
	// Read .env file and export variables for this process
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	// Populate cfg variable with exported environment variables
	var cfg config.EnvConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello-world", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world..")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)))

	// TODO: Gracefully shutdown the server here..
}
