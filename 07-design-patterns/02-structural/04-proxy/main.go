package main

import (
	"fmt"
	"sync"
	"time"
)

type Requester interface {
	Request(string) string
}

type DefaultRequester struct {

}

func (r DefaultRequester) Request(url string) string {
	return fmt.Sprintf(`{ "url": "%s", "time": %v }`, url, time.Now().Unix())
}

type CachedRequester struct {
	requester	Requester
	cache		map[string]string
	mutex		sync.Mutex
}

func NewCachedRequester(requester Requester) *CachedRequester {
	return &CachedRequester{
		requester: requester,
		cache: make(map[string]string),
		mutex: sync.Mutex{},
	}
}

func (r CachedRequester) Request(url string) string {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if val, ok := r.cache[url]; ok {
		fmt.Println("fetched from cache")
		return val
	}

	val := r.requester.Request(url)
	r.cache[url] = val

	fmt.Println("stored to cache")

	return val
}


func main() {
	defaultRequester := DefaultRequester{}
	cachedRequester := NewCachedRequester(defaultRequester)

	fmt.Println(cachedRequester.Request("localhost"))
	time.Sleep(time.Second)
	fmt.Println(cachedRequester.Request("localhost"))
}