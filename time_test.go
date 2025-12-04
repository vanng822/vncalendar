package vncalendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFromSolarTime(t *testing.T) {
	date, _ := time.Parse("Jan 2, 2006 at 3:04pm", "Sep 16, 2014 at 3:04pm")
	lunarTime := FromSolarTime(date)
	assert.Equal(t, 23, lunarTime.Day())
	assert.Equal(t, time.August, lunarTime.Month())
	assert.Equal(t, 2014, lunarTime.Year())
}

func TestAdd(t *testing.T) {
	l, _ := time.Parse("Jan 2, 2006 at 3:04pm", "Sep 16, 2014 at 3:04pm")
	n := l.Add(time.Duration(24 * time.Hour))
	assert.Equal(t, l.Day(), n.Day()-1)
}

func TestLunarDate(t *testing.T) {
	date, _ := time.Parse("Jan 2, 2006 at 3:04pm", "Sep 16, 2014 at 3:04pm")
	lunarTime := FromSolarTime(date)
	y, m, d := lunarTime.Date()
	assert.Equal(t, 2014, y)
	assert.Equal(t, time.August, m)
	assert.Equal(t, 23, d)
}

func TestDate(t *testing.T) {
	date, _ := time.Parse("Jan 2, 2006 at 3:04pm", "Sep 16, 2014 at 3:04pm")
	lunarTime := Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), date.Second(), 0)
	y, m, d := lunarTime.Date()
	assert.Equal(t, 2014, y)
	assert.Equal(t, time.August, m)
	assert.Equal(t, 23, d)
}

func TestDateEqual(t *testing.T) {
	date, _ := time.Parse("Jan 2, 2006 at 3:04pm", "Sep 16, 2014 at 3:04pm")
	date2, _ := time.Parse("Jan 2, 2006 at 3:04pm", "Sep 16, 2014 at 3:04pm")
	lunarTime := FromSolarTime(date)
	lunarTime2 := FromSolarTime(date2)
	assert.True(t, lunarTime.Equal(lunarTime2))
}

func TestFormatDefault(t *testing.T) {
	date, _ := time.Parse("Jan 2, 2006 at 3:04pm", "Sep 16, 2014 at 3:04pm")
	lunarTime := FromSolarTime(date)
	assert.Equal(t, lunarTime.Format(""), "2014-08-23")
}

func TestFormatVietnamese(t *testing.T) {
	date, _ := time.Parse("Jan 2, 2006 at 3:04pm", "Sep 16, 2014 at 3:04pm")
	lunarTime := FromSolarTime(date)
	assert.Equal(t, lunarTime.Format("%[3]s/%[2]s/%[1]s"), "23/08/2014")
}

func TestDateVNTimeZone(t *testing.T) {
	date := Date(2017, time.May, 21, 16, 59, 59, 0)
	assert.Equal(t, 2017, date.Year())
	assert.Equal(t, time.April, date.Month())
	assert.Equal(t, 26, date.Day())

	date2 := Date(2017, time.May, 21, 17, 0, 1, 0)
	assert.Equal(t, 2017, date2.Year())
	assert.Equal(t, time.April, date2.Month())
	assert.Equal(t, 27, date2.Day())
}

func TestIsTheFirstTomorrow(t *testing.T) {
	d := Date(2020, time.April, 22, 12, 12, 12, 0)
	assert.True(t, d.IsTheFirstNextDay())

	d = Date(2020, time.April, 23, 12, 12, 12, 0)
	assert.False(t, d.IsTheFirstNextDay())
}

func TestIsTheFifteenTomorrow(t *testing.T) {
	d := Date(2020, time.April, 6, 12, 12, 12, 0)
	assert.True(t, d.IsTheFifteenNextDay())

	d = Date(2020, time.April, 7, 12, 12, 12, 0)
	assert.False(t, d.IsTheFifteenNextDay())
}

func TestLastDayOfMonth(t *testing.T) {
	d := Date(2018, time.July, 22, 12, 12, 12, 0)

	last := d.LastDayOfMonth()
	assert.Equal(t, 2018, last.Year())
	assert.Equal(t, time.June, last.Month())
	assert.Equal(t, 29, last.Day())
	assert.Equal(t, 2018, last.SolarTime().Year())
	assert.Equal(t, time.August, last.SolarTime().Month())
	assert.Equal(t, 10, last.SolarTime().Day())
}

func TestFirstDayOfMonth(t *testing.T) {
	d := Date(2018, time.July, 22, 12, 12, 12, 0)

	first := d.FirstDayOfMonth()
	assert.Equal(t, 2018, first.Year())
	assert.Equal(t, time.June, first.Month())
	assert.Equal(t, 1, first.Day())
	assert.Equal(t, 2018, first.SolarTime().Year())
	assert.Equal(t, time.July, first.SolarTime().Month())
	assert.Equal(t, 13, first.SolarTime().Day())
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
	invalidDateString := "2023-13-40"
	_, err = ParseDate(invalidDateString)
	assert.Error(t, err)
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
