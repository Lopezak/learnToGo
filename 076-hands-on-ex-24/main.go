package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 0
	for i < 25 {
		r := rand.Intn(250)
		fmt.Printf("%d.\t%d\n", i, r)
		if 0 <= r && r <= 100 {
			fmt.Println("\tBetween 0 and 100")
		} else if 101 <= r && r <= 200 {
			fmt.Println("\tBetween 101 and 200")
		} else {
			fmt.Println("\tBetween 201 and 250")
		}
		i++
	}
}
