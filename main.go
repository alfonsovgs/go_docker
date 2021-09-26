package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Message struct {
	value string `json:"value"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello World")
	})

	e.POST("/send", func(c echo.Context) error {
		return sendHandler(c)
	})

	e.GET("/messages", func(c echo.Context) error {
		r, err := getMessagesHandler(c)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, r)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func initStore() (*sql.DB, error) {
	return nil, nil
}

func sendHandler(c echo.Context) error {
	return nil
}

func getMessagesHandler(c echo.Context) (*Message, error) {
	return nil, nil
}
