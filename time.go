package vncalendar

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
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

func (t VNDate) LunarDate() LunarDate {
	return t.lunarDate
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

func (t VNDate) IsTheFirstNextDay() bool {
	d := t.AddDate(0, 0, 1)
	return d.Day() == 1
}

func (t VNDate) IsTheFifteenNextDay() bool {
	d := t.AddDate(0, 0, 1)
	return d.Day() == 15
}

func (t VNDate) LastDayOfMonth() VNDate {
	var date VNDate
	date = t
	d := t
	for range 31 {
		d = d.AddDate(0, 0, 1)
		// next month
		if d.LunarDate().Day == 1 {
			date = d.AddDate(0, 0, -1)
			break
		}
	}
	return date
}

func (t VNDate) FirstDayOfMonth() VNDate {
	date := t
	start := t
	for range 31 {
		// first day
		if start.Day() == 1 {
			date = start
			break
		}
		start = start.AddDate(0, 0, -1)
	}
	return date
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

var dateFormatRe = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)

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
		return VNDate{}, errors.New("invalid date format")
	}
	year, err = strconv.Atoi(res[1])
	if err != nil {
		return VNDate{}, errors.New("invalid date - year")
	}

	// Unsure how good the algorithm is outside this range so limit it for now
	if 1800 > year || year > 2040 {
		return VNDate{}, errors.New("not supported year range")
	}

	month, err = strconv.Atoi(res[2])
	if err != nil {
		return VNDate{}, err
	}
	if 1 > month || month > 12 {
		return VNDate{}, errors.New("invalid date - month")
	}

	day, err = strconv.Atoi(res[3])
	if err != nil {
		return VNDate{}, errors.New("invalid date - day")
	}
	if 1 > day || day > 31 {
		return VNDate{}, errors.New("invalid date - day")
	}

	valid, validDate := Validate(year, month, day)
	if !valid {
		return VNDate{}, errors.New("invalid date")
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
