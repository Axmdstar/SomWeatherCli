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

func DateTimeStrings(dateStr string) (string, string) {
	var date strings.Builder
	var clock strings.Builder

	parsedTime, err := time.Parse("2006-01-02T15:04", dateStr)
	if err != nil {
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

	return date.String(), clock.String()
}

func GetWmoCodeData(code string) {
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

	fmt.Println(wmocodeArray)
	return ""
}
