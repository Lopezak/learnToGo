package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	m1 := m["James"]
	fmt.Println(m1)

	m2 := m["Q"]
	fmt.Println(m2)

	if _, ok := m["Q"]; !ok {
		fmt.Println("No Q")
	}
}
