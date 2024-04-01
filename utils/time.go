package utils

import "time"

var TimeFormat = "2006-01-02 15:04:05"

func NowStr() string {
	return time.Now().Format(TimeFormat)
}
