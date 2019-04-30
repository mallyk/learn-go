package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("my name is", p.first, p.last, "and I am", p.age)
}

func main() {
	fmt.Println("hello")
	p := person{
		first: "Dominic",
		last:  "Torreto",
		age:   35,
	}

	p.speak()
}
