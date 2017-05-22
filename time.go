package vncalendar

import (
	"fmt"
	"time"
)

var (
	TimeZoneOffset  = 7
	VietNamTimeZone = time.FixedZone("ICT", 7*60*60)
)

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
	return newVNDate(time.Now().In(VietNamTimeZone), TimeZoneOffset)
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int) VNDate {
	solarTime := time.Date(year, month, day, hour, min, sec, nsec, time.UTC).In(VietNamTimeZone)
	return newVNDate(solarTime, TimeZoneOffset)
}

func FromSolarTime(t time.Time) VNDate {
	return newVNDate(t.In(VietNamTimeZone), TimeZoneOffset)
}

func (t VNDate) SolarTime() time.Time {
	return t.solarTime
}

func (t VNDate) Add(d time.Duration) VNDate {
	return newVNDate(t.solarTime.Add(d), t.timeZoneOffset)
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
	return fmt.Sprintf("%s-%s-%s (%s-%s-%s)",
		padd(t.Year()), padd(int(t.Month())), padd(t.Day()),
		padd(t.solarTime.Year()), padd(int(t.solarTime.Month())), padd(t.solarTime.Day()))
}

// Format using Sprintf where inputs are string with zero padd
// First position is year, 2nd month, 3rth day
// Default is %[1]s-%[2]s-%[3]s
func (t VNDate) Format(layout string) string {
	if layout == "" {
		layout = "%[1]s-%[2]s-%[3]s"
	}
	return fmt.Sprintf(layout, padd(t.Year()), padd(int(t.Month())), padd(t.Day()))
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
