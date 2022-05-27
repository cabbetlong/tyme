package tyme

import (
	"testing"
	"time"
)

func TestTyme_BeforeS(t *testing.T) {
	assert := assertComparasion(t)

	assert(Parse("2022-05-20 16:14:53").BeforeS("2022-05-20 16:14:54"), true)
	assert(Parse("2022-05-20 16:14:54").BeforeS("2022-05-20 16:14:53"), false)
	assert(Parse("2022-05-20 16:14:54").BeforeS("2022-05-20 16:14:54"), false)
}

func TestTyme_AfterS(t *testing.T) {
	assert := assertComparasion(t)

	assert(Parse("2022-05-20 16:14:53").AfterS("2022-05-20 16:14:54"), false)
	assert(Parse("2022-05-20 16:14:54").AfterS("2022-05-20 16:14:53"), true)
	assert(Parse("2022-05-20 16:14:54").AfterS("2022-05-20 16:14:54"), false)
}

func TestTyme_BetweenS(t *testing.T) {
	assert := assertComparasion(t)

	assert(Parse("2022-05-20 16:14:54").BetweenS("2022-05-20 16:14:53", "2022-05-20 16:14:55"), true)
	assert(Parse("2022-05-20 16:14:53").BetweenS("2022-05-20 16:14:53", "2022-05-20 16:14:55"), false)
	assert(Parse("2022-05-20 16:14:55").BetweenS("2022-05-20 16:14:53", "2022-05-20 16:14:55"), false)
	assert(Parse("2022-05-20 16:14:52").BetweenS("2022-05-20 16:14:53", "2022-05-20 16:14:55"), false)
	assert(Parse("2022-05-20 16:14:56").BetweenS("2022-05-20 16:14:53", "2022-05-20 16:14:55"), false)
}

func TestTyme_BeginOfMinute(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").BeginOfMinute(), "2022-05-20 16:14:00")
	assert(Parse("2022-05-20 16:14:58.99999").BeginOfMinute(), "2022-05-20 16:14:00")
}

func TestTyme_BeginOfHour(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").BeginOfHour(), "2022-05-20 16:00:00")
	assert(Parse("2022-05-20 17:14:58.99999").BeginOfHour(), "2022-05-20 17:00:00")
}

func TestTyme_BeginOfDay(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").BeginOfDay(), "2022-05-20 00:00:00")
	assert(Parse("2021-04-20 17:14:58.99999").BeginOfDay(), "2021-04-20 00:00:00")
}

func TestTyme_BeginOfWeek(t *testing.T) {
	assert := assertT(t)

	// Monday is the first day of the week
	assert(Parse("2022-05-16").BeginOfWeek(), "2022-05-16 00:00:00") // Monday
	assert(Parse("2022-05-17").BeginOfWeek(), "2022-05-16 00:00:00") // Tuesday
	assert(Parse("2022-05-18").BeginOfWeek(), "2022-05-16 00:00:00") // Wednesday
	assert(Parse("2022-05-19").BeginOfWeek(), "2022-05-16 00:00:00") // Thursday
	assert(Parse("2022-05-20").BeginOfWeek(), "2022-05-16 00:00:00") // Friday
	assert(Parse("2022-05-21").BeginOfWeek(), "2022-05-16 00:00:00") // Saturday
	assert(Parse("2022-05-22").BeginOfWeek(), "2022-05-16 00:00:00") // Sunday

	// Sunday is the first day of the week
	SetFirstDayOfWeek(time.Sunday)
	assert(Parse("2022-05-15").BeginOfWeek(), "2022-05-15 00:00:00") // Sunday
	assert(Parse("2022-05-16").BeginOfWeek(), "2022-05-15 00:00:00") // Monday
	assert(Parse("2022-05-17").BeginOfWeek(), "2022-05-15 00:00:00") // Tuesday
	assert(Parse("2022-05-18").BeginOfWeek(), "2022-05-15 00:00:00") // Wednesday
	assert(Parse("2022-05-19").BeginOfWeek(), "2022-05-15 00:00:00") // Thursday
	assert(Parse("2022-05-20").BeginOfWeek(), "2022-05-15 00:00:00") // Friday
	assert(Parse("2022-05-21").BeginOfWeek(), "2022-05-15 00:00:00") // Saturday

	SetFirstDayOfWeek(time.Monday)
}

