package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func StringToTime(date string) (time.Time, error) {
	now, err := time.Parse(apiDateLayout, date)
	if err != nil {
		return time.Now(), err
	}
	return now, nil

}
