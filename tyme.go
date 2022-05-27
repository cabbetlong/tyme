package tyme

import (
	"time"
)

type Tyme struct {
	time.Time
}

func (t *Tyme) String() string {
	return t.Format(globalDTLayout)
}

// ==> Comparison

// Equal
func (t *Tyme) Equal(r *Tyme) bool {
	return t.Time.Equal(r.Time)
}

// EqualS
// NOTE: This function is not safely! If tyme cannot parse the string, it will panic.
func (t *Tyme) EqualS(s string) bool {
	r := MustParse(s)

	return t.Time.Equal(r.Time)
}

// Before
func (t *Tyme) Before(r *Tyme) bool {
	return t.Time.Before(r.Time)
}

// BeforeS checks the given time string is before t
// NOTE: This function is not safely! If tyme cannot parse the string, it will panic.
func (t *Tyme) BeforeS(s string) bool {
	r := MustParse(s)

	return t.Time.Before(r.Time)
}

// After
func (t *Tyme) After(r *Tyme) bool {
	return t.Time.After(r.Time)
}

// AfterS checks the given time string is after t.
// NOTE: This function is not safely! If tyme can't parse the string, it will panic.
func (t *Tyme) AfterS(s string) bool {
	r := MustParse(s)

	return t.After(r)
}

// Between
func (t *Tyme) Between(begin, end *Tyme) bool {
	return t.After(begin) && t.Before(end)
}

// AfterS returns true if the time is before the time in string format.
// NOTE: This function is not safely! If tyme can't parse the string, it will panic.
func (t *Tyme) BetweenS(begin, end string) bool {
	return t.AfterS(begin) && t.BeforeS(end)
}

// ==> Basic functions

//BeginOfMinute returns the begin of the minute
func (t *Tyme) BeginOfMinute() *Tyme {
	return With(t.Truncate(time.Minute))
}

// BeginOfHour returns the begin of the hour
func (t *Tyme) BeginOfHour() *Tyme {
	return With(t.Truncate(time.Hour))
}

// BeginOfDay returns the begin of the day
func (t *Tyme) BeginOfDay() *Tyme {
	return With(t.Truncate(24 * time.Hour))
}

// BeginOfWeek returns the begin of the week
func (t *Tyme) BeginOfWeek() *Tyme {
	d := t.BeginOfDay()
	weekday := int(d.Weekday())

	if firstDayOfWeek != time.Sunday {
		weekday = (weekday + 6) % 7
	}

	return With(d.AddDate(0, 0, -weekday))
}

// BeginOfMonth returns the begin of the month
func (t *Tyme) BeginOfMonth() *Tyme {
	y, m, _ := t.Date()
	return With(time.Date(y, m, 1, 0, 0, 0, 0, globalLoc))
}

// BeginOfQuarter returns the begin of the quarter
func (t *Tyme) BeginOfQuarter() *Tyme {
	offset := (int(t.Month()) - 1) % 3
	return t.AddMonths(-offset).BeginOfMonth()
}

// BeginOfYear returns the begin of the year
func (t *Tyme) BeginOfYear() *Tyme {
	y, _, _ := t.Date()
	return With(time.Date(y, time.January, 1, 0, 0, 0, 0, globalLoc))
}

// EndOfWeek returns the end of the week
func (t *Tyme) EndOfMinute() *Tyme {
	return t.BeginOfMinute().AddNanoseconds(int64(time.Minute - time.Nanosecond))
}

// EndOfHour returns the end of the hour
func (t *Tyme) EndOfHour() *Tyme {
	return t.BeginOfHour().AddHours(1).AddNanoseconds(-1)
}

// EndOfDay returns the end of the day
func (t *Tyme) EndOfDay() *Tyme {
	y, m, d := t.Date()
	ot := time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), globalLoc)
	return With(ot)
}

// EndOfWeek returns the end of the week
func (t *Tyme) EndOfWeek() *Tyme {
	return t.BeginOfWeek().AddDays(7).AddNanoseconds(-1)
}

