package progress

import (
	"fmt"
	"leetcode_progress_bar/internal/util"
	"testing"
)

func init() {
	InitData()
	InitCrawler()
}

func TestCrawlSubmission(t *testing.T) {
	outputMap, err := CrawlSubmission()
	if err != nil {
		util.PrintError(err.Error())
	}
	fmt.Println(len(outputMap))
}
