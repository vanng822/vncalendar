package vncalendar

import (
	"time"
)

var TimeZoneOffset int

func init() {
	now := time.Now()
	_, offset := now.Zone()
	TimeZoneOffset = int(offset/3600)
}


type VNTime struct {
	solarTime time.Time
	lunarDate LunarDate
	timeZoneOffset int
}

func newVNTime(solarTime time.Time, timeZoneOffset int) VNTime {
	t := VNTime{solarTime: solarTime, timeZoneOffset: timeZoneOffset}
	t.lunarDate = Solar2lunar(t.solarTime.Year(), int(t.solarTime.Month()), t.solarTime.Day(), timeZoneOffset)
	
	return t 
}

func Now() VNTime {
	return newVNTime(time.Now(), TimeZoneOffset)
}

func Date(year, month, day, hour, min, sec, nsec, timeZoneOffset int) VNTime {
	return VNTime{solarTime:time.Now(), timeZoneOffset: TimeZoneOffset}
}

func FromSolarTime(t time.Time) VNTime {
	return newVNTime(t, TimeZoneOffset)
}

func (t VNTime) SolarTime() time.Time {
	return t.solarTime
}

func (t VNTime) Add(d time.Duration) VNTime {
	return VNTime{}
}

func (t VNTime) AddDate(years int, months int, days int) VNTime {
	return newVNTime(t.solarTime.AddDate(years, months, days), t.timeZoneOffset)
}

func (t VNTime) Before(u VNTime) bool {
	return false
}

func (t VNTime) After(u VNTime) bool {
	return false
}

func (t VNTime) Equal(u VNTime) bool {
	return false
}

func (t VNTime) String() string {
	return ""
}

func (t VNTime) Format(layout string) string {
	return ""
}

func (t VNTime) Clock() (hour, min, sec int) {
	return
}

func (t VNTime) Nanosecond (nanosecond int) {
	return
}

func (t VNTime) Second() (second int) {
	return
}

func (t VNTime) Minute() (minute int) {
	return
}

func (t VNTime) Hour() (hour int) {
	return
}

func (t VNTime) Day() int {
	return t.lunarDate.Day
}

func (t VNTime) Date() (year, month, day int) {
	return t.lunarDate.Year, t.lunarDate.Month, t.lunarDate.Day
}

func (t VNTime) Month() int {
	return t.lunarDate.Month
}

func (t VNTime) Year() int {
	return t.lunarDate.Year
}
