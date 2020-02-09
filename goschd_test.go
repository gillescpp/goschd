package goschd

import (
	"testing"
	"time"
)

func TestPeriodListTest(t *testing.T) {
	ph := PeriodList{}
	ph.HoursOnly = true
	ph.TimeSlot = make([][2]time.Time, 0)

	// test 1 : no TimeSlot
	//next = now + interval
	from := time.Time{}
	intr := time.Duration(10 * time.Minute)
	fromResp := time.Now().Add(intr)
	n := ph.NextStepInPeriod(from, intr, false)
	t.Log("NextStepInPeriod", from, intr, false, "=", n)
	if n.Sub(fromResp) < 0 {
		t.Errorf("NextStepInPeriod 1 for %v = %v, not %v", from, n, fromResp)
	}
	n = ph.NextStepInPeriod(from, intr, true)
	t.Log("NextStepInPeriod", from, intr, true, "=", n)
	if n.Sub(fromResp) < 0 {
		t.Errorf("NextStepInPeriod 1 for %v = %v, not %v", from, n, fromResp)
	}
	//next = from + interval
	from = time.Date(2010, 6, 22, 7, 25, 0, 0, time.Local)
	intr = time.Duration(30 * time.Minute)
	n = ph.NextStepInPeriod(from, intr, false)
	fromResp = time.Date(2010, 6, 22, 7, 55, 0, 0, time.Local)
	t.Log("NextStepInPeriod", from, intr, false, "=", n)
	if n != fromResp {
		t.Errorf("NextStepInPeriod 3 for %v = %v, not %v", from, n, fromResp)
	}
	n = ph.NextStepInPeriod(from, intr, true)
	fromResp = time.Date(2010, 6, 22, 8, 00, 0, 0, time.Local)
	t.Log("NextStepInPeriod", from, intr, true, "=", n)
	if n != fromResp {
		t.Errorf("NextStepInPeriod 4 for %v = %v, not %v", from, n, fromResp)
	}

	//test 2: TimeSlot, 9:30 to 10:50 and 21:00 to 22:00
	ph.TimeSlot = append(ph.TimeSlot, [2]time.Time{
		time.Date(0, 0, 0, 21, 00, 0, 0, time.Local),
		time.Date(0, 0, 0, 22, 00, 0, 0, time.Local),
	})
	ph.TimeSlot = append(ph.TimeSlot, [2]time.Time{
		time.Date(0, 0, 0, 9, 30, 0, 0, time.Local),
		time.Date(0, 0, 0, 10, 50, 0, 0, time.Local),
	})

	// 7:10, it=10 mins
	from = time.Date(2010, 6, 25, 7, 10, 0, 0, time.Local)
	intr = time.Duration(10 * time.Minute)
	fromResp = time.Date(0, 0, 0, 9, 30, 0, 0, time.Local)
	n = ph.NextStepInPeriod(from, intr, false)
	t.Log("NextStepInPeriod", from, intr, false, "=", n)
	if n != fromResp {
		t.Errorf("NextStepInPeriod 3 for %v = %v, not %v", from, n, fromResp)
	}

}
