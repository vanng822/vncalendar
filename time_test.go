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
