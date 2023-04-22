package events

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func InitEventsRoutes(e *echo.Echo) {
	e.POST("/events", postEventsHandler)
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
		return err
	}

	return c.JSON(http.StatusCreated, eventFromDb)
}
