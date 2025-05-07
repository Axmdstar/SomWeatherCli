package api_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"somweathercli/api"
)

func Test_WmoCode(t *testing.T) {
	t.Run("Read Wmo Code from json file", func(t *testing.T) {
		dataR, err := api.GetWmoCodefile()
		if err != nil {
			t.Error("Error Opening file")
		}

		// fmt.Println(string(dataR))
		wmocodeArray := api.NewWmoCode()
		err = json.Unmarshal(dataR, wmocodeArray)
		if err != nil {
			fmt.Println(err)
			t.Error("failed to umarshal json")
			return
		}

		if wmocodeArray == nil {
			t.Error("Wmo Code is Empty")
			return
		}

		for k, v := range wmocodeArray.WmoCodes {
			if k == "" {
				t.Error("Key is empty")
				return
			}

			if v.Day == (api.DayCycle{}) {
				t.Error("Day field is empty")
			}
			// fmt.Printf("Key: %v is OK \n", k)
		}
	})

	t.Run("one row table ", func(t *testing.T) {
		// Api Request dummy data
		dummy := `
		{
  "latitude": 6.0,
  "longitude": 48.0,
  "generationtime_ms": 0.0133514404296875,
  "utc_offset_seconds": 10800,
  "timezone": "Africa/Cairo",
  "timezone_abbreviation": "GMT+3",
  "elevation": 121.0,
  "current_units": {
    "time": "iso8601",
    "interval": "seconds",
    "weather_code": "wmo code"
  },
  "current": {
    "time": "2025-05-04T15:15",
    "interval": 900,
    "weather_code": 3
  }
}
 `
		dataR, err := api.GetWmoCodefile()
		if err != nil {
			t.Error("Error Opening file")
		}

		// Create struct for Current weather
		dummyCurrentWther := api.NEWCurrentWther()
		err = json.Unmarshal([]byte(dummy), dummyCurrentWther)
		if err != nil {
			t.Error(err)
		}
		// fmt.Printf("dummy Struct Current Weather: %v ", dummyCurrentWther)

		// Create Struct for Dummy Wmo Code
		dummyWmoCode := api.NewWmoCode()
		err = json.Unmarshal(dataR, dummyWmoCode)
		if err != nil {
			t.Error(err)
		}

		wcode := strconv.Itoa(dummyCurrentWther.Current.WeatherCode)
		emoji, description, err := api.GetWmoCodeData(wcode)
		if err != nil {
			t.Error(err)
		}

		want := "☁️"
		wantdes := "Cloudy"
		if emoji != want {
			t.Errorf("got %v Want %v ", emoji, want)
		}

		if description != wantdes {
			t.Errorf("got %v Want %v ", description, wantdes)
		}

		// date, clock :=	api.DateTimeStrings(dummyCurrentWther.Current.Time)
	})
}
