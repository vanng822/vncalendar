package vncalendar

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolar2Lunar(t *testing.T) {
	result := Solar2lunar(2014, 9, 23, 0)
	if result.Day != 30 {
		t.Errorf("Day expected to be 30 but got '%s'", result.Day)
	}
	if result.Month != 8 {
		t.Errorf("Month expected to be 8 but got '%s'", result.Month)
	}
	if result.Year != 2014 {
		t.Errorf("Year expected to be 2014 but got '%s'", result.Year)
	}
	
	assert.Equal(t, false, result.Leap)
}


func TestLunar2solar(t *testing.T) {
	result := Lunar2solar(2014, 8, 30, 0, 0)
	if result.Day != 23 {
		t.Errorf("Day expected to be 23 but got '%s'", result.Day)
	}
	if result.Month != 9 {
		t.Errorf("Month expected to be 9 but got '%s'", result.Month)
	}
	if result.Year != 2014 {
		t.Errorf("Year expected to be 2014 but got '%s'", result.Year)
	}
}

func TestSolar2LunarLeapMonth(t *testing.T) {
	lunarDate := Solar2lunar(2006, 9, 12, 7)
	assert.Equal(t, 20, lunarDate.Day)
	assert.Equal(t, 7, lunarDate.Month)
	assert.Equal(t, 2006, lunarDate.Year)
	assert.Equal(t, true, lunarDate.Leap)
}

func TestLunar2SolarLeapMonth(t *testing.T) {
	lunarDate := Lunar2solar(2006, 7, 20, 7, 7)
	assert.Equal(t, 12, lunarDate.Day)
	assert.Equal(t, 9, lunarDate.Month)
	assert.Equal(t, 2006, lunarDate.Year)
}