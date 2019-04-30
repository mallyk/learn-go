package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	v1 := truck{
		fourWheel: true,
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
	}

	v2 := sedan{
		luxury: false,
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.color)
	fmt.Println(v2.color)
}
