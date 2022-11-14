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

type User struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

type Users struct {
	Users []User `json:"users"`
}

var users = []User{
	{
		Email: "iamsharduld@gmail.com",
		Token: "2347FD2F854ECC36E6BD335DDD88F",
	},
	{
		Email: "mcoblenz@ucsd.edu",
		Token: "82C89F7F83E5BFA34297FC9E59985",
	},
}

func getUserFromToken(token string) (User, error) {
	for _, user := range users {
		if user.Token == token {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("invalid token")
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

	// Weather
	e.POST("/weather1", GetWeatherData1)
	e.POST("/weather2", GetWeatherData2)
	e.POST("/weather3", GetWeatherData3)

	// Stock Data (SPY)
	e.POST("/stockPrice1", GetStockData1)
	e.POST("/stockPrice2", GetStockData2)
	e.POST("/stockPrice3", GetStockData3)

	// Heart Rate Data
	e.POST("/heartRate1", GetHeartRateData1)
	e.POST("/heartRate2", GetHeartRateData2)
	e.POST("/heartRate3", GetHeartRateData3)

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
	Date  string `json:"date" validate:"required"`
	Token string `json:"token" validate:"required"`
}

type WeatherObj struct {
	Temperature float32 `json:"temperature"`
	WindSpeed   float32 `json:"wind_speed"`
}

func GetWeatherData1(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return err
	}
	user, err := getUserFromToken(q.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(user)

	WeatherObj := &WeatherObj{
		Temperature: 72.01,
		WindSpeed:   14.11,
	}
	return c.JSON(http.StatusOK, WeatherObj)
}

func GetWeatherData2(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user, err := getUserFromToken(q.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(user)

	WeatherObj := &WeatherObj{
		Temperature: 72.01,
		WindSpeed:   14.11,
	}
	return c.JSON(http.StatusOK, WeatherObj)
}

func GetWeatherData3(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return c.String(http.StatusOK, "Invalid request")
	}
	user, err := getUserFromToken(q.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(user)

	WeatherObj := &WeatherObj{
		Temperature: 72.01,
		WindSpeed:   14.11,
	}
	return c.JSON(http.StatusOK, WeatherObj)
}

type StockObj struct {
	Price float32 `json:"price"`
}

func GetStockData1(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return err
	}
	user, err := getUserFromToken(q.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(user)
	StockObj := &StockObj{
		Price: 340.23,
	}
	return c.JSON(http.StatusOK, StockObj)
}

func GetStockData2(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user, err := getUserFromToken(q.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(user)
	StockObj := &StockObj{
		Price: 340.23,
	}
	return c.JSON(http.StatusOK, StockObj)
}

func GetStockData3(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return c.String(http.StatusOK, "Invalid request")
	}
	user, err := getUserFromToken(q.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(user)
	StockObj := &StockObj{
		Price: 340.23,
	}
	return c.JSON(http.StatusOK, StockObj)
}

type HeartRate struct {
	HeartRate float32 `json:"heart_rate"`
}

func GetHeartRateData1(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return err
	}
	user, err := getUserFromToken(q.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(user)
	HeartRateObj := &HeartRate{
		HeartRate: 75,
	}
	return c.JSON(http.StatusOK, HeartRateObj)
}

func GetHeartRateData2(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user, err := getUserFromToken(q.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(user)
	HeartRateObj := &HeartRate{
		HeartRate: 75,
	}
	return c.JSON(http.StatusOK, HeartRateObj)
}

func GetHeartRateData3(c echo.Context) error {
	q := new(QueryByDateRequest)
	if err := c.Bind(q); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// Checking if required fields are present
	if err := c.Validate(q); err != nil {
		return c.String(http.StatusOK, "Invalid request")
	}
	user, err := getUserFromToken(q.Token)
	if err != nil {
		return fmt.Errorf("invalid token")
	}
	fmt.Println(user)
	HeartRateObj := &HeartRate{
		HeartRate: 75,
	}
	return c.JSON(http.StatusOK, HeartRateObj)
}
