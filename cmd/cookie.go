package cmd

import (
	"github.com/spf13/cobra"
	"leetcode_progress_bar/internal/progress"
	"leetcode_progress_bar/internal/util"
)

var cookie string

var cookieCmd = &cobra.Command{
	Use:   "cookie",
	Short: "Set your cookie to crawl your leetcode",
	Long:  "Set cookie from your browser to crawl leetcode website.",
	Run: func(cmd *cobra.Command, args []string) {
		progress.OverwriteFile(progress.CookiePath, cookie)
		util.PrintSuccess("Cookie set successfully!")
	},
}

func init() {
	cookieCmd.Flags().StringVarP(&cookie, "set", "s", "", "please copy and paste your leetcode cookie.")
}
