package utils

import "time"

func ParseDatetime(datetimeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04" // The layout string for the datetime format
	datetime, err := time.Parse(layout, datetimeStr)
	if err != nil {
		return time.Time{}, err
	}
	return datetime, nil
}
