package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Создаем новый Echo-инстанс
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Маршруты
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the URL Shortener Service!")
	})

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
