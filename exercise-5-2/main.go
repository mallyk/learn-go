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

	people := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	printPeople(people)

}

func printPeople(people map[string]person) {

	for k, v := range people {
		fmt.Println("person", k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, iceCream := range v.favoriteIceCream {
			fmt.Println("fav ice cream", i, ":", iceCream)
		}

	}
}
