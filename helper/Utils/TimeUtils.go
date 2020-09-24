package Utils

import (
	"time"
)

func GetThisWeekStartEnd()(time.Time,time.Time){
	now := time.Now()
	mondayOffset := int(time.Monday - now.Weekday())
	if mondayOffset > 0 {
		mondayOffset = -6
	}
	sundayOffset := int(time.Saturday - now.Weekday())
	sundayOffset += 2
	if sundayOffset < 0{
		sundayOffset = 6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(),0,0,0,0,time.Local).AddDate(0,0,mondayOffset)
	weekEnd   := time.Date(now.Year(), now.Month(), now.Day(),0,0,0,0,time.Local).AddDate(0,0,sundayOffset)
	return weekStart, weekEnd
}