func TestTyme_BeginOfMonth(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").BeginOfMonth(), "2022-05-01 00:00:00")
	assert(Parse("2022-04-12 17:00:53").BeginOfMonth(), "2022-04-01 00:00:00")
}

func TestTyme_BeginOfQuarter(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-01-20 16:14:53").BeginOfQuarter(), "2022-01-01 00:00:00")
	assert(Parse("2022-02-20 16:14:53").BeginOfQuarter(), "2022-01-01 00:00:00")
	assert(Parse("2022-03-20 16:14:53").BeginOfQuarter(), "2022-01-01 00:00:00")
	assert(Parse("2022-04-20 16:14:53").BeginOfQuarter(), "2022-04-01 00:00:00")
	assert(Parse("2022-05-20 16:14:53").BeginOfQuarter(), "2022-04-01 00:00:00")
	assert(Parse("2022-06-20 16:14:53").BeginOfQuarter(), "2022-04-01 00:00:00")
	assert(Parse("2022-07-20 16:14:53").BeginOfQuarter(), "2022-07-01 00:00:00")
	assert(Parse("2022-08-20 16:14:53").BeginOfQuarter(), "2022-07-01 00:00:00")
	assert(Parse("2022-09-20 16:14:53").BeginOfQuarter(), "2022-07-01 00:00:00")
	assert(Parse("2022-10-20 16:14:53").BeginOfQuarter(), "2022-10-01 00:00:00")
	assert(Parse("2022-11-20 16:14:53").BeginOfQuarter(), "2022-10-01 00:00:00")
	assert(Parse("2022-12-20 16:14:53").BeginOfQuarter(), "2022-10-01 00:00:00")
}

func TestTyme_BeginOfYear(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").BeginOfYear(), "2022-01-01 00:00:00")
	assert(Parse("2021-05-20 16:14:53").BeginOfYear(), "2021-01-01 00:00:00")
}

func TestTyme_EndOfMinute(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").EndOfMinute(), "2022-05-20 16:14:59.999999999")
	assert(Parse("2022-05-20 16:14:58.99999").EndOfMinute(), "2022-05-20 16:14:59.999999999")
}

func TestTyme_EndOfHour(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").EndOfHour(), "2022-05-20 16:59:59.999999999")
	assert(Parse("2022-05-20 17:14:58.99999").EndOfHour(), "2022-05-20 17:59:59.999999999")
}

func TestTyme_EndOfDay(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").EndOfDay(), "2022-05-20 23:59:59.999999999")
	assert(Parse("2021-04-20 17:14:58.99999").EndOfDay(), "2021-04-20 23:59:59.999999999")
}

func TestTyme_EndOfWeek(t *testing.T) {
	assert := assertT(t)

	// Monday is the first day of the week
	assert(Parse("2022-05-16").EndOfWeek(), "2022-05-22 23:59:59.999999999") // Monday
	assert(Parse("2022-05-17").EndOfWeek(), "2022-05-22 23:59:59.999999999") // Tuesday
	assert(Parse("2022-05-18").EndOfWeek(), "2022-05-22 23:59:59.999999999") // Wednesday
	assert(Parse("2022-05-19").EndOfWeek(), "2022-05-22 23:59:59.999999999") // Thursday
	assert(Parse("2022-05-20").EndOfWeek(), "2022-05-22 23:59:59.999999999") // Friday
	assert(Parse("2022-05-21").EndOfWeek(), "2022-05-22 23:59:59.999999999") // Saturday
	assert(Parse("2022-05-22").EndOfWeek(), "2022-05-22 23:59:59.999999999") // Sunday

	// Sunday is the first day of the week
	SetFirstDayOfWeek(time.Sunday)
	assert(Parse("2022-05-15").EndOfWeek(), "2022-05-21 23:59:59.999999999") // Sunday
	assert(Parse("2022-05-16").EndOfWeek(), "2022-05-21 23:59:59.999999999") // Monday
	assert(Parse("2022-05-17").EndOfWeek(), "2022-05-21 23:59:59.999999999") // Tuesday
	assert(Parse("2022-05-18").EndOfWeek(), "2022-05-21 23:59:59.999999999") // Wednesday
	assert(Parse("2022-05-19").EndOfWeek(), "2022-05-21 23:59:59.999999999") // Thursday
	assert(Parse("2022-05-20").EndOfWeek(), "2022-05-21 23:59:59.999999999") // Friday
	assert(Parse("2022-05-21").EndOfWeek(), "2022-05-21 23:59:59.999999999") // Saturday

	SetFirstDayOfWeek(time.Monday)
}

