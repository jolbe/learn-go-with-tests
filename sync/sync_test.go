package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})

	t.Run("it runs safely concurrently with reads", func(t *testing.T) {
		counter := NewCounter()
		var wg sync.WaitGroup

		// spawn writers
		for range 100 {
			wg.Go(func() {
				for range 10 {
					counter.Inc()
				}
			})
		}

		// spawn readers (while writers are still running)
		for range 1000 {
			wg.Go(func() {
				_ = counter.Value()
			})
		}

		wg.Wait()

		assertCount(t, counter, 1000)
	})
}

func assertCount(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d; want %d", got.Value(), want)
	}
}
