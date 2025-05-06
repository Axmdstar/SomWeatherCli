package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Api's Endpoint
// https://api.open-meteo.com/v1/forecast?

// Query Params
// latitude=6&
// longitude=48&
// daily=weather_code&
// hourly=weather_code&
// current=weather_code&
// timezone=Africa%2FCairo

var endpoint string = "https://api.open-meteo.com/v1/forecast?latitude=6&longitude=48&current=weather_code&timezone=Africa%2FCairo"

func GetCurrentWthr() (*CurrentWther, error) {
	res, err := http.Get(endpoint)
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

	data := NEWCurrentWther()
	err = json.Unmarshal(dataR, data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}
