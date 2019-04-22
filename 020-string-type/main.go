package main

import "fmt"

func main() {
	s := "Hello"          //string
	t := `string literal` //represented by back ticks
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Println(t)
	fmt.Printf("%T\n", t)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println()
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

	bt := []byte(t)
	fmt.Println(bt)
	fmt.Printf("%T\n", bt)

	for i := 0; i < len(t); i++ {
		fmt.Printf("%#U", t[i])
	}
}
