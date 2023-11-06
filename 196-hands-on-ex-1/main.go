package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error marshalling to JSON")
	} else {
		fmt.Println("JSON marshall successful: ")
		fmt.Println(string(bs))
	}

	var result []user

	err = json.Unmarshal(bs, &result)
	if err != nil {
		fmt.Println("Error unmarshalling from JSON")
	} else {
		fmt.Println("JSON unmarshall successful: ")
		fmt.Println(result)
	}

}
