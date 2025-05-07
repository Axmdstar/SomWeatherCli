package api

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetWmoCodefile() ([]byte, error) {
	file, err := os.Open("./api/wmo_code.json")
	if err != nil {
		return nil, err
	}

	dataR, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dataR, nil
}

// returns date and clock
func DateTimeStrings(dateStr string) (string, string, error) {
	var date strings.Builder
	var clock strings.Builder

	parsedTime, err := time.Parse("2006-01-02T15:04", dateStr)
	if err != nil {
		return "", "", err
	}

	date.WriteString(strconv.Itoa(parsedTime.Year()))
	date.WriteString("/")
	date.WriteString(parsedTime.Month().String())
	date.WriteString("/")
	date.WriteString(strconv.Itoa(parsedTime.Day()))

	h, m, _ := parsedTime.Clock()
	clock.WriteString(strconv.Itoa(h))
	clock.WriteString(":")
	clock.WriteString(strconv.Itoa(m))

	return date.String(), clock.String(), nil
}

// Returns Emoji and description
func GetWmoCodeData(code string) (emoji, description string, Error error) {
	dataR, err := GetWmoCodefile()
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	WmoCode := NewWmoCode()
	err = json.Unmarshal(dataR, WmoCode)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	fmt.Println(WmoCode.WmoCodes[code])
	return WmoCode.WmoCodes[code].Day.Emoji, WmoCode.WmoCodes[code].Day.Description, nil
}

func CurrentWtherformatter(data *CurrentWther) string {
	// get wmo_code
	file, err := os.Open("./api/wmo_code.json")
	if err != nil {
		fmt.Println(err)
	}

	dataR, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(dataR))
	wmocodeArray := NewWmoCode()
	err = json.Unmarshal(dataR, wmocodeArray)
	if err != nil {
		fmt.Println(err)
	}

	date, clock, err := DateTimeStrings(data.Current.Time)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(date, clock)

	code := strconv.Itoa(data.Current.WeatherCode)
	emoji, description, err := GetWmoCodeData(code)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(emoji, description)
	return fmt.Sprintf("date %v, clock %v, emoji %v, description %v", date, clock, emoji, description)
}