func TestTyme_EndOfMonth(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").EndOfMonth(), "2022-05-31 23:59:59.999999999")
	assert(Parse("2022-04-12 17:00:53").EndOfMonth(), "2022-04-30 23:59:59.999999999")
}

func TestTyme_EndOfQuarter(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-01-20 16:14:53").EndOfQuarter(), "2022-03-31 23:59:59.999999999")
	assert(Parse("2022-02-20 16:14").EndOfQuarter(), "2022-03-31 23:59:59.999999999")
	assert(Parse("2022-03-20 16").EndOfQuarter(), "2022-03-31 23:59:59.999999999")
	assert(Parse("2022-04-20").EndOfQuarter(), "2022-06-30 23:59:59.999999999")
	assert(Parse("2022-05-20").EndOfQuarter(), "2022-06-30 23:59:59.999999999")
	assert(Parse("2022-06-20").EndOfQuarter(), "2022-06-30 23:59:59.999999999")
	assert(Parse("2022-07-20").EndOfQuarter(), "2022-09-30 23:59:59.999999999")
	assert(Parse("2022-08-20").EndOfQuarter(), "2022-09-30 23:59:59.999999999")
	assert(Parse("2022-09-20").EndOfQuarter(), "2022-09-30 23:59:59.999999999")
	assert(Parse("2022-10-20").EndOfQuarter(), "2022-12-31 23:59:59.999999999")
	assert(Parse("2022-11-20").EndOfQuarter(), "2022-12-31 23:59:59.999999999")
	assert(Parse("2022-12-20").EndOfQuarter(), "2022-12-31 23:59:59.999999999")
}

func TestTyme_EndOfYear(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20 16:14:53").EndOfYear(), "2022-12-31 23:59:59.999999999")
	assert(Parse("2021-05-20 16:14:53").EndOfYear(), "2021-12-31 23:59:59.999999999")
}

func TestTyme_AddNanoseconds(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20").AddNanoseconds(-1), "2022-05-19 23:59:59.999999999")
	assert(Parse("2022-05-20 00:00:00").AddNanoseconds(1), "2022-05-20 00:00:00.000000001")
}

func TestTyme_AddAddMilliseconds(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20").AddMilliseconds(-1), "2022-05-19 23:59:59.999")
	assert(Parse("2022-05-20 00:00:00").AddMilliseconds(1), "2022-05-20 00:00:00.001")
}

func TestTyme_AddSeconds(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20").AddSeconds(-1), "2022-05-19 23:59:59")
	assert(Parse("2022-05-20").AddSeconds(-60), "2022-05-19 23:59:00")
	assert(Parse("2022-05-20 00:00:00").AddSeconds(1), "2022-05-20 00:00:01")
	assert(Parse("2022-05-20 00:00:00").AddSeconds(60), "2022-05-20 00:01:00")
}

func TestTyme_AddMinutes(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20").AddMinutes(-1), "2022-05-19 23:59:00")
	assert(Parse("2022-05-20").AddMinutes(-60), "2022-05-19 23:00:00")
	assert(Parse("2022-05-20").AddMinutes(1), "2022-05-20 00:01:00")
	assert(Parse("2022-05-20").AddMinutes(60), "2022-05-20 01:00:00")
}

func TestTyme_AddHours(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20").AddHours(1), "2022-05-20 01:00:00")
	assert(Parse("2022-05-20").AddHours(24), "2022-05-21 00:00:00")
	assert(Parse("2022-05-20").AddHours(-1), "2022-05-19 23:00:00")
	assert(Parse("2022-05-20").AddHours(-24), "2022-05-19 00:00:00")
}

func TestTyme_AddDays(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20").AddDays(1), "2022-05-21 00:00:00")
	assert(Parse("2022-05-20").AddDays(24), "2022-06-13 00:00:00")
	assert(Parse("2022-05-20").AddDays(-1), "2022-05-19 00:00:00")
	assert(Parse("2022-05-20").AddDays(-24), "2022-04-26 00:00:00")
}

