package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for key, value := range xi {
		fmt.Printf("The value is %v and the index is %v\n", value, key)
	}
}
