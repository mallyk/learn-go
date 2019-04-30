package main

import "fmt"

type person struct {
	first            string
	last             string
	favoriteIceCream []string
}

func main() {

	p1 := person{
		first:            "James",
		last:             "Bond",
		favoriteIceCream: []string{"rocky road", "chocolate", "chunky monkey"},
	}

	p2 := person{
		first:            "Miss",
		last:             "Moneypenny",
		favoriteIceCream: []string{"butterscotch", "monster cookie"},
	}

	printPerson(p1)
	printPerson(p2)

}

func printPerson(p person) {
	fmt.Println(p.first)
	fmt.Println(p.last)
	for k, v := range p.favoriteIceCream {
		fmt.Println("fav ice cream", k, ":", v)
	}
}
