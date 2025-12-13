package vncalendar

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func GetDatesBetween(fromDate, toDate VNDate) []VNDate {
	var dates []VNDate
	for fromDate.Before(toDate) {
		dates = append(dates, fromDate)
		fromDate = fromDate.AddDate(0, 0, 1)
	}
	dates = append(dates, fromDate)
	return dates
}

const DefaultSolarLayout = "2006-01-02 15:04:05.999999999 -0700 MST"

// Layout uses time package layout format
// if error occurs, VNDate with zero value is returned
func ParseFromSolarString(dateStr, layout string) (VNDate, error) {
	if layout == "" {
		layout = DefaultSolarLayout
	}
	solarTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return VNDate{}, err
	}
	return FromSolarTime(solarTime), nil
}

var (
	dateFormatRe             = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)
	errInvalidDateFormat     = errors.New("invalid date format")
	errNotSupportedYearRange = errors.New("not supported year range")
	errInvalidMonth          = errors.New("invalid date - month")
	errInvalidDay            = errors.New("invalid date - day")
	errInvalidDate           = errors.New("invalid date")
)

// ParseDate parse date string in format "YYYY-MM-DD"
// date: string lunar date in format "YYYY-MM-DD"
// return zero value VNDate and error if invalid format or invalid date
func ParseDate(date string) (VNDate, error) {
	var (
		year, day int
		month     int
		err       error
	)

	res := dateFormatRe.FindStringSubmatch(date)

	if len(res) != 4 {
		return VNDate{}, errInvalidDateFormat
	}
	year, err = strconv.Atoi(res[1])
	if err != nil {
		return VNDate{}, err
	}

	// Unsure how good the algorithm is outside this range so limit it for now
	if 1800 > year || year > 2040 {
		return VNDate{}, errNotSupportedYearRange
	}

	month, err = strconv.Atoi(res[2])
	if err != nil {
		return VNDate{}, err
	}
	if 1 > month || month > 12 {
		return VNDate{}, errInvalidMonth
	}

	day, err = strconv.Atoi(res[3])
	if err != nil {
		return VNDate{}, errInvalidDay
	}
	if 1 > day || day > 31 {
		return VNDate{}, errInvalidDay
	}

	valid, validDate := Validate(year, month, day)
	if !valid {
		return VNDate{}, errInvalidDate
	}

	return validDate, nil
}

func Validate(year, month, day int) (bool, VNDate) {
	// just convert back and forth to verify date
	testSolar := Lunar2solar(year, month, day, false, TimeZoneOffset)
	testLunar := Solar2lunar(testSolar.Year, testSolar.Month, testSolar.Day, TimeZoneOffset)

	if testLunar.Year != year || testLunar.Month != month || testLunar.Day != day {
		return false, VNDate{}
	}

	return true, Date(testSolar.Year, time.Month(testSolar.Month), testSolar.Day, 12, 0, 0, 0)
}
