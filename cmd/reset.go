package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"leetcode_progress_bar/internal/progress"
	"leetcode_progress_bar/internal/util"
)

var fileType string

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset and cleanup",
	Long:  "Reset and cleanup the data, log or cookie",
	Run: func(cmd *cobra.Command, args []string) {
		switch fileType {
		case "all":
			progress.RemoveDir()
		case "data":
			progress.RemoveFile(progress.DataPath)
			progress.RemoveFile(progress.LogPath)
		case "log":
			progress.RemoveFile(progress.LogPath)
		case "cookie":
			progress.RemoveFile(progress.CookiePath)
		default:
			util.PrintError("Invalid file type. Please input file type: 'data', 'log', 'cookie', or 'all'.")
			return
		}

		util.PrintSuccess(fmt.Sprintf("Reset and cleanup %v", fileType))
	},
}

func init() {
	resetCmd.Flags().StringVarP(&fileType, "file", "f", "", "Please input file type: 'data', 'log', 'cookie', or 'all'.")
}
