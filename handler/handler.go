package handler

import (
	"net/http"

	"github.com/labstack/echo"

	s "myapp/service"
)

func GetWeatherData(c echo.Context) error {
	data, err := s.GetWeatherData("31/10/2022")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}
