// Package racer shows how to use select on multiple channels
package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

// RacerSynchronous case
func RacerSynchronous(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// RacerChannels case
func RacerChannels(a, b string) (winner string) {
	chA := make(chan time.Duration)
	chB := make(chan time.Duration)
	go measureResponseTimeChannel(a, chA)
	go measureResponseTimeChannel(b, chB)
	aDuration := <-chA
	bDuration := <-chB

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTimeChannel(url string, ch chan time.Duration) {
	start := time.Now()
	http.Get(url)
	ch <- time.Since(start)
}

// RacerGoroutines case
func RacerGoroutines(a, b string) (winner string) {
	ch := make(chan string)

	go func() {
		http.Get(a)
		ch <- a
	}()
	go func() {
		http.Get(b)
		ch <- b
	}()

	return <-ch
}
