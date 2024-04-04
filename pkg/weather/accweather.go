package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccuWeather struct{}

type LocationKeyResponse []struct {
	Key string `json:"Key"`
}

type AccWeatherResponse struct {
	Headline struct {
		Category string `json:"Category"`
	} `json:"Headline"`
}

func (aw *AccuWeather) GetWeather(location string) (string, error) {
	locationKey, _ := getLocationKey(location)

	url := fmt.Sprintf("%s/forecasts/v1/daily/1day/%s?apikey=%s", conf.AccWeatherDomain, locationKey, conf.AccWeatherKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var weatherResponse AccWeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return "", err
	}

	if len(weatherResponse.Headline.Category) > 0 {
		return weatherResponse.Headline.Category, nil
	}

	return "", err
}

// приватная функция для получения ключа города
func getLocationKey(location string) (string, error) {
	url := fmt.Sprintf("%s/locations/v1/cities/search?q=%s&apikey=%s", conf.AccWeatherDomain, location, conf.AccWeatherKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var locationKeyResponse LocationKeyResponse
	err = json.Unmarshal(body, &locationKeyResponse)
	if err != nil {
		return "", err
	}

	if len(locationKeyResponse[0].Key) > 0 {
		return locationKeyResponse[0].Key, nil
	}

	return "", err
}
