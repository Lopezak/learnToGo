package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("***This is where initialization for my program occurs***")
}

func main() {

	fmt.Println("This is the ifVersion")
	for i := 0; i < 20; i++ {
		ifVersion()
	}

	fmt.Println("This is the switchVersion")
	for i := 0; i < 20; i++ {
		switchVersion()
	}
}

func ifVersion() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("Y = %v\n", y)
	fmt.Printf("X = %v\n", x)

	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater than or queal to 4 and less than or equal to 6")
	} else if y != 5 {
		fmt.Println("Y is not 5")
	} else {
		fmt.Println("None of the previous cases were met")
	}
}

func switchVersion() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("Y = %v\n", y)
	fmt.Printf("X = %v\n", x)

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or queal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println("Y is not 5")
	default:
		fmt.Println("None of the previous cases were met")
	}
}
