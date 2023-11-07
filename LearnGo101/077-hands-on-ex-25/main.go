package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 0
	for i < 25 {
		r := rand.Intn(300)
		fmt.Printf("%d.\t%d\n", i, r)
		switch {
		case 0 <= r && r <= 100:
			fmt.Println("\tBetween 0 and 100")
		case 101 <= r && r <= 200:
			fmt.Println("\tBetween 101 and 200")
		case 201 <= r && r <= 250:
			fmt.Println("\tBetween 201 and 250")
		default:
			fmt.Println("\tNumber is greater than 250")
		}
		i++
	}
}
