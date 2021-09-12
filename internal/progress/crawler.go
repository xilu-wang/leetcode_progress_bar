package progress

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseURL = "https://leetcode.com/api/submissions"
	batchSize = 20
	acceptance = "Accepted"
)

type Submission struct {
	Timestamp int64	`json:"timestamp"`
	TitleSlug string `json:"title_slug"`
	StatusDisplay string `json:"status_display"`
}

var (
	httpClient *http.Client
)

func InitCrawler() {
	httpClient = &http.Client{}
}

func CrawlSubmission() (map[string]int64, error) {
	// validate cookie
	if len(cookie) == 0 {
		cookieLines := ReadFile(CookiePath)
		if len(cookieLines) == 0 {
			panic("ERROR: cookie not found, please set cookie")
		}

		cookie = cookieLines[0]
		if len(cookie) == 0 {
			panic("ERROR: invalid cookie, please reset cookie")
		}
	}

	hasNext := true
	offset := 0
	curTime := time.Now().Unix()

	header := map[string]string {
		"Cookie": cookie,
	}

	if len(dataMap) == 0 {
		dataMap = make(map[string]int64)
	}

	newData := make(map[string]int64)

	for hasNext && curTime > lastCrawlTimestamp {
		url := fmt.Sprintf("%v?offset=%d&limit=%d", baseURL, offset, batchSize)
		respStr, err := get(url, header)
		if err != nil {
			panic(err)
		}
		var submissions []Submission
		data, _, _, _ := jsonparser.Get([]byte(respStr), "submissions_dump")
		if err := json.Unmarshal(data, &submissions); err != nil {
			continue
		}

		for _, s := range submissions {
			if s.StatusDisplay != acceptance {
				continue
			}

			timestamp := dataMap[s.TitleSlug]
			if timestamp == 0 || timestamp > s.Timestamp {
				dataMap[s.TitleSlug] = s.Timestamp
				newData[s.TitleSlug] = s.Timestamp
			}

			if s.Timestamp < curTime {
				curTime = s.Timestamp
			}
		}

		hasNext, _ = jsonparser.GetBoolean([]byte(respStr), "has_next")
		offset += batchSize
	}

	// write to data file
	var newDataStr string
	for k, v := range newData {
		newDataStr += fmt.Sprintf("%v:%v\n", k, v)
	}
	AppendFile(data, newDataStr)

	return dataMap, nil
}

func get(url string, header map[string]string) (string, error){
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, err := httpClient.Do(req)
	if res == nil || err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}