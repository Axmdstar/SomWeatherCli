package api

// {
//   "latitude": 6.0,
//   "longitude": 48.0,
//   "generationtime_ms": 0.0133514404296875,
//   "utc_offset_seconds": 10800,
//   "timezone": "Africa/Cairo",
//   "timezone_abbreviation": "GMT+3",
//   "elevation": 121.0,
//   "current_units": {
//     "time": "iso8601",
//     "interval": "seconds",
//     "weather_code": "wmo code"
//   },
//   "current": {
//     "time": "2025-05-04T15:15",
//     "interval": 900,
//     "weather_code": 3
//   }
// }

type currentunits struct {
	Time        string `json:"time"`
	Interval    string `json:"interval"`
	WeatherCode string `json:"weather_code"`
}

type current struct {
	Time        string `json:"time"`
	Interval    int    `json:"interval"`
	WeatherCode int    `json:"weather_code"`
}

type CurrentWther struct {
	Latitude             float32      `json:"latitude"`
	Longitude            float32      `json:"longitude"`
	GenerationtimeMs     float64      `json:"generationtime_ms"`
	UtcOffsetSeconds     int          `json:"utc_offset_seconds"`
	TimeZone             string       `json:"timezone"`
	TimeZoneAbbreviation string       `json:"timezone_abbreviation"`
	Elevation            float64      `json:"elevation"`
	CurrentUnits         currentunits `json:"current_units"`
	Current              current      `json:"current"`
}

func NEWCurrentWther() *CurrentWther {
	return &CurrentWther{}
}

// "0":{
// 		"day":{
// 			"description":"Sunny",
// 			"emoji":"☀️"
// 		},
// 		"night":{
// 			"description":"Clear",
// 			"emoji":"🌕"
// 		}
// 	}

type DayCycle struct {
	Description string `json:"description"`
	Emoji       string `json:"emoji"`
}

type WmoCode struct {
	Day   DayCycle `json:"day"`
	Night DayCycle `json:"night"`
}

type WmoCodeArray struct {
	WmoCodes map[string]WmoCode `json:"wmo_codes"`
}

func NewWmoCode() *WmoCodeArray {
	return &WmoCodeArray{}
}

// {
//   "latitude": 5.125,
//   "longitude": 46.25,
//   "generationtime_ms": 0.019073486328125,
//   "utc_offset_seconds": 10800,
//   "timezone": "Africa/Cairo",
//   "timezone_abbreviation": "GMT+3",
//   "elevation": 241.0,
//   "daily_units": {
// "time": "iso8601",
// "weather_code": "wmo code"
//   },
//   "daily": {
//     "time": [
//       "2025-05-08",
//       "2025-05-09",
//       "2025-05-10",
//       "2025-05-11",
//       "2025-05-12",
//       "2025-05-13",
//       "2025-05-14"
//     ],
//     "weather_code": [
//       96,
//       3,
//       3,
//       3,
//       2,
//       2,
//       2
//     ]
//   }
// }

type dailyUnits struct {
	Time        string `json:"time"`
	WeatherCode string `json:"weather_code"`
}

type daily struct {
	Time        []string `json:"time"`
	WeatherCode []int    `json:"weather_code"`
}

type DailyWeather struct {
	Latitude             float32    `json:"latitude"`
	Longitude            float32    `json:"longitude"`
	GenerationtimeMs     float64    `json:"generationtime_ms"`
	UtcOffsetSeconds     int        `json:"utc_offset_seconds"`
	TimeZone             string     `json:"timezone"`
	TimeZoneAbbreviation string     `json:"timezone_abbreviation"`
	Elevation            float64    `json:"elevation"`
	DailyUnits           dailyUnits `json:"daily_units"`
	Daily                daily      `json:"daily"`
}

func NewDailyWeather() *DailyWeather {
	return &DailyWeather{}
}
