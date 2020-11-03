package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// 通过唯一通道来修改缓存，所有并发的goroutine都把请求发送给这个唯一通道

type entry struct {
	res   result
	ready chan struct{}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(resp chan result) {
	<-e.ready
	resp <- e.res
}

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan result
}

type Func func(key string) (interface{}, error)

type Memo struct {
	requests chan request
}

func (memo *Memo) Get(key string) (interface{}, error) {
	request := &request{key: key, response: make(chan result)}
	memo.requests <- *request
	res := <-request.response
	return res.value, res.err
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.Server(f)
	return memo
}

func (memo *Memo) Server(f Func) {
	cache := make(map[string]*entry)

	for request := range memo.requests {
		e := cache[request.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[request.key] = e
			go e.call(f, request.key)
		}
		go e.deliver(request.response)
	}
}

func getBody(key string) (interface{}, error) {
	res, err := http.Get(key)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body), nil
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
	memo := New(getBody)
	var n sync.WaitGroup
	for _, url := range incommingUrls {
		n.Add(1)
		go func(key string) {
			value, err := memo.Get(key)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(value)
			}
			n.Done()
		}(url)
	}
	n.Wait()
}
