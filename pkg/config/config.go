package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AccWeatherDomain     string
	AccWeatherKey        string
	OpenWeatherMapDomain string
	OpenWeatherMapKey    string
}

func Load() *Config {
	_ = godotenv.Load()

	config := Config{
		AccWeatherDomain:     os.Getenv("ACC_WEATHER_DOMAIN"),
		AccWeatherKey:        os.Getenv("ACC_WEATHER_KEY"),
		OpenWeatherMapDomain: os.Getenv("OPEN_WEATHER_MAP_DOMAIN"),
		OpenWeatherMapKey:    os.Getenv("OPEN_WEATHER_MAP_KEY"),
	}

	return &config
}
