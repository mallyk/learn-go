package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\t\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("The inner loop: %d\n", j)
		}
	}
}

func conditionOnly() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
}

func justAFor() {
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
}
