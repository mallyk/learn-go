package main

import "fmt"

func main() {
	favSport := "basketball"
	switch favSport {
	case "skiing":
		fmt.Println("won't print")
	case "surfing":
		fmt.Println("will print")
	case "basketball":
		fmt.Println("yarrrr")
	}
}
