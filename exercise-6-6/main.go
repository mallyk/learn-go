package main

import "fmt"

func main() {
	func() {
		fmt.Println("woo anonymous func.")
	}()
}
