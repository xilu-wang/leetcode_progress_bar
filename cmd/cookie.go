package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"leetcode_progress_bar/internal/progress"
)

var cookie string

var cookieCmd = &cobra.Command{
	Use:   "cookie",
	Short: "set cookie for leetcode",
	Long:  "set cookie from your leetcode website.",
	Run: func(cmd *cobra.Command, args []string) {
		progress.OverwriteFile(progress.CookiePath, cookie)
		fmt.Println("cookie set successfully!")
	},
}

func init() {
	cookieCmd.Flags().StringVarP(&cookie, "set", "s", "", "please copy and paste your leetcode cookie.")
}
