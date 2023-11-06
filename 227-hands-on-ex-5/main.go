package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// v, ok := <-c
	// fmt.Println(v, ok)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Println(v, ok)
	}
}
