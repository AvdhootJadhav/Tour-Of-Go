package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

// Concurrency Exercise for Channels

func Walk(tree *tree.Tree, ch chan int) {
	if tree == nil {
		return
	}
	Walk(tree.Left, ch)
	ch <- tree.Value
	Walk(tree.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {

	first := make(chan int)
	second := make(chan int)

	go Walk(t1, first)
	go Walk(t2, second)

	for i := 1; i <= 10; i++ {
		v1 := <-first
		v2 := <-second

		if v1 != v2 {
			return false
		}
	}
	return true
}

func channels_exercise() {
	fmt.Println("Both trees are same? :", Same(tree.New(1), tree.New(2)))
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Cache struct {
	store map[string]bool
	lock  sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *Cache) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	cache.lock.Lock()
	check := cache.store[url]
	cache.lock.Unlock()
	if check {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		cache.lock.Lock()
		cache.store[url] = true
		cache.lock.Unlock()
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	cache.lock.Lock()
	cache.store[url] = true
	cache.lock.Unlock()
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, cache)
	}
	time.Sleep(time.Second * 1)
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

// fetcher is a populated fakeFetcher.
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

func main() {
	t := time.Now()
	cache := Cache{
		store: map[string]bool{},
		lock:  sync.Mutex{},
	}
	Crawl("https://golang.org/", 4, fetcher, &cache)
	fmt.Println("Time taken : ", time.Since(t))
}
