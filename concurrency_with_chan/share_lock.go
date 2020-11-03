package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// 通过使用共享锁来守护缓存的并发设计模型
// 所有的goroutine同时抢锁，谁抢到谁注册缓存，其它相同请求的goroutine
// 只需要等待注册的entry ready就好了, 不必重复发请求

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type Func func(key string) (interface{}, error)

func getBody(key string) (interface{}, error) {
	res, err := http.Get(key)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func (m *Memo) Get(key string) (value interface{}, err error) {
	m.mu.Lock()
	e := m.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mu.Unlock()

		e.res.value, e.res.err = m.f(key)
		close(e.ready)
	} else {
		m.mu.Unlock()
		<-e.ready
	}

	return e.res.value, e.res.err
}

func main() {

	incommingUrls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://golang.org",
		"http://gopl.io",
		"https://godoc.org",
		"http://gopl.io",
	}

	m := New(getBody)

	var n sync.WaitGroup
	for _, url := range incommingUrls {
		n.Add(1)
		go func(url string) {
			value, err := m.Get(url)
			if err != nil {
				log.Println(err.Error())
			} else {
				log.Println(value.([]byte))
			}
			n.Done()
		}(url)
	}
	n.Wait()
}
