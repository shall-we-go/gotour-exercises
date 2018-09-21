/*
https://tour.golang.org/concurrency/10
*/
package main

import (
	"fmt"
	"sync"
)

var doneUrls = make(map[string]bool)
var mux sync.Mutex

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, quit chan bool) {
	mux.Lock()
	if doneUrls[url] {
		mux.Unlock()
		quit <- true
		return
	}
	doneUrls[url] = true
	mux.Unlock()
	if depth <= 0 {
		quit <- true
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		quit <- true
		return
	}
	fmt.Printf("found: %s %q %d\n", url, body, len(urls))
	childQuit := make(chan bool, len(urls))
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, childQuit)
	}
	for i := 0; i < len(urls); i++ {
		<-childQuit
	}
	quit <- true
	return
}

func main() {
	quit := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, quit)
	<-quit
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (this fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := this[url]; ok {
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
