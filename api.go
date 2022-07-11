package tyme

import (
	"errors"
	"time"
)

var (
	globalLoc      *time.Location = time.UTC
	globalDLayout  string         = "2006-01-02"
	globalTLayout  string         = "15:04:05"
	globalDTLayout string         = "2006-01-02 15:04:05"
	globalFormats                 = []string{
		"2006", "2006-01", "2006-01-02", "2006-01-02 15", "2006-01-02 15:4", "2006-01-02 15:04:05", "01-02",
		"2006.01", "2006.01.02", "2006.01.02 15", "2006.01.02 15:4", "2006.01.02 15:04:05", "01.02",
		"2006/01", "2006/01/02", "2006/01/02 15", "2006/01/02 15:4", "2006/01/02 15:04:05", "01/02",
		"15:04:05", "15:04", "15",
		"15:04:05 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02T15:04:05Z0700", "2006-01-02T15:04:05Z07",
		"01/02/2006", "01/02/2006 15:04:05", "2006/01/02", "20060102", "2006/01/02 15:04:05",
		time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850,
		time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
		time.Kitchen, time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
	}

	firstDayOfWeek = time.Monday
)

// ==> Global settings
// SetDTLayout 设置全局Datetime layout
func SetDTLayout(format string) {
	globalDTLayout = format
}

// SetDLayout 设置全局Date layout
func SetDLayout(format string) {
	globalDLayout = format
}

// SetTLayout 设置全局Time layout
func SetTLayout(format string) {
	globalTLayout = format
}

// SetTimezone 设置全局时区
func SetTimezone(timezone string) error {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	globalLoc = loc
	return nil
}

// AppendFormats 追加时间格式
func AppendFormats(formats ...string) {
	globalFormats = append(globalFormats, formats...)
}

// SetFirstDayOfWeek 设置一周以周天或周一开始
func SetFirstDayOfWeek(weekday time.Weekday) {
	if weekday != time.Monday && weekday != time.Sunday {
		return
	}

	firstDayOfWeek = weekday
}

// ==> Specific time

// With 将time.Time对象包装为*Tyme
func With(t time.Time) *Tyme {
	return &Tyme{Time: t}
}

// Parse 解析日期字符串为*Tyme
// NOTE: 解析失败，返回zero value
func Parse(s string) *Tyme {
	t, err := StrictParse(s)
	if err != nil {
		return &Tyme{Time: time.Time{}}
	}

	return t
}

// StrictParse 解析日期字符串为*Tyme
func StrictParse(s string) (*Tyme, error) {
	for _, format := range globalFormats {
		ot, err := time.ParseInLocation(format, s, globalLoc)

		if err == nil {
			t := &Tyme{Time: ot}
			return t, nil
		}
	}

	return nil, errors.New("Could not parse string as time: " + s)
}

// MustParse 解析日期字符串为*Tyme
// NOTE: 解析失败，panic
func MustParse(s string) *Tyme {
	t, err := StrictParse(s)
	if err != nil {
		panic(err)
	}

	return t
}

// ==> Basic functions

// Now 返回当前时间
func Now() *Tyme {
	return With(time.Now())
}

// ==> Weekdays

// Monday 返回当日周一日期
func Monday() *Tyme {
	return Now().Monday()
}

// Tuesday 返回当日周二日期
func Tuesday() *Tyme {
	return Now().Tuesday()
}

// Wendesday 返回当日周三日期
func Wendesday() *Tyme {
	return Now().Wendesday()
}

// Thursday 返回当日周四日期
func Thursday() *Tyme {
	return Now().Thursday()
}

// Friday 返回当日周五日期
func Friday() *Tyme {
	return Now().Friday()
}

// Saturday 返回当日周六日期
func Saturday() *Tyme {
	return Now().Saturday()
}

// Sunday 返回当日周日日期
func Sunday() *Tyme {
	return Now().Sunday()
}

// ==> Relative days

// Today returns the today's date
func Today() *Tyme {
	return Now().BeginOfDay()
}

// Yesterday returns the yesterday's date
func Yesterday() *Tyme {
	return Today().AddDays(-1)
}

// Tomorrow returns the tomorrow's date
func Tomorrow() *Tyme {
	return Today().AddDays(1)
}
