package tyme

import "errors"

type SeriesStep = int

const (
	Millisecond = iota
	Second
	Minute
	Hour
	Day
	Week
	Month
	Quarter
	Year
)

type series[T any] struct {
	Begin *Tyme
	End   *Tyme

	StepType SeriesStep
	Step     uint
}

func Series[T any](begin, end *Tyme, stepType SeriesStep, step uint) *series[T] {
	return &series[T]{
		Begin:    begin,
		End:      end,
		StepType: stepType,
		Step:     step,
	}
}

func (s *series[T]) As(iteratee func(*Tyme) (T, error)) ([]T, error) {
	add, err := s.addFn()
	if err != nil {
		return nil, err
	}

	var res []T
	for t := s.Begin; t.Before(s.End); t = add(t, int(s.Step)) {
		item, err := iteratee(t)
		if err != nil {
			return nil, err
		}
		res = append(res, item)
	}
	return res, nil
}

func (s *series[T]) addFn() (func(*Tyme, int) *Tyme, error) {
	switch s.StepType {
	case Millisecond:
		return func(t *Tyme, n int) *Tyme { return t.AddMilliseconds(n) }, nil
	case Second:
		return func(t *Tyme, n int) *Tyme { return t.AddSeconds(n) }, nil
	case Minute:
		return func(t *Tyme, n int) *Tyme { return t.AddMinutes(n) }, nil
	case Hour:
		return func(t *Tyme, n int) *Tyme { return t.AddHours(n) }, nil
	case Day:
		return func(t *Tyme, n int) *Tyme { return t.AddDays(n) }, nil
	case Week:
		return func(t *Tyme, n int) *Tyme { return t.AddWeeks(n) }, nil
	case Month:
		return func(t *Tyme, n int) *Tyme { return t.AddMonths(n) }, nil
	case Quarter:
		return func(t *Tyme, n int) *Tyme { return t.AddMonths(n * 3) }, nil
	case Year:
		return func(t *Tyme, n int) *Tyme { return t.AddYears(n) }, nil
	default:
		return nil, errors.New("Invalid step type")
	}
}
