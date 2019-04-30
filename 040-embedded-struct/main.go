package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

//we create VALUES of a certain TYPE that are stored in VARIABLES
//and those VARIABLES have IDENTIFIERS

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk) //inner type promoted to outer type
	fmt.Println(p2)
}