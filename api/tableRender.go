package api

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Table rendering functionality
// Using tablewriter package to create and render tables in the console

func WriteTableCurrent(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Time", "Description", "Icon"})
	table.AppendBulk(data)
	table.Render()
}

func WriteTableDaily(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Description", "Icon"})
	table.AppendBulk(data)
	table.Render()
}
