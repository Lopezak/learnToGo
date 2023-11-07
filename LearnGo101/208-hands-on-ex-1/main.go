package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go foo()
	wg.Add(1)
	go bar()

	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("From the routine 1", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("From the routine 2", i)
	}
	wg.Done()
}
