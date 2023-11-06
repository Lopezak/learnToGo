package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var counter int

var mu sync.Mutex

func main() {
	for i := 0; i < 100; i++ {
		go incrementer()
	}
	wg.Wait()
	fmt.Println("Final counter = ", counter)
}

func incrementer() {
	wg.Add(1)
	mu.Lock()
	c := counter
	//fmt.Println("Counter = ", c)
	//runtime.Gosched()
	c++
	counter = c
	fmt.Println("Go routines: ", runtime.NumGoroutine())
	fmt.Println("Counter = ", counter)
	mu.Unlock()
	wg.Done()
}
