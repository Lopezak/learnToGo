package main

import "fmt"

func main() {
	xs := make([]int, 0, 10)

	//fmt.Println(cap(xs))
	//fmt.Println(len(xs))

	xs = append(xs, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51)

	fmt.Println(xs)
	fmt.Println(xs[:5])
	fmt.Println(xs[5:])
	fmt.Println(xs[2:7])
	fmt.Println(xs[1:6])
	/*
		for i, v := range xs {
			fmt.Printf("Iteration %v is %v with type %T\n", i, v, v)
		}
	*/
}
