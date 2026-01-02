// Package cmd defines the command-line interface (CLI) commands and their behavior using the Cobra library and group our commands.
package cmd

import (
	"fmt"

	"somweathercli/api"

	"github.com/spf13/cobra"
)

// DailyCmd represents the daily command
// shows 7 days forecast of Weather in somali, returns in a nice and beauty format to understand it easly
// Using cobra.Command struct to define the command
// Importing and using api package to get and format the weather data
var DailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "daily forecast",
	Long:  `shows 7 days forecast of Weather in somali, returns in a nice and beauty format to understand it easly`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get daily weather data
		data, err := api.GetDailyWthr()
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		// pass data to formatter and write table
		api.WriteTableDaily(api.DailyWeatherFormatter(data))
	},
}

// Initialize the DailyCmd and add it to the root command
func init() {
	rootCmd.AddCommand(DailyCmd)
}
