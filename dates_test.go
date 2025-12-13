package vncalendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDatesBetween(t *testing.T) {
	fromDate := Date(2025, time.December, 10, 12, 12, 12, 0)
	toDate := Date(2025, time.December, 15, 12, 12, 12, 0)
	dates := GetDatesBetween(fromDate, toDate)
	expected := []VNDate{
		fromDate,
		fromDate.AddDate(0, 0, 1),
		fromDate.AddDate(0, 0, 2),
		fromDate.AddDate(0, 0, 3),
		fromDate.AddDate(0, 0, 4),
		fromDate.AddDate(0, 0, 5),
	}
	assert.Equal(t, len(dates), 6)
	assert.Equal(t, dates, expected)
	// including from and to date
	assert.Equal(t, dates[0].solarTime.Day(), 10)
	assert.Equal(t, dates[5].solarTime.Day(), 15)
}

func TestParseFromSolarString(t *testing.T) {
	testDateString := "2025-12-02 12:04:05.999999999 +0700 ICT"
	lunarDate, err := ParseFromSolarString(testDateString, DefaultSolarLayout)
	assert.NoError(t, err)
	assert.Equal(t, 13, lunarDate.Day())
	assert.Equal(t, time.October, lunarDate.Month())
	assert.Equal(t, 2025, lunarDate.Year())

	invalidDateString := "2023-13-40 12:04:05.999999999 +0700 ICT"
	_, err = ParseFromSolarString(invalidDateString, DefaultSolarLayout)
	assert.Error(t, err)
}

func TestParseDate(t *testing.T) {
	testDateString := "2023-10-15"
	d, err := ParseDate(testDateString)
	assert.NoError(t, err)
	assert.Equal(t, 2023, d.Year())
	assert.Equal(t, time.October, d.Month())
	assert.Equal(t, 15, d.Day())

	invalidDateString := "2023-13-20"
	_, err = ParseDate(invalidDateString)
	assert.Error(t, err)
	assert.Equal(t, "invalid date - month", err.Error())

	invalidDateString = "2025-08-32"
	_, err = ParseDate(invalidDateString)
	assert.Error(t, err)
	assert.Equal(t, "invalid date - day", err.Error())

	// only 29 days for this month in lunar calendar
	invalidDateString = "2025-08-30"
	_, err = ParseDate(invalidDateString)
	assert.Error(t, err)
	assert.Equal(t, "invalid date", err.Error())

	invalidDateString = "1790-08-20"
	_, err = ParseDate(invalidDateString)
	assert.Error(t, err)
	assert.Equal(t, "not supported year range", err.Error())

	// auto test many dates
	startDate := time.Now()
	for i := range 1001 {
		dt := startDate.AddDate(0, 0, i)
		lunarDate := FromSolarTime(dt)
		dateString := lunarDate.Format("%[1]s-%[2]s-%[3]s")
		parsedDate, err := ParseDate(dateString)
		assert.NoError(t, err)
		assert.Equal(t, lunarDate.Year(), parsedDate.Year())
		assert.Equal(t, lunarDate.Month(), parsedDate.Month())
		assert.Equal(t, lunarDate.Day(), parsedDate.Day())
	}
}

func TestSub(t *testing.T) {
	// dynamic test
	now := time.Now()
	vnDate1 := FromSolarTime(now)
	vnDate2 := FromSolarTime(now.Add(48 * time.Hour))
	diff := vnDate2.Sub(vnDate1)
	assert.Equal(t, 48*time.Hour, diff)

	// static test
	date1, _ := time.Parse("2006-01-02 15:04:05", "2014-09-16 15:04:00")
	date2, _ := time.Parse("2006-01-02 15:04:05", "2014-09-17 14:04:00")
	vnDate3 := FromSolarTime(date1)
	vnDate4 := FromSolarTime(date2)
	diff = vnDate4.Sub(vnDate3)
	assert.Equal(t, 23*time.Hour, diff)
}
