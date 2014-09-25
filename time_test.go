package vncalendar

import (
	"testing"
	"time"
	//"fmt"
	"github.com/stretchr/testify/assert"
)

func TestFromSolarTime(t *testing.T) {
    date, _ := time.Parse("Jan 2, 2006 at 3:04pm", "Sep 16, 2014 at 3:04pm")
    lunarTime := FromSolarTime(date)
    assert.Equal(t, 23, lunarTime.Day())
	assert.Equal(t, 8, lunarTime.Month())
	assert.Equal(t, 2014, lunarTime.Year())
}