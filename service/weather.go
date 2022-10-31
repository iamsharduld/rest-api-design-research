package service

type WeatherObj struct {
	temperature float32
	windSpeed   float32
}

func GetWeatherData(dayOfWeek string) (WeatherObj, error) {
	// Get the weather data
	WeatherObj := &WeatherObj{
		temperature: 72.0,
		windSpeed:   14.11,
	}
	return *WeatherObj, nil
}
