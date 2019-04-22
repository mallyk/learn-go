package main

import "fmt"

func main() {
	x := 42
	if x == 40 {
		fmt.Println("value is 40")
	} else if x == 41 {
		fmt.Println("value is 41")
	} else {
		fmt.Println("its neither 40 or 41")
	}
}
