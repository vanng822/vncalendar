package vncalendar

import (
	"fmt"
	"time"
)

var (
	TimeZoneOffset int
)

func init() {
	now := time.Now()
	_, offset := now.Zone()
	TimeZoneOffset = int(offset / 3600)
}

type VNDate struct {
	solarTime      time.Time
	lunarDate      LunarDate
	timeZoneOffset int
}

func newVNDate(solarTime time.Time, timeZoneOffset int) VNDate {
	t := VNDate{solarTime: solarTime, timeZoneOffset: timeZoneOffset}
	t.lunarDate = Solar2lunar(t.solarTime.Year(), int(t.solarTime.Month()), t.solarTime.Day(), timeZoneOffset)

	return t
}

func Today() VNDate {
	return newVNDate(time.Now(), TimeZoneOffset)
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int) VNDate {
	solarTime := time.Date(year, month, day, hour, min, sec, nsec, time.UTC)
	return newVNDate(solarTime, TimeZoneOffset)
}

func FromSolarTime(t time.Time) VNDate {
	return newVNDate(t, TimeZoneOffset)
}

func (t VNDate) SolarTime() time.Time {
	return t.solarTime
}

func (t VNDate) Add(d time.Duration) VNDate {
	return newVNDate(t.solarTime.Add(d), TimeZoneOffset)
}

func (t VNDate) AddDate(years int, months int, days int) VNDate {
	return newVNDate(t.solarTime.AddDate(years, months, days), t.timeZoneOffset)
}

func (t VNDate) Before(u VNDate) bool {
	return t.solarTime.Before(u.solarTime)
}

func (t VNDate) After(u VNDate) bool {
	return t.solarTime.After(u.solarTime)
}

func (t VNDate) Equal(u VNDate) bool {
	return t.solarTime.Equal(u.solarTime)
}

func (t VNDate) String() string {
	return fmt.Sprintf("%d-%d-%d (%d-%d-%d)", t.Year(), t.Month(), t.Day(), t.solarTime.Year(), t.solarTime.Month(), t.solarTime.Day())
}

func (t VNDate) Format(layout string) string {
	return ""
}

func (t VNDate) Day() int {
	return t.lunarDate.Day
}

func (t VNDate) Date() (year int, month time.Month, day int) {
	return t.lunarDate.Year, time.Month(t.lunarDate.Month), t.lunarDate.Day
}

func (t VNDate) Month() time.Month {
	return time.Month(t.lunarDate.Month)
}

func (t VNDate) Year() int {
	return t.lunarDate.Year
}
