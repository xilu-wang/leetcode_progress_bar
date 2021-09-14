package progress

import (
	"fmt"
	"sort"
	"time"
)

const (
	Day TimeInterval = 0
	Month TimeInterval = 1
	Year TimeInterval = 2
)

func ShowProgressBar(interval TimeInterval) {

	dataMap, err := CrawlSubmission()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	switch interval {
	case Day:
		fmt.Println(buildSummary(dataMap, 10))
	case Month:
		fmt.Println(buildSummary(dataMap, 7))
	case Year:
		fmt.Println(buildSummary(dataMap, 4))
	}
}

func buildSummary(dataMap map[string]int64, end int) string {
	dailyMap := make(map[string]int)
	var keys []string

	for _, v := range dataMap {
		timeStrFull := time.Unix(v, 0).Format("2006-01-02 15:04:05")
		timeStr := timeStrFull[:end]
		dailyMap[timeStr] = dailyMap[timeStr] + 1
	}

	for k, _ := range dailyMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	output := "\n"
	for _, key := range keys {
		output += fmt.Sprintf("%v: %v questions \n", key, dailyMap[key])
	}

	return output
}