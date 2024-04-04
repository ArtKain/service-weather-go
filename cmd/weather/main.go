// Package main provides ...
package main

import (
	"fmt"
	"os"
	"service-weather/pkg/weather"
	"strings"
)

func main() {

	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("not suitable for passing arguments")
		os.Exit(1)
	}

	provider := os.Args[1]
	location := os.Args[2]
	var sendTo []string

	if len(os.Args) == 4 {
		sendTo = strings.Split(os.Args[3], ":")
	}

	decorated := weather.Decorator(provider)
	response, err := decorated.GetWeather(location)

	if err != nil {
		return
	}

	if len(sendTo) > 0 {
		sendToProvider := sendTo[0]
		switch sendToProvider {
		//
		}
	} else {
		fmt.Println(response)
	}
}
