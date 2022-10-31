package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "myapp/docs/echosimple"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host rest-api-design-research.herokuapp.com:8080
// @BasePath /
// @schemes http

func main() {
	// Echo instance
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Weather APIs
	e.POST("/weather", GetWeatherData)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}

type QueryByDateRequest struct {
	Date string `json:"date" validate:"required"`
}

type WeatherObj struct {
	Temperature float32 `json:"temperature"`
	WindSpeed   float32 `json:"wind_speed"`
}

func GetWeatherData(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return err
	}
	fmt.Println(q)
	WeatherObj := &WeatherObj{
		Temperature: 72.01,
		WindSpeed:   14.11,
	}
	return c.JSON(http.StatusOK, WeatherObj)
}

// func GetWeatherDataType2(c echo.Context) error {
// 	q := new(QueryByDateRequest)
// 	if err := c.Bind(q); err != nil {
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}
// 	// Checking if required fields are present
// 	if err := c.Validate(q); err != nil {
// 		return err
// 	}
// 	fmt.Println(q)
// 	WeatherObj := &WeatherObj{
// 		Temperature: 72.01,
// 		WindSpeed:   14.11,
// 	}
// 	return c.JSON(http.StatusOK, WeatherObj)
// }

// func GetWeatherDataType3(c echo.Context) error {
// 	q := new(QueryByDateRequest)
// 	if err := c.Bind(q); err != nil {
// 		return c.String(http.StatusBadRequest, err.Error())
// 	}
// 	// Checking if required fields are present
// 	if err := c.Validate(q); err != nil {
// 		return err
// 	}
// 	fmt.Println(q)
// 	WeatherObj := &WeatherObj{
// 		Temperature: 72.01,
// 		WindSpeed:   14.11,
// 	}
// 	return c.JSON(http.StatusOK, WeatherObj)
// }
