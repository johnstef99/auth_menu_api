package main

import "time"

type Weekday int

const (
	_ Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func isSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}

func getTodayWeekday() Weekday {
	t := time.Now()
	var w Weekday
	if tw := t.Weekday(); tw == 0 {
		w = Weekday(Sunday)
	} else {
		w = Weekday(tw)
	}
	return w
}
