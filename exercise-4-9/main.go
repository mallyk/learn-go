package main

import "fmt"

func main() {
	bondsThings := []string{"shaken, not stirred", "martinis", "women"}
	moneypennyThings := []string{"James Bond", "literature", "computer science"}
	drNoThings := []string{"being evil", "ice cream", "sunsets"}

	x := map[string][]string{
		"bond_james":      bondsThings,
		"moneypenny_miss": moneypennyThings,
		"no_dr":           drNoThings,
	}

	for k, v := range x {
		fmt.Println(k)
		for i, z := range v {
			fmt.Println(i, z)
		}
	}

	x["evil_dr"] = []string{"mr bigglesworth", "money", "sharks with lasers"}

	for k, v := range x {
		fmt.Println(k)
		for i, z := range v {
			fmt.Println(i, z)
		}
	}
}
