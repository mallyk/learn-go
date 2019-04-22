package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken", "Not Stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	x := [][]string{jb, mp}
	fmt.Println(x)

	for k, v := range x {
		fmt.Println(k, v)
		for i, z := range v {
			fmt.Println(i, z)
		}
	}
}
