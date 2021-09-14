package main

import (
	"leetcode_progress_bar/cmd"
	"leetcode_progress_bar/internal/progress"
	"log"
)

func main() {

	progress.InitData()

	progress.InitCrawler()

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
