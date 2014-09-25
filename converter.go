package vncalendar

import (
	"math"
)

type SolarDate struct {
	Year, Month, Day int
}

type LunarDate struct {
	Year, Month, Day int
	Leap             bool
}

func jdFromDate(dd, mm, yyyy int) int {
	var a, y, m, jd int
	a = int((14 - mm) / 12)
	y = yyyy + 4800 - a
	m = mm + 12*a - 3
	jd = dd + int((153*m+2)/5) + 365*y + int(y/4) - int(y/100) + int(y/400) - 32045
	if jd < 2299161 {
		jd = dd + int((153*m+2)/5) + 365*y + int(y/4) - 32083
	}
	return jd
}

func jdToDate(jd int) SolarDate {
	var a, b, c, d, e, m, day, month, year int
	date := SolarDate{}

	if jd > 2299160 {
		// After 5/10/1582, Gregorian calendar
		a = jd + 32044
		b = int((4*a + 3) / 146097)
		c = a - int((b*146097)/4)
	} else {
		b = 0
		c = jd + 32082
	}
	d = int((4*c + 3) / 1461)
	e = c - int((1461*d)/4)
	m = int((5*e + 2) / 153)
	day = e - int((153*m+2)/5) + 1
	month = m + 3 - 12*int(m/10)
	year = b*100 + d - 4800 + int(m/10)
	date.Day = day
	date.Month = month
	date.Year = year

	return date
}

func newMoon(ak int) float64 {
	var T, T2, T3, dr, Jd1, M, mPr, F, C1, deltat float64
	k := float64(ak)
	T = k / 1236.85 // Time in Julian centuries from 1900 January 0.5
	T2 = T * T
	T3 = T2 * T
	dr = math.Pi / 180
	Jd1 = 2415020.75933 + 29.53058868*k + 0.0001178*T2 - 0.000000155*T3
	Jd1 = Jd1 + 0.00033*math.Sin((166.56+132.87*T-0.009173*T2)*dr)  // Mean new moon
	M = 359.2242 + 29.10535608*k - 0.0000333*T2 - 0.00000347*T3     // Sun's mean anomaly
	mPr = 306.0253 + 385.81691806*k + 0.0107306*T2 + 0.00001236*T3 // Moon's mean anomaly
	F = 21.2964 + 390.67050646*k - 0.0016528*T2 - 0.00000239*T3     // Moon's argument of latitude
	C1 = (0.1734-0.000393*T)*math.Sin(M*dr) + 0.0021*math.Sin(2*dr*M)
	C1 = C1 - 0.4068*math.Sin(mPr*dr) + 0.0161*math.Sin(dr*2*mPr)
	C1 = C1 - 0.0004*math.Sin(dr*3*mPr)
	C1 = C1 + 0.0104*math.Sin(dr*2*F) - 0.0051*math.Sin(dr*(M+mPr))
	C1 = C1 - 0.0074*math.Sin(dr*(M-mPr)) + 0.0004*math.Sin(dr*(2*F+M))
	C1 = C1 - 0.0004*math.Sin(dr*(2*F-M)) - 0.0006*math.Sin(dr*(2*F+mPr))
	C1 = C1 + 0.0010*math.Sin(dr*(2*F-mPr)) + 0.0005*math.Sin(dr*(2*mPr+M))
	if T < -11 {
		deltat = 0.001 + 0.000839*T + 0.0002261*T2 - 0.00000845*T3 - 0.000000081*T*T3
	} else {
		deltat = -0.000278 + 0.000265*T + 0.000262*T2
	}
	
	return Jd1 + C1 - deltat
}

func sunLongitude(jdn float64) float64 {
	var T, T2, dr, M, L0, DL, L float64
	T = (jdn - 2451545.0) / 36525.0 // Time in Julian centuries from 2000-01-01 12:00:00 GMT
	T2 = T * T
	dr = math.Pi / 180                                             // degree to radian
	M = 357.52910 + 35999.05030*T - 0.0001559*T2 - 0.00000048*T*T2 // mean anomaly, degree
	L0 = 280.46645 + 36000.76983*T + 0.0003032*T2                  // mean longitude, degree
	DL = (1.914600 - 0.004817*T - 0.000014*T2) * math.Sin(dr*M)
	DL = DL + (0.019993-0.000101*T)*math.Sin(dr*2*M) + 0.000290*math.Sin(dr*3*M)
	L = L0 + DL // true longitude, degree
	L = L * dr
	L = L - math.Pi*2*float64(int(L/(math.Pi*2))) // Normalize to (0, 2*PI)
	return L
}

