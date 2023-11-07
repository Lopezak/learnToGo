package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		fmt.Printf("%v:\t%v\n", i+1, getRandintAndDescribe())
	}
}

func getRandintAndDescribe() (result string) {
	r := rand.Intn(5)

	switch r {
	case 0:
		return "the number is 0"
	case 1:
		return "the number is 1"
	case 2:
		return "the number is 2"
	case 3:
		return "the number is 3"
	case 4:
		return "the number is 4"
	default:
		return "the number is not in range [0,5)"
	}
}
