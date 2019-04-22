package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print 2")
	case 3 == 3:
		fmt.Println("prints")
	case 4 == 4:
		fmt.Println("no default fall through")
	}
}

func fallThroughCase() {
	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print 2")
	case 3 == 3:
		fmt.Println("prints")
		fallthrough //don't use
	case 4 == 4:
		fmt.Println("no default fall through")
	}
}

func defaultCause() {
	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print 2")
	case 3 == 3:
		fmt.Println("prints")
	case 4 == 4:
		fmt.Println("no default fall through")
	default:
		fmt.Println("default")
	}
}
