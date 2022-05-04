package controllers

import (
	"net/http"

	"github.com/devazizi/golang-ticket-booking.git/services"
	"github.com/labstack/echo/v4"
)

func IndexTickets(c echo.Context) error {

	type Empty struct {
		Data string `json:"message"`
	}

	return c.JSON(http.StatusOK, services.Response{Status: true, Content: Empty{Data: "hello world"}})
}
