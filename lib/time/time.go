// Package internaltime handles time conversions
package internaltime

import (
	"time"

	"github.com/car-journal/api-backend/config"
)

func ConvertStringToTime(layout string, timeStr string) (time.Time, error) {
	loc, _ := time.LoadLocation(config.Get(config.DefaultTimeZone))
	res, err := time.ParseInLocation(layout, timeStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	return res, nil
}
