package cmd

import (
	"fmt"

	"somweathercli/api"

	"github.com/spf13/cobra"
)

var DailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "daily forecast",
	Long:  `shows 7 days forecast of Weather in somali, returns in a nice and beauty format to understand it easly`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := api.GetDailyWthr()
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		api.WriteTableDaily(api.DailyWeatherFormatter(data))
	},
}

func init() {
	rootCmd.AddCommand(DailyCmd)
}
