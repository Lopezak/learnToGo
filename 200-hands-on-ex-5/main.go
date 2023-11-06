package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type SortByLast []user

func (sbl SortByLast) Len() int           { return len(sbl) }
func (sbl SortByLast) Swap(i, j int)      { sbl[i], sbl[j] = sbl[j], sbl[i] }
func (sbl SortByLast) Less(i, j int) bool { return sbl[i].Last < sbl[j].Last }

type SortByAge []user

func (sba SortByAge) Len() int           { return len(sba) }
func (sba SortByAge) Swap(i, j int)      { sba[i], sba[j] = sba[j], sba[i] }
func (sba SortByAge) Less(i, j int) bool { return sba[i].Age < sba[j].Age }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	sort.Sort(SortByAge(users))
	fmt.Println(users)
	sort.Sort(SortByLast(users))
	fmt.Println(users)
}
