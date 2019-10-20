package goschd

import (
	"time"
)

const (
	//IntevalType types I : Interval
	IntevalType = "I"
	//FixedHoursType types H : Hours
	FixedHoursType = "H"
)

// PeriodList is a list a from->to period
type PeriodList struct {
	TimeSlot  [][2]time.Time
	HoursOnly bool
}

// NextStepInPeriod return next date
func (p PeriodList) NextStepInPeriod(from time.Time, interval time.Duration, fixedStep bool) time.Time {
	if interval.Nanoseconds() <= 0 {
		return time.Time{}
	} else if len(p.TimeSlot) == 0 {
		//next, no time slot
		if fixedStep {
			return from.Truncate(interval).Add(interval)
		}
		return from.Add(interval)
	} else {
		//search next available
		bt := from
		if fixedStep {
			bt = from.Truncate(interval).Add(interval)
		}

		//test each timeslot
		bestNext := time.Time{}
		for _, c := range p.TimeSlot {
			currentBestNext := time.Time{}
			//invalid timeslot
			if !c[0].IsZero() && !c[1].IsZero() && c[0].After(c[1]) {
				continue
			}
			nbt := bt.Add(interval)
			//use starting time
			if !fixedStep && !c[0].IsZero() && c[0].After(nbt) {
				currentBestNext = c[0]
			} else {
				//inc step...
				for ; ; nbt = nbt.Add(interval) {
					if (c[0].IsZero() || nbt.After(c[0])) && (c[1].IsZero() || nbt.Before(c[1])) {
						currentBestNext = nbt
						break
					}
				}
			}
			//ts evaluation terminated
			if !currentBestNext.IsZero() && (bestNext.IsZero() && currentBestNext.Before(bestNext)) {
				bestNext = currentBestNext
			}
		}
		return bestNext
	}
}

// Monthdays is a list of a month days
type Monthdays struct {
	Days     [31]bool // checked days
	Firstday bool     // first day of the month (1)
	Lastday  bool     // last day of the month according to the current month
}

//TimeIn check if dt is in current Monthdays
/*
func (e *Monthdays) TimeIn(dt time.Time) bool {
	dm := dt.Day()
	//1st day of the month
	if (e.Firstday || e.Days[0]) && (dm == 1) {
		return true
	}
	//last day of the month
	if e.Lastday && (dt.Month() != dt.Add(24*time.Hour).Month()) {
		return true
	}
	//is checked day
	return (e.Days[dm-1])
}
*/

// Event is a schedule définition
type Event struct {
	Type        string        // I/H IntevalType or FixedHoursType
	Interval    time.Duration // for interval type
	HoursPeriod PeriodList    // for interval type
	Hours       []time.Time   // for fixed hours type : list of hours
	Fixed       bool          // for fixed hours type : true = fixed time, false = "from" (event fired even if time excedeed)
	//others common time constrainst (none if empty/nil)
	Weekdays [7]bool   //Allowed weekdays (index: time consts time.Sunday, ...)
	Montdays Monthdays //Allowed monthdays (1-31)
	Months   [12]bool  //Allowed months (0=January)
}

//CalcNext returne the next event time
/*
func (e *Event) CalcNext(from time.Time) {
	var dt time.Time
	if e == nil || (e.Type != IntevalType && e.Type == FixedHoursType ) {
		return dt
	}

	//calc next hour in the day
	if e.Type == IntevalType {
		//next base step
		inc := from.Add(e.Interval)
		if Fixed {
			currentDay := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
			for inc = currentDay.Add(e.Interval) ; inc.Before(from) ; inc = inc.Add(e.Interval) {
			}
		}
		//check
	} else {
		//hours
	}
}*/

// EventSet contains a list a event to launch a task
type EventSet struct {
	Events           []Event
	ForbiddenPeriods PeriodList
}

//CalcNext returne the next event time
/*
func (e *EventSet) CalcNext(from time.Time) {
	nt := time.Time{}
	if e == nil || len(e.Events.Events) == 0 {
		return nt
	}

	//check ForbiddenPeriods : if from is included in a forbidden period
	// , event is postponed to the end of it
	if len(e.ForbiddenPeriods.TimeSlot) > 0 {
		for _, f := range e.ForbiddenPeriods.TimeSlot {
			if from.After(f[0]) && from.Before(f[1]) {
				from = f[1].Add(time.Nanosecond)
			}
		}
	}

	//get the next closest event time in the Events list
	for _, n := range e.Events {
		currentNext := from
		//check constrainsts

		// next eligible month
		wdfound := false
		for i := 0 ; !wdfound && (i<12) ; i++ {
			wdfound = n.Months [int(currentNext.Month)-1]
			if !wdfound {
				currentNext = currentNext.AddDate(0, 1, 0)
			} else i > 0 {
				//ok but initial month of currentNext ko, start at the begining of another month
				currentNext = time.Date(currentNext.Year, currentNext.Month, 1, 0, 0, 0, 0, currentNext.Location())
			}
		}
		if !wdfound {
			continue
		}

		//next eligible day of the mont
		wdfound = false
		currentMonth := currentNext.Month()
		for i := 0 ; !wdfound && (currentMonth == urrentNext.Month()) ; i++ {
			wdfound = n.Montdays.TimeIn(currentMonth)
			if !wdfound {
				currentNext = currentNext.AddDate(0, 0, 1)
			}
		}
		if !wdfound {
			continue
		}

		// next eligible week day
		wdfound = false
		for i := 0 ; !wdfound && (i<7) ; i++ {
			wdfound = n.Weekdays[int(currentNext.Weekday)]
			if !wdfound {
				currentNext = currentNext.AddDate(0, 0, 1)
			}
		}
		if !wdfound {
			continue
		}


		if n.Type == IntevalType {
			from = f[1]
		}
	}




	return nt
}
*/

// Scheduler instance
type Scheduler struct {
}
