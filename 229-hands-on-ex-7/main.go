package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	go launch10(c)

	i := 0
	for v := range c {
		i++
		fmt.Println(i, v)
	}

	fmt.Println("Exiting")
}

func launch10(c chan int) {
	var wg sync.WaitGroup
	for i := 1; i < 11; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			wg.Done()
		}()
	}
	// go func() {
	// 	wg.Wait()
	// 	close(c)
	// }()
	wg.Wait()
	close(c)
}
