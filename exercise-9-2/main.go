package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am", p.age, "years old.")
}

type human interface {
	speak()
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println("hello")
	saySomething(&p1)
}

func saySomething(h human) {
	h.speak()
}
