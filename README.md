# Tyme
Useful time utils for golang.
## Features
### Get Tyme instance
- Getting *Tyme from time.Time
```go
// Get *Tyme from time.Time
t := With(time.Now())
```
- 3 methods for getting *Tyme from a time string
```go
// return zero value if parse failed
t := Parse("2022-05-22")
// returns error if parse failed
t, err := StrictParse("2022.05.22")
// panic if parse failed
t := MustParse("2022/05/22")
```
- Settings global formats to parse different string of time
> There are some built-in formats, and also yout can append some custom formats:
```go
AppendFormats("2006-01-02 15:04:05", "...")
```

### Convert to string
- Override the String interface
```go
Now().String() // use the default layout: 2006-01-02 15:04:05
```
- Layout
| You can also use layout functions to get date, time string.
```go
SetDTLayout("2016.01.02 15:04:05") // set datetime layout
SetDayout("2016.01.02 15:04:05") // set date layout
SetTayout("2016.01.02 15:04:05") // set time layout

Now().DateLayout() // 2022.05.20
Now().TimeLayout() // 16:14:53
```

### Comparison
- Comparison between Tyme objects:
```go
t1 := Parse("2022-05-20 16:14:53")
t2 := Parse("2022-05-20 16:14:55")
t3 := Parse("2022-05-20 16:14:54")

t1.Equal(t2) // false
t1.Before(t2) // true
t1.After(t2) // false
t3.Between(t1, t2) // true
```
- Comparison with time string:
> **NOTE:** All of below functions is not safely! If tyme cannot parse the string, it will panic.
```go
t1 := Now()

t1.EqualS("2022-05-20")
t1.BeforeS("2022.05.20")
t1.AfterS("2022/05/20")
t1.BetweenS("2022-05-20", "2022-05-22")
```

### Time calculations
Increase or decrease time according to different time units:
```go
t1 := Parse("2022-05-20 16:14:53")

t := t1.AddYears(1).AddMonths(-1).AddSeconds(-1)
```

### Get meaningful days easily
```go
today := Today() // the date of today
yesterday := Yesterday() // the date of yesterday
tomorrow = Tomorrow() // the date of tomorrow
monday = Monday() // the date of the Monday of current week
tuesday = TuesDay() // the date of the TuesDay of current week
// and so on ... 
```

### Get a specific point in time
```go
t := Now()

t.BeginOfMinute()
t.BeginOfHour()
t.BeginOfDay()
t.BeginOfWeek()
t.BeginOfMonth()
t.BeginOfQuarter()
t.BeginOfYear()

t.EndOfMinute()
t.EndOfHour()
t.EndOfDay()
t.EndOfWeek()
t.EndOfMonth()
t.EndOfQuarter()
t.EndOfYear()
```

### Series
```go
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
```
