package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int64
	color string
}

func main() {
	car1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Hyundai",
		model: "Sonata",
		doors: 4,
		color: "Gray",
	}

	fmt.Println(car1)
	fmt.Println(car1.model, car1.electric)

	car2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Ford",
		model: "F150",
		doors: 4,
		color: "White",
	}

	fmt.Println(car2)
	fmt.Println(car2.model, car2.electric)
}
