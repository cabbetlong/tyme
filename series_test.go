package tyme

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSeries_TestAs(t *testing.T) {
	t1 := With(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC))
	t2 := With(time.Date(2022, 1, 7, 0, 0, 0, 0, time.UTC))
	series, _ := Series[string](t1, t2, Day, 1).As(func(t *Tyme) (string, error) {
		return t.DateLayout(), nil
	})

	assert.Equal(t, []string{
		"2022-01-01",
		"2022-01-02",
		"2022-01-03",
		"2022-01-04",
		"2022-01-05",
		"2022-01-06"}, series)
}