func TestTyme_AddWeeks(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20").AddWeeks(1), "2022-05-27 00:00:00")
	assert(Parse("2022-05-20").AddWeeks(-1), "2022-05-13 00:00:00")
}

func TestTyme_AddMonths(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20").AddMonths(1), "2022-06-20 00:00:00")
	assert(Parse("2022-05-20").AddMonths(8), "2023-01-20 00:00:00")
	assert(Parse("2022-05-20").AddMonths(-8), "2021-09-20 00:00:00")
}

func TestTyme_AddYears(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-20").AddYears(1), "2023-05-20 00:00:00")
	assert(Parse("2022-05-20").AddYears(-1), "2021-05-20 00:00:00")
}

func TestTyme_Monday(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-15").Monday(), "2022-05-09 00:00:00")
	assert(Parse("2022-05-16").Monday(), "2022-05-16 00:00:00")
	SetFirstDayOfWeek(time.Sunday)
	assert(Parse("2022-05-15").Monday(), "2022-05-16 00:00:00")
	assert(Parse("2022-05-16").Monday(), "2022-05-16 00:00:00")
	SetFirstDayOfWeek(time.Monday)
}

func TestTyme_Tuesday(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-15").Tuesday(), "2022-05-10 00:00:00")
	assert(Parse("2022-05-16").Tuesday(), "2022-05-17 00:00:00")
	SetFirstDayOfWeek(time.Sunday)
	assert(Parse("2022-05-15").Tuesday(), "2022-05-17 00:00:00")
	assert(Parse("2022-05-16").Tuesday(), "2022-05-17 00:00:00")
	SetFirstDayOfWeek(time.Monday)
}

func TestTyme_Wendesday(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-15").Wendesday(), "2022-05-11 00:00:00")
	assert(Parse("2022-05-16").Wendesday(), "2022-05-18 00:00:00")
	SetFirstDayOfWeek(time.Sunday)
	assert(Parse("2022-05-15").Wendesday(), "2022-05-18 00:00:00")
	assert(Parse("2022-05-16").Wendesday(), "2022-05-18 00:00:00")
	SetFirstDayOfWeek(time.Monday)
}

func TestTyme_Thursday(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-15").Thursday(), "2022-05-12 00:00:00")
	assert(Parse("2022-05-16").Thursday(), "2022-05-19 00:00:00")
	SetFirstDayOfWeek(time.Sunday)
	assert(Parse("2022-05-15").Thursday(), "2022-05-19 00:00:00")
	assert(Parse("2022-05-16").Thursday(), "2022-05-19 00:00:00")
	SetFirstDayOfWeek(time.Monday)
}

func TestTyme_Friday(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-15").Friday(), "2022-05-13 00:00:00")
	assert(Parse("2022-05-16").Friday(), "2022-05-20 00:00:00")
	SetFirstDayOfWeek(time.Sunday)
	assert(Parse("2022-05-15").Friday(), "2022-05-20 00:00:00")
	assert(Parse("2022-05-16").Friday(), "2022-05-20 00:00:00")
	SetFirstDayOfWeek(time.Monday)
}

func TestTyme_Saturday(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-15").Saturday(), "2022-05-14 00:00:00")
	assert(Parse("2022-05-16").Saturday(), "2022-05-21 00:00:00")
	SetFirstDayOfWeek(time.Sunday)
	assert(Parse("2022-05-15").Saturday(), "2022-05-21 00:00:00")
	assert(Parse("2022-05-16").Saturday(), "2022-05-21 00:00:00")
	SetFirstDayOfWeek(time.Monday)
}

func TestTyme_Sunday(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022-05-15").Sunday(), "2022-05-15 00:00:00")
	assert(Parse("2022-05-16").Sunday(), "2022-05-22 00:00:00")
	SetFirstDayOfWeek(time.Sunday)
	assert(Parse("2022-05-15").Sunday(), "2022-05-15 00:00:00")
	assert(Parse("2022-05-16").Sunday(), "2022-05-15 00:00:00")
	SetFirstDayOfWeek(time.Monday)
}
