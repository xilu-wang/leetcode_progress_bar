package cmd

import (
	"github.com/spf13/cobra"
	"leetcode_progress_bar/internal/progress"
	"log"
)

var time string

var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "leetcode progress bar",
	Long:  "leetcode progress bar to show your submission record daily, monthly and yearly.",
	Run: func(cmd *cobra.Command, args []string) {
		switch time {
		case "day":
			progress.ShowProgressBar(progress.Day)
		case "month":
			progress.ShowProgressBar(progress.Month)
		case "year":
			progress.ShowProgressBar(progress.Year)
		default:
			log.Fatalf("invalid time interval")
		}
	},
}

func init() {
	progressCmd.Flags().StringVarP(&time, "time", "t", "", "please input time interval")
}