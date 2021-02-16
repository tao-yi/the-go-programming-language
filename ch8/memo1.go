package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache from race condition
	cache map[string]*entry
}

// Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		// this is a repeat request for this key
		memo.mu.Unlock()
		<-e.ready // wait for ready condition
	}

	return e.res.value, e.res.err
}

func httpGetbody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := New(httpGetbody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

var incomingURLsOnce sync.Once

func incomingURLs() []string {
	var urls []string
	incomingURLsOnce.Do(func() {
		urls = append(urls, "https://www.baidu.com")
		urls = append(urls, "https://www.baidu.com")
		urls = append(urls, "https://www.bilibili.com")
	})
	return urls
}
