package main

import "fmt"

func main() {
	xs := []string{
		"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)",
		"Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)",
		"Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)",
		"Coffee (GF)",
		"Cookies & Cream",
		"Fresh Basil (GF)",
		"Garden Mint Chip (GF)",
		"Lavender Lemon Honey (GF)",
		"Lemon Bar",
		"Madagascar Vanilla (GF)",
		"Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy Sunbutter Cinnamon (GF, V)",
		"Orange Cream (GF) ",
		"Peanut Butter Cookie Dough",
		"Raspberry Sorbet (GF, V)",
		"Salty Caramel (GF)",
		"Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)",
		"Vanilla Caramel Blondie",
		"Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)",
	}

	for i, flavor := range xs {
		fmt.Printf("Number %v: \t %v\n", i+1, flavor)
	}

	xs = append(xs, "Mint Mint Chocolate Chocolate Chip")

	fmt.Printf("Number %v: \t %v\n", len(xs), xs[len(xs)-1])

	xs = append(xs[:15], xs[16:]...)

	fmt.Print("***After delete***\n\n")

	for i, flavor := range xs {
		fmt.Printf("Number %v: \t %v\n", i+1, flavor)
	}

}
