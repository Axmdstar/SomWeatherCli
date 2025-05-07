package cmd

import (
	"fmt"

	"somweathercli/api"

	"github.com/spf13/cobra"
)

var CurrentCmd = &cobra.Command{
	Use:   "now",
	Short: "Current weather",
	Long:  `show current Weather in somali, returns in a nice and beauty format to understand it easly`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := api.GetCurrentWthr()
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		fmt.Println(api.CurrentWtherformatter(data))
	},
}

func init() {
	rootCmd.AddCommand(CurrentCmd)
}
