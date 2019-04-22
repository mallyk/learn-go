package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barney"])

	v, ok := m["Barney"]
	fmt.Println(v)
	fmt.Println(ok)

	//check if value exists and print it
	if v, ok := m["James"]; ok {
		fmt.Println("we found it", v)
	}

	//add element to map
	m["Todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	//delete value from a map
	delete(m, "Todd")
	fmt.Println(m)

	//delete value from map with ok
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println(v)
		delete(m, "Miss Moneypenny")
		fmt.Println(m)
	}
}
