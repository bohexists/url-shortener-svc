package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func StartServer(e *echo.Echo, address string) {
	// Запуск сервера в отдельной горутине
	go func() {
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Канал для захвата системных сигналов
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Ожидание сигнала для завершения
	<-quit

	log.Println("Shutting down server...")

	// Контекст с таймаутом для корректного завершения работы
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down the server: %v", err)
	}

	log.Println("Server stopped gracefully.")
}
