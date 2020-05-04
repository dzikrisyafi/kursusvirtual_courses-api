package date_utils

import "time"

const (
	apiDBLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
