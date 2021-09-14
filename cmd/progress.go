package cmd

import (
	"github.com/spf13/cobra"
	"leetcode_progress_bar/internal/progress"
	"leetcode_progress_bar/internal/util"
)

var time string

var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Show your leetcode progress",
	Long:  "A leetcode progress report to show your submissions daily, monthly or yearly.",
	Run: func(cmd *cobra.Command, args []string) {
		switch time {
		case "day":
			progress.ShowProgressBar(progress.Day)
		case "month":
			progress.ShowProgressBar(progress.Month)
		case "year":
			progress.ShowProgressBar(progress.Year)
		default:
			util.PrintError("Invalid time interval. Please input 'day', 'month', or 'year'.")
		}
	},
}

func init() {
	progressCmd.Flags().StringVarP(&time, "time", "t", "", "Please input time interval: 'day', 'month', or 'year'.")
}