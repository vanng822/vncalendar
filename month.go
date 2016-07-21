package vncalendar

import (
	//"fmt"
	"time"
)

var Months = []time.Month{
	time.January,
	time.February,
	time.March,
	time.April,
	time.May,
	time.June,
	time.July,
	time.August,
	time.September,
	time.October,
	time.November,
	time.December,
}

// Given Year/Month in Gregorian Calendar
// Return a list of date in that month with corresponding dates
// Lunar calendar 
func GetMonthDates(year int, month time.Month) []VNDate {
	var dates []VNDate
	start := time.Date(year, month, 1, 12, 0, 0, 1, time.UTC)
	for i := 0; i <= 31; i++ {
		d := FromSolarTime(start.AddDate(0, 0, i))
		// next month
		if d.SolarTime().Month() != month {
			break
		}
		dates = append(dates, d)
	}
	return dates
}

func GetYearMonthDates(year int) map[time.Month][]VNDate {
	months := make(map[time.Month][]VNDate)
	for _, m := range Months {
		months[m] = GetMonthDates(year, m)
	}
	
	return months
}
