package main

import (
	"events/db"
	"events/env"
	"events/events"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	env.Load()

	db.MongoConnect()
	defer db.MongoDisconnect()

	e := echo.New()
	e.Use(middleware.Logger())
	events.InitEventsRoutes(e)

	runServer(e)
}

func runServer(e *echo.Echo) {
	serverPort := env.Env["SERVER_PORT"]
	if serverPort == "" {
		serverPort = "8000"
	}
	fmt.Printf("starting server on: %s\n", serverPort)

	s := http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: e,
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
