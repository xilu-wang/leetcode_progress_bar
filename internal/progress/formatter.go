package progress

import (
	"fmt"
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
		fmt.Printf("ERROR: %v", err)
	}

	switch interval {
	case Day:
		fmt.Println(buildDailySummary(dataMap))
	case Month:
		fmt.Println(buildMonthlySummary(dataMap))
	case Year:
		fmt.Println(buildYearlySummary(dataMap))
	}
}

func buildDailySummary(dataMap map[string]int64) string {
	dailyMap := make(map[string]int)

	for _, v := range dataMap {
		timeStrFull := time.Unix(v, 0).Format("2006-01-02 15:04:05")
		timeStr := timeStrFull[:10]
		dailyMap[timeStr] = dailyMap[timeStr] + 1
	}

	var output string
	for k, v := range dailyMap {
		output += fmt.Sprintf("%v: %v questions \n", k, v)
	}

	return output
}

func buildMonthlySummary(dataMap map[string]int64) string {
	monthlyMap := make(map[string]int)

	for _, v := range dataMap {
		timeStrFull := time.Unix(v, 0).Format("2006-01-02 15:04:05")
		timeStr := timeStrFull[:7]
		monthlyMap[timeStr] = monthlyMap[timeStr] + 1
	}

	var output string
	for k, v := range monthlyMap {
		output += fmt.Sprintf("%v: %v questions \n", k, v)
	}

	return output
}

func buildYearlySummary(dataMap map[string]int64) string {
	yearlyMap := make(map[string]int)

	for _, v := range dataMap {
		timeStrFull := time.Unix(v, 0).Format("2006-01-02 15:04:051")
		timeStr := timeStrFull[:4]
		yearlyMap[timeStr] = yearlyMap[timeStr] + 1
	}

	var output string
	for k, v := range yearlyMap {
		output += fmt.Sprintf("%v: %v questions \n", k, v)
	}

	return output
}
