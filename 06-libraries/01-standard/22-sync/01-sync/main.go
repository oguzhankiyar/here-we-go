package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	Sample("Cond", Cond)
	Sample("Map", Map)
	Sample("Mutex", Mutex)
	Sample("RWMutex", RWMutex)
	Sample("Once", Once)
	Sample("Pool", Pool)
	Sample("WaitGroup", WaitGroup)
}

func Cond() {
	data := make(map[string]string)

	wg := sync.WaitGroup{}
	wg.Add(2)

	m := sync.Mutex{}
	c := sync.NewCond(&m)

	go func() {
		c.L.Lock()
		for len(data) == 0 {
			c.Wait()
		}
		fmt.Println("one", data["one"])
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		c.L.Lock()
		for len(data) == 0 {
			c.Wait()
		}
		fmt.Println("two", data["two"])
		c.L.Unlock()
		wg.Done()
	}()

	c.L.Lock()
	data["one"] = "foo"
	data["two"] = "bar"
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}

func Map() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	m := sync.Map{}

	for i := 1; i <= 10; i++ {
		go func(i int) {
			m.Store(i, i)

			wg.Done()
		}(i)
	}

	wg.Wait()

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%v:%v ", key, value)
		return true
	})

	fmt.Println()
}

func Mutex() {
	m := sync.Mutex{}

	var counter int

	fn := func() {
		m.Lock()
		counter++
		fmt.Printf("%v ", counter)
		time.Sleep(100 * time.Millisecond)
		m.Unlock()
	}

	for i := 1; i <= 10; i++ {
		go fn()
	}

	time.Sleep(1200 * time.Millisecond)

	fmt.Println()
}

func RWMutex() {
	wg := sync.WaitGroup{}
	wg.Add(8)

	m := sync.RWMutex{}

	counter := 0

	read := func() {
		m.RLock()

		fmt.Println("read:", counter)

		m.RUnlock()
		wg.Done()
	}

	write := func() {
		m.Lock()

		counter++
		fmt.Println("write:", counter)

		m.Unlock()
		wg.Done()
	}

	for i := 1; i <= 4; i++ {
		go write()
	}

	for i := 1; i <= 4; i++ {
		go read()
	}

	wg.Wait()
}

func Once() {
	once := sync.Once{}

	for i := 1; i <= 10; i++ {
		once.Do(func() {
			fmt.Println("Hello")
		})
	}
}

func Pool() {
	type user struct {
		id int
		name string
	}

	pool := sync.Pool{
		New: func() interface{} {
			return &user{
				id: rand.Int(),
				name: "",
			}
		},
	}

	for i := 1; i <= 5; i++ {
		newUser := pool.Get().(*user)
		fmt.Println(newUser.id)
		pool.Put(newUser)
	}
}

func WaitGroup() {
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		time.Sleep(500 * time.Millisecond)
		wg.Done()
	}()

	fmt.Println("Waiting the wg")
	wg.Wait()
	fmt.Println("Done")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}