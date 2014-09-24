package vncalendar

import (
	"time"
)

type VNTime struct {
	solarTime time.Time
}

func Now() VNTime {
	return VNTime{}
}

func Date(year, month, day, hour, min, sec, nsec, timeZoneOffset int) VNTime {
	return VNTime{}
}

func FromSolarTime(t time.Time) VNTime {
	return VNTime{}
}

func (t VNTime) SolarTime() time.Time {
	return time.Time{}
}

func (t VNTime) Add(d time.Duration) VNTime {
	return VNTime{}
}

func (t VNTime) AddDate(years int, months int, days int) VNTime {
	return VNTime{}
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

func (t VNTime) Date() (year, month, day int) {
	return
}

func (t VNTime) Month() (month int) {
	return
}

func (t VNTime) Year() (year int) {
	return
}
