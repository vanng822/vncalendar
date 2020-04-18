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
