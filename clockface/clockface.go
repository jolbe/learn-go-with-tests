// Package clockface provides functions that calculate the position of the hour hands.
// of an analogue clock,.
package clockface

import (
	"math"
	"time"
)

const (
	secondsPerMinute = 60
	minutesPerHour   = 60
	hoursPerClock    = 12
)

const (
	secondHandRadiansPerSecond = 2 * math.Pi / secondsPerMinute
	minuteHandRadiansPerSecond = 2 * math.Pi / (secondsPerMinute * minutesPerHour)
	hourHandRadiansPerSecond   = 2 * math.Pi / (secondsPerMinute * minutesPerHour * hoursPerClock)
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X, Y float64
}

func SecondsInRadians(t time.Time) float64 {
	return float64(t.Second()) * secondHandRadiansPerSecond
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

func MinutesInRadians(t time.Time) float64 {
	elapsedSeconds := t.Minute()*secondsPerMinute + t.Second()
	return float64(elapsedSeconds) * minuteHandRadiansPerSecond
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

func HoursInRadians(t time.Time) float64 {
	elapsedSeconds := (t.Hour()%hoursPerClock)*minutesPerHour*secondsPerMinute +
		t.Minute()*secondsPerMinute +
		t.Second()
	return float64(elapsedSeconds) * hourHandRadiansPerSecond
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
