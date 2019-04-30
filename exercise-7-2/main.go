package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	j := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println("before changeMe", j)
	fmt.Println(&j)

	changeMe(&j)
	fmt.Println("after changeMe", j)
	fmt.Println(&j)
}

func changeMe(p *person) {
	(*p).first = "Miss"
	(*p).last = "Moneypenny"
	(*p).age = 27

	//could also do p.first...
}