func getSunLongitude(jd, timeZoneOffset int) int {
	return int(sunLongitude(float64(jd)-float64(0.5)-float64(timeZoneOffset)/24.0) / math.Pi * 6)
}

func getNewMoonDay(k, timeZoneOffset int) int {
	return int(newMoon(k) + 0.5 + float64(timeZoneOffset)/24)
}

func getLunarMonth11(yyyy, timeZoneOffset int) int {
	var k, off, nm, sunLong int
	off = jdFromDate(31, 12, yyyy) - 2415021
	k = int(float64(off) / 29.530588853)
	nm = getNewMoonDay(k, timeZoneOffset)
	sunLong = getSunLongitude(nm, timeZoneOffset) // sun longitude at local midnight
	if sunLong >= 9 {
		nm = getNewMoonDay(k-1, timeZoneOffset)
	}
	return nm
}

func getLeapMonthOffset(a11, timeZoneOffset int) int {
	var k, last, arc, i int
	k = int((float64(a11)-2415021.076998695)/29.530588853 + 0.5)
	last = 0
	i = 1 // We start with the month following lunar month 11
	arc = getSunLongitude(getNewMoonDay(k+i, timeZoneOffset), timeZoneOffset)

	for arc != last && i < 14 {
		last = arc
		i++
		arc = getSunLongitude(getNewMoonDay(k+i, timeZoneOffset), timeZoneOffset)
	}
	return i - 1
}

func Solar2lunar(yyyy, mm, dd, timeZoneOffset int) LunarDate {
	var k, dayNumber, monthStart, a11, b11, lunarDay, lunarMonth, lunarYear,
		diff, leapMonthDiff int
	var lunarLeap bool

	dayNumber = jdFromDate(dd, mm, yyyy)

	k = int((float64(dayNumber) - 2415021.076998695) / 29.530588853)
	monthStart = getNewMoonDay(k+1, timeZoneOffset)
	if monthStart > dayNumber {
		monthStart = getNewMoonDay(k, timeZoneOffset)
	}
	a11 = getLunarMonth11(yyyy, timeZoneOffset)
	b11 = a11
	if a11 >= monthStart {
		lunarYear = yyyy
		a11 = getLunarMonth11(yyyy-1, timeZoneOffset)
	} else {
		lunarYear = yyyy + 1
		b11 = getLunarMonth11(yyyy+1, timeZoneOffset)
	}
	lunarDay = dayNumber - monthStart + 1
	diff = int((monthStart - a11) / 29)
	lunarLeap = false
	lunarMonth = diff + 11
	if b11-a11 > 365 {
		leapMonthDiff = getLeapMonthOffset(a11, timeZoneOffset)
		if diff >= leapMonthDiff {
			lunarMonth = diff + 10
			if diff == leapMonthDiff {
				lunarLeap = true
			}
		}
	}
	if lunarMonth > 12 {
		lunarMonth = lunarMonth - 12
	}
	if lunarMonth >= 11 && diff < 4 {
		lunarYear -= 1
	}
	res := LunarDate{}
	res.Day = lunarDay
	res.Month = lunarMonth
	res.Year = lunarYear
	res.Leap = lunarLeap
	return res
}

func Lunar2solar(lunarYear, lunarMonth, lunarDay, lunarLeap, timeZoneOffset int) SolarDate {
	var k, a11, b11, off, leapOff, leapMonth, monthStart int

	if lunarMonth < 11 {
		a11 = getLunarMonth11(lunarYear-1, timeZoneOffset)
		b11 = getLunarMonth11(lunarYear, timeZoneOffset)
	} else {
		a11 = getLunarMonth11(lunarYear, timeZoneOffset)
		b11 = getLunarMonth11(lunarYear+1, timeZoneOffset)
	}
	k = int(0.5 + (float64(a11)-2415021.076998695)/29.530588853)
	off = lunarMonth - 11
	if off < 0 {
		off += 12
	}
	if b11-a11 > 365 {
		leapOff = getLeapMonthOffset(a11, timeZoneOffset)
		leapMonth = leapOff - 2
		if leapMonth < 0 {
			leapMonth += 12
		}
		if lunarLeap != 0 && lunarMonth != lunarLeap {
			return SolarDate{Day: 0, Month: 0, Year: 0}
		} else if lunarLeap != 0 || off >= leapOff {
			off += 1
		}
	}
	monthStart = getNewMoonDay(k+off, timeZoneOffset)
	return jdToDate(monthStart + lunarDay - 1)
}
