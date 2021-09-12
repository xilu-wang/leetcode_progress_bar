package progress

import (
	"fmt"
	"testing"
)

func init() {
	InitData()
	InitCrawler()
}

func TestCrawlSubmission(t *testing.T) {
	outputMap, _ := CrawlSubmission()
	fmt.Println(len(outputMap))
}
