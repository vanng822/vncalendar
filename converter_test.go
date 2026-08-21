package vncalendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolar2Lunar(t *testing.T) {
	result := Solar2lunar(2014, 9, 23, 0)
	assert.Equal(t, 30, result.Day)
	assert.Equal(t, 8, result.Month)
	assert.Equal(t, 2014, result.Year)
	assert.Equal(t, false, result.Leap)
}

func TestLunar2solar(t *testing.T) {
	result := Lunar2solar(2014, 8, 30, false, 0)
	assert.Equal(t, 23, result.Day)
	assert.Equal(t, 9, result.Month)
	assert.Equal(t, 2014, result.Year)
}

func TestSolar2LunarLeapMonth(t *testing.T) {
	lunarDate := Solar2lunar(2006, 9, 12, 7)
	assert.Equal(t, 20, lunarDate.Day)
	assert.Equal(t, 7, lunarDate.Month)
	assert.Equal(t, 2006, lunarDate.Year)
	assert.Equal(t, true, lunarDate.Leap)

	lunarDate = Solar2lunar(2006, 8, 13, 7)
	assert.Equal(t, 20, lunarDate.Day)
	assert.Equal(t, 7, lunarDate.Month)
	assert.Equal(t, 2006, lunarDate.Year)
	assert.Equal(t, false, lunarDate.Leap)

	lunarDate = Solar2lunar(2012, 6, 12, 7)
	assert.Equal(t, 23, lunarDate.Day)
	assert.Equal(t, 4, lunarDate.Month)
	assert.Equal(t, 2012, lunarDate.Year)
	assert.Equal(t, true, lunarDate.Leap)

	lunarDate = Solar2lunar(2012, 5, 13, 7)
	assert.Equal(t, 23, lunarDate.Day)
	assert.Equal(t, 4, lunarDate.Month)
	assert.Equal(t, 2012, lunarDate.Year)
	assert.Equal(t, false, lunarDate.Leap)
}

// Dates before 2000-01-01 give a negative Julian century T, and so a negative
// sun longitude before normalization. Truncating towards zero instead of
// flooring there used to put these a whole month early.
func TestSolar2LunarBefore2000(t *testing.T) {
	for _, tc := range []struct {
		year, month, day             int
		wantYear, wantMonth, wantDay int
	}{
		{1900, 1, 1, 1899, 12, 1},
		{1917, 1, 1, 1916, 12, 8},
		{1936, 1, 1, 1935, 12, 7},
		{1955, 1, 1, 1954, 12, 8},
		{1974, 1, 1, 1973, 12, 9},
		{1990, 1, 1, 1989, 12, 5},
		{1998, 1, 1, 1997, 12, 4},
	} {
		lunarDate := Solar2lunar(tc.year, tc.month, tc.day, 7)
		assert.Equal(t, tc.wantDay, lunarDate.Day)
		assert.Equal(t, tc.wantMonth, lunarDate.Month)
		assert.Equal(t, tc.wantYear, lunarDate.Year)
	}
}

// Tet for a few years before 2000, which the same bug left intact but which
// pin the month boundaries either side of the dates above.
func TestSolar2LunarTetBefore2000(t *testing.T) {
	for _, tc := range []struct{ year, month, day, wantYear int }{
		{1970, 2, 6, 1970},
		{1991, 2, 15, 1991},
		{1999, 2, 16, 1999},
	} {
		lunarDate := Solar2lunar(tc.year, tc.month, tc.day, 7)
		assert.Equal(t, 1, lunarDate.Day)
		assert.Equal(t, 1, lunarDate.Month)
		assert.Equal(t, tc.wantYear, lunarDate.Year)
	}
}

func TestSolar2LunarRoundTrip(t *testing.T) {
	for jd := jdFromDate(1, 1, 1900); jd <= jdFromDate(31, 12, 2100); jd++ {
		solarDate := jdToDate(jd)
		lunarDate := Solar2lunar(solarDate.Year, solarDate.Month, solarDate.Day, 7)
		back := Lunar2solar(lunarDate.Year, lunarDate.Month, lunarDate.Day, lunarDate.Leap, 7)
		if back != solarDate {
			t.Fatalf("round trip failed: %v -> %v -> %v", solarDate, lunarDate, back)
		}
	}
}

func TestLunar2SolarLeapMonth(t *testing.T) {
	solarDate := Lunar2solar(2006, 7, 20, true, 7)
	assert.Equal(t, 12, solarDate.Day)
	assert.Equal(t, 9, solarDate.Month)
	assert.Equal(t, 2006, solarDate.Year)

	solarDate = Lunar2solar(2006, 7, 20, false, 7)
	assert.Equal(t, 13, solarDate.Day)
	assert.Equal(t, 8, solarDate.Month)
	assert.Equal(t, 2006, solarDate.Year)

	solarDate = Lunar2solar(2012, 4, 23, false, 7)
	assert.Equal(t, 13, solarDate.Day)
	assert.Equal(t, 5, solarDate.Month)
	assert.Equal(t, 2012, solarDate.Year)

	solarDate = Lunar2solar(2012, 4, 23, true, 7)
	assert.Equal(t, 12, solarDate.Day)
	assert.Equal(t, 6, solarDate.Month)
	assert.Equal(t, 2012, solarDate.Year)
}
