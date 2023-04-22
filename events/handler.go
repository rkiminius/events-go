package events

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitEventsRoutes(e *echo.Echo) {
	e.GET("/events", getEventsHandler)
	e.POST("/events", postEventsHandler)
}

func getEventsHandler(c echo.Context) error {
	eventsFromDb, err := getAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, eventsFromDb)
}

func postEventsHandler(c echo.Context) error {
	e := new(Event)
	if err := c.Bind(e); err != nil {
		return err
	}

	if err := validateEventInput(e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	eventFromDb, err := insert(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, eventFromDb)
}
