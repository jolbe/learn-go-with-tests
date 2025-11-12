// Package concurrency is checking websites availability
package concurrency

import (
	"sync"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	ch := make(chan result)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Go(func() {
			ch <- result{url, wc(url)}
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		results[v.string] = v.bool
	}

	return results
}

func CheckWebsitesConcurrent(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	ch := make(chan result)

	for _, url := range urls {
		go func() {
			ch <- result{url, wc(url)}
		}()
	}

	for range urls {
		res := <-ch
		results[res.string] = res.bool
	}

	return results
}

func CheckWebsitesMutex(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			results[url] = wc(url)
		})
		// wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	mu.Lock()
		// 	results[url] = wc(url)
		// 	mu.Unlock()
		// }()
	}

	wg.Wait()
	return results
}

func CheckWebsitesSequential(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}

