package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	go launch10(c)
	/*
		The above function "launch10(c)" has to be a go routine in and of itself. This is due to the fact that we are unable to send
		values onto a channel until there is a receiver ready. Channels are a way of communication, not memory, and thus there
		is a need for a conversation. In a conversation there must be 2 parties present, unless we are leaving a message. Therefore,
		our options are to put the "launch10(c)" into a go routine such that we will hit the for range "receiver" following shortly OR
		to make a buffer large enough to hold all the data that we will be passing. Assuming that we do not know the size of the data that
		will be passed, it makes more sense to put "launch10(c)" into a go routine. Without the go routine, we are trying to pass data with
		no receiver ready and we deadlock waiting for a receiver. The go routine allows these functions to run concurrently using the first
		concurrent function to launch all 10 other concurrent "anonymous" functions and the "main" go routine operates the receiver
	*/

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