// EndOfMonth returns the end of the month
func (t *Tyme) EndOfMonth() *Tyme {
	return t.BeginOfMonth().AddMonths(1).AddNanoseconds(-1)
}

// EndOfQuarter returns the end of the quarter
func (t *Tyme) EndOfQuarter() *Tyme {
	return t.BeginOfQuarter().AddMonths(3).AddNanoseconds(-1)
}

// EndOfYear returns the end of the year
func (t *Tyme) EndOfYear() *Tyme {
	return t.BeginOfYear().AddYears(1).AddNanoseconds(-1)
}

// ==> Weekdays

// Monday returns the monday time of current week
func (t *Tyme) Monday() *Tyme {
	if firstDayOfWeek == time.Monday {
		return t.BeginOfWeek()
	}

	return t.BeginOfWeek().AddDays(1)
}

// Tuesday returns the tuesday time of current week
func (t *Tyme) Tuesday() *Tyme {
	if firstDayOfWeek == time.Monday {
		return t.BeginOfWeek().AddDays(1)
	}

	return t.BeginOfWeek().AddDays(2)
}

// Wendesday returns the wendesday time of current week
func (t *Tyme) Wendesday() *Tyme {
	if firstDayOfWeek == time.Monday {
		return t.BeginOfWeek().AddDays(2)
	}

	return t.BeginOfWeek().AddDays(3)
}

// Thursday returns the thursday time of current week
func (t *Tyme) Thursday() *Tyme {
	if firstDayOfWeek == time.Monday {
		return t.BeginOfWeek().AddDays(3)
	}

	return t.BeginOfWeek().AddDays(4)
}

// Friday returns the friday time of current week
func (t *Tyme) Friday() *Tyme {
	if firstDayOfWeek == time.Monday {
		return t.BeginOfWeek().AddDays(4)
	}

	return t.BeginOfWeek().AddDays(5)
}

// Saturday returns the saturday time of current week
func (t *Tyme) Saturday() *Tyme {
	if firstDayOfWeek == time.Monday {
		return t.BeginOfWeek().AddDays(5)
	}

	return t.BeginOfWeek().AddDays(6)
}

// Sunday returns the sunday time of current week
func (t *Tyme) Sunday() *Tyme {
	if firstDayOfWeek == time.Monday {
		return t.BeginOfWeek().AddDays(6)
	}

	return t.BeginOfWeek()
}

// ==> Caclulations

// AddNanoseconds adds nanoseconds to the time
func (t *Tyme) AddNanoseconds(nano int64) *Tyme {
	return With(t.Add(time.Duration(nano) * time.Nanosecond))
}

// AddMilliseconds adds millisecond to the time
func (t *Tyme) AddMilliseconds(milli int) *Tyme {
	return With(t.Add(time.Duration(milli) * time.Millisecond))
}

// AddSeconds adds seconds to the time
func (t *Tyme) AddSeconds(seconds int) *Tyme {
	return With(t.Add(time.Duration(seconds) * time.Second))
}

// AddMinutes adds minutes to the time
func (t *Tyme) AddMinutes(minutes int) *Tyme {
	return With(t.Add(time.Duration(minutes) * time.Minute))
}

// AddHours adds hours to the time
func (t *Tyme) AddHours(hours int) *Tyme {
	return With(t.Add(time.Duration(hours) * time.Hour))
}

// AddDays adds days to the time
func (t *Tyme) AddDays(days int) *Tyme {
	return With(t.AddDate(0, 0, days))
}

// AddWeeks adds weeks to the time
func (t *Tyme) AddWeeks(weeks int) *Tyme {
	return With(t.AddDate(0, 0, weeks*7))
}

// AddMonths adds months to the time
func (t *Tyme) AddMonths(months int) *Tyme {
	return With(t.AddDate(0, months, 0))
}

// AddYears adds years to the time
func (t *Tyme) AddYears(years int) *Tyme {
	return With(t.AddDate(years, 0, 0))
}

// ==> Format
// DateLayout
func (t *Tyme) DateLayout() string {
	return t.Format(globalDLayout)
}

// TimeLayout
func (t *Tyme) TimeLayout() string {
	return t.Format(globalTLayout)
}
