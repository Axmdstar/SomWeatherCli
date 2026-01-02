// Package api handles communication with the Open-Meteo API to fetch weather data.
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// NOTE: Api's Endpoint
// https://api.open-meteo.com/v1/forecast?

// NOTE:
// Query Params
// latitude=6&
// longitude=48&
// daily=weather_code&
// hourly=weather_code&
// current=weather_code&
// timezone=Africa%2FCairo

// NOTE: Open-Meteo API Endpoints
var (
	CurrentWthrendpoint string = "https://api.open-meteo.com/v1/forecast?latitude=6&longitude=48&current=weather_code&timezone=Africa%2FCairo"
	DailyWthrendpoint   string = "https://api.open-meteo.com/v1/forecast?latitude=5.1521&longitude=46.1996&daily=weather_code&timezone=Africa%2FCairo"
)

// GetCurrentWthr fetching Current Weather data from Open-Meteo API
func GetCurrentWthr() (*CurrentWther, error) {
	res, err := http.Get(CurrentWthrendpoint)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Ensure the response body is closed after reading
	defer res.Body.Close()

	// Read the response body
	dataR, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Create a new instance of CurrentWther to hold the unmarshaled data
	data := NEWCurrentWther()
	err = json.Unmarshal(dataR, data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

// GetDailyWthr fetching Daily Weather data from Open-Meteo API
func GetDailyWthr() (*DailyWeather, error) {
	res, err := http.Get(DailyWthrendpoint)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	dataR, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data := NewDailyWeather()
	err = json.Unmarshal(dataR, data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}
