package progress

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type TimeInterval int

const (
	BaseDir = "/tmp/leetcode_progress"
	LogPath = "/tmp/leetcode_progress/log"
	DataPath = "/tmp/leetcode_progress/data"
	CookiePath = "/tmp/leetcode_progress/cookie"
)

var (
	dataMap map[string]int64
	cookie string
)

func InitData() {
	// create leetcode_progress dir under /tmp if not exist
	validateBaseDir(DataPath)

	// check cookie
	if fileExist(cookie) {
		cookieLines := ReadFile(CookiePath)
		if len(cookieLines) == 0 {
			fmt.Println("WARNING: cookie not found, please set cookie")
		} else if len(cookieLines[0]) == 0 {
			fmt.Println("WARNING: invalid cookie, please reset cookie")
		} else {
			cookie = cookieLines[0]
		}
	}

	// init dataMap, key is the title-slug, value is the first accepted timestamp
	dataMap = make(map[string]int64)
	if fileExist(DataPath) {
		buildDataMap()
	}
}

func OverwriteFile(path string, content string) {
	var f *os.File
	var err error
	if f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755); err != nil {
		if f, err = os.Create(path); err != nil {
			panic(err)
		}
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}(f)

	w := bufio.NewWriter(f)
	if _, err := w.WriteString(content); err != nil {
		panic(err)
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}

	return
}

func AppendFile(path string, content string) {
	var f *os.File
	var err error
	if f, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend); err != nil {
		if f, err = os.Create(path); err != nil {
			panic(err)
		}
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}(f)

	w := bufio.NewWriter(f)
	if _, err := w.WriteString(content); err != nil {
		panic(err)
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}

	return
}

func ReadFile(path string) []string {
	var f *os.File
	var err error
	if f, err = os.Open(path); err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}(f)

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func GetLog() int64 {
	var lastLogTimestamp int64 = 0

	if fileExist(LogPath) {
		lines := ReadFile(LogPath)
		if len(lines) == 0 {
			return lastLogTimestamp
		}

		lastLogTimestamp, _ = strconv.ParseInt(lines[0], 10, 64)
	}

	return lastLogTimestamp
}

func UpdateLog(time int64) {
	if fileExist(LogPath) {
		OverwriteFile(LogPath, fmt.Sprintf("%v", time))
	}
}

func RemoveFile(path string) {
	if fileExist(path) {
		if e := os.Remove(path); e != nil {
			panic(e)
		}
	}
}

func RemoveDir() {
	info, err := os.Stat(BaseDir)
	if err == nil && info.IsDir() {
		if e := os.RemoveAll(BaseDir); e != nil {
			panic(e)
		}
	}
}

func buildDataMap() {
	var f *os.File
	var err error
	if f, err = os.Open(DataPath); err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		if len(line) != 2 {
			continue
		}

		titleSlug := line[0]
		timestamp, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			continue
		}

		existTime := dataMap[titleSlug]
		if existTime == 0 || timestamp < existTime {
			// always pick the earliest time
			dataMap[titleSlug] = timestamp
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func validateBaseDir(p string) {
	baseDir := path.Dir(p)
	info, err := os.Stat(baseDir)
	if err != nil || !info.IsDir() {
		if dirErr := os.MkdirAll(baseDir, 0755); dirErr != nil {
			panic(dirErr)
		}
	}
}

func fileExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
