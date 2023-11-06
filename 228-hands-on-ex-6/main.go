package main

import "fmt"

func main() {
	c := make(chan int)
	addVals(c)
	receive(c)

	fmt.Println("Exiting program")
}

func addVals(c chan int) <-chan int {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
