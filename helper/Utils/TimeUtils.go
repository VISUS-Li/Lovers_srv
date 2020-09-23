package Utils

import (
	"fmt"
	"time"
)

func GetThisWeekStartEnd(){
	now := time.Now()
	mondayOffset := int(time.Monday - now.Weekday())
	if mondayOffset > 0 {
		mondayOffset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(),0,0,0,0,time.Local).AddDate(0,0,mondayOffset)
	fmt.Println(weekStart)
}
