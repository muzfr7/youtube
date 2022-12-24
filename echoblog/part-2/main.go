package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello-world", func(c echo.Context) error {
		// return c.String(http.StatusOK, "Hello world..")

		res := Response{
			Message: "Hello world..",
		}

		return c.JSON(http.StatusOK, res)

	})

	e.Start("127.0.0.1:8080")
}
