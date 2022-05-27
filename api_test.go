package tyme

import (
	"testing"
)

func assertT(t *testing.T) func(*Tyme, string) {
	SetDTLayout("2006-01-02 15:04:05.999999999")
	return func(tyme *Tyme, expected string) {
		actualStr := tyme.String()
		if actualStr != expected {
			t.Errorf("actual: %v, expected: %v", tyme, expected)
		}
	}
}

func assertComparasion(t *testing.T) func(bool, bool) {
	return func(actual bool, expected bool) {
		if actual != expected {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}
	}
}

func TestParse(t *testing.T) {
	assert := assertT(t)

	assert(Parse("2022"), "2022-01-01 00:00:00")
	assert(Parse("2022-02"), "2022-02-01 00:00:00")
	assert(Parse("2022.02"), "2022-02-01 00:00:00")
	assert(Parse("2022/02"), "2022-02-01 00:00:00")
	assert(Parse("2022-02-21"), "2022-02-21 00:00:00")
	assert(Parse("2022.02.21"), "2022-02-21 00:00:00")
	assert(Parse("2022/02/21"), "2022-02-21 00:00:00")
	assert(Parse("2022-02-21 23"), "2022-02-21 23:00:00")
	assert(Parse("2022.02.21 23"), "2022-02-21 23:00:00")
	assert(Parse("2022/02/21 23"), "2022-02-21 23:00:00")
	assert(Parse("2022-02-21 23:02"), "2022-02-21 23:02:00")
	assert(Parse("2022.02.21 23:02"), "2022-02-21 23:02:00")
	assert(Parse("2022/02/21 23:02"), "2022-02-21 23:02:00")
	assert(Parse("2022-02-21 23:02:59"), "2022-02-21 23:02:59")
	assert(Parse("2022.02.21 23:02:59"), "2022-02-21 23:02:59")
	assert(Parse("2022/02/21 23:02:59"), "2022-02-21 23:02:59")
}

func TestToday(t *testing.T) {
	assert := assertT(t)

	assert(Today(), "2022-05-27 00:00:00")
}
