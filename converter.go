
package vncalendar

type SolarDate struct {
	Year, Month, Day int
}

type LunarDate struct {
	Year, Month, Day, Leap int

}

func Lunar2solar(lunar_day, lunar_month, lunar_year, lunar_leap, time_zone int) *SolarDate {
	return nil
}

func Solar2lunar(dd, mm,  yyyy, time_zone int) *LunarDate {
	return nil	
}