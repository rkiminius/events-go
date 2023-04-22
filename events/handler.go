package events

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func InitEventsRoutes(e *echo.Echo) {
	route := "/events"
	e.GET(route, getEventsHandler)
	e.GET(route+"/:id", getEventByIdHandler)
	e.POST(route, postEventsHandler)
}

func getEventsHandler(c echo.Context) error {
	eventsFromDb, err := getAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, eventsFromDb)
}

func getEventByIdHandler(c echo.Context) error {
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	eventFromId, err := getById(objId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, eventFromId)
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
