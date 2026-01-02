// Package cmd defines the command-line interface (CLI) commands and their behavior using the Cobra library and group our commands.
package cmd

// Packages are used to organize Go source code for better reusability and readability.
// Packages are a collection of Go sources files that reside in the same directory.
// Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "somweathercli",
	Short: "Somali Weather CLI Application",
	Long:  `Somweathercli is a command-line application that provides Weather information in Somali.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.somweathercli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
