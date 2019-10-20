package goschd

import (
	"testing"
	"time"
)

func TestPeriodListTest(t *testing.T) {
	ph := PeriodList{}
	ph.HoursOnly = true
	ph.TimeSlot = make([][2]time.Time, 0)
	ph.TimeSlot = append(ph.TimeSlot, [2]time.Time{
		time.Date(0, 0, 0, 8, 30, 0, 0, nil),
		time.Date(0, 0, 0, 10, 50, 0, 0, nil),
	})

	from := time.Date(2010, 6, 25, 7, 10, 0, 0, nil)
	intr := time.Duration(10 * time.Minute)

	fromResp := time.Date(2010, 6, 25, 8, 30, 0, 0, nil)
	n := ph.NextStepInPeriod(from, intr, false)
	if n != fromResp {
		t.Errorf("NextStepInPeriod for %v = %v, not %v", from, n, fromResp)
	}
}
