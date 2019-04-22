package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:3]) //up to but not including index 3

	//loop without range
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	//append to slice
	x = append(x, 77, 88, 99, 114)
	fmt.Println(x)

	y := []int{234, 456, 987}
	x = append(x, y...)
	fmt.Println(x)

	//delete from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	//make slice of a certain size
	z := make([]int, 10, 100)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z[0] = 42
	z[9] = 99
	z = append(z, 3423)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	//range over slice
	for k, v := range x {
		fmt.Println(k, v)
	}

}
