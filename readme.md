# SomWeather Cli tool

## Overview

This project is a command-line interface (CLI) tool written in Go that fetches and displays weather data from a public API.

## Features

- Fetch current weather data.
- Fetch weather forecast for the next 7 days.
- Display weather information in a user-friendly table format.

## Usage

to run the tool, use the following command:

```bash
go run main.go daily
or
./somwther daily
```

returns the daily weather forecast for the next 7 days.

```bash
go run main.go now
or
./somwther now
```

returns the current weather data.

## Implementation Details

- Create two packages/directories: `cmd` for CLI commands and `api` for API interactions and formatting tables.
- Use the `cobra` package for CLI command handling.
- Use the `net/http` package to make HTTP requests to the weather API.
- Use the `encoding/json` package to parse JSON responses.
- Format and display the weather data in a readable table format using the `text/tabwriter`

## References

Api Link: `https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&daily=weather_code&timezone=Africa%2FCairo`

Api Doc: `https://open-meteo.com/en/docs?timezone=Africa%2FCairo&daily=weather_code`

Cli flags Cobra: `https://github.com/spf13/cobra`
