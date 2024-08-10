package vncalendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMonthDates(t *testing.T) {
	assert.Equal(t, 31, len(GetMonthDates(2016, time.July)))
	// Leap
	assert.Equal(t, 29, len(GetMonthDates(2016, time.February)))
	assert.Equal(t, 28, len(GetMonthDates(2017, time.February)))
	assert.Equal(t, 30, len(GetMonthDates(2016, time.November)))
	assert.Equal(t, 29, len(GetMonthDates(2024, time.February)))
}

func TestGetYearMonthDates(t *testing.T) {
	assert.Equal(t, 12, len(GetYearMonthDates(2016)))
}
