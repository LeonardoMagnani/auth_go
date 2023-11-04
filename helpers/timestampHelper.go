package helpers

import (
	"time"
)

func Timestamp(timestampString string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return time.Now(), err
	}

	timestamp, err := time.ParseInLocation(layout, timestampString, location)
	if err != nil {
		return time.Now(), err
	}

	return timestamp, nil
}
