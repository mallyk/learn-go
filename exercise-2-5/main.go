package main

import "fmt"

func main() {
	x := `raw string 
	with
	junk
	"in between"
	literal`
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
