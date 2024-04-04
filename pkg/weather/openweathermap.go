package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherResponse struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
}

type OpenWeatherMap struct{}

func (owm *OpenWeatherMap) GetWeather(location string) (string, error) {
	url := fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", conf.OpenWeatherMapDomain, location, conf.OpenWeatherMapKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return "", err
	}

	if len(weatherResponse.Weather) > 0 {
		return weatherResponse.Weather[0].Main, nil
	}

	return string(body), nil
}
