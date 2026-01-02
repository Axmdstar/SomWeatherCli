// Package cmd defines the command-line interface (CLI) commands and their behavior using the Cobra library and group our commands.
package cmd

import (
	"fmt"

	"somweathercli/api"

	"github.com/spf13/cobra"
)

// CurrentCmd represents the current command
// show current Weather in somali, returns in a nice and beauty format to understand it easly
// Using cobra.Command struct to define the command
// Importing and using api package to get and format the weather data
var CurrentCmd = &cobra.Command{
	Use:   "now",
	Short: "Current weather",
	Long:  `show current Weather in somali, returns in a nice and beauty format to understand it easly`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get current weather data
		data, err := api.GetCurrentWthr()
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		// pass data to formatter
		date, clock, description, emoji := api.CurrentWtherformatter(data)
		// prepare data for table
		dataArray := [][]string{
			{date, clock, description, emoji},
		}
		// write table
		api.WriteTableCurrent(dataArray)
	},
}

// Initialize the CurrentCmd and add it to the root command
func init() {
	rootCmd.AddCommand(CurrentCmd)
}
