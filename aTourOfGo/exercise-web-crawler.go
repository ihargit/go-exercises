package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeUrlMap struct {
	mu sync.Mutex
	v  map[string]string
}

func (c *SafeUrlMap) Set(key string, body string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key] = body
}

func (c *SafeUrlMap) Value(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	body, ok := c.v[key]
	return body, ok
}

func Crawl(url string, depth int, fetcher Fetcher, sm SafeUrlMap) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	sm.Set(url, body)
	//fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if _, ok :=sm.Value(u); !ok {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, sm)
		}
	}
	return
}

var wg sync.WaitGroup

func main() {
	safeMap := SafeUrlMap{v: make(map[string]string)}
	
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, safeMap)
	wg.Wait()
	
	for url := range safeMap.v {
		body, _ := safeMap.Value(url)
		fmt.Printf("found: %s %q\n", url, body)
	}
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
