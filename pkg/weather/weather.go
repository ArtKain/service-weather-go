package weather

import "service-weather/pkg/config"

var conf = config.Load()

const (
	openWeatherMap string = "open-weather-map"
	accuWeather    string = "accu-weather"
)

type Getter interface {
	GetWeather(location string) (string, error)
}

type DecoratorStruct struct {
	Getter
}

func Decorator(provider string) *DecoratorStruct {
	var wp Getter

	if provider == accuWeather {
		wp = &AccuWeather{}
	}

	if provider == openWeatherMap {
		wp = &OpenWeatherMap{}

	}

	return &DecoratorStruct{Getter: wp}
}

func (wd *DecoratorStruct) GetWeather(location string) (string, error) {
	baseResult, err := wd.Getter.GetWeather(location)
	if err != nil {
		return "", err
	}
	return baseResult, nil
}
