package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	// _, err := sqrtWithFmtErrorF(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("hello")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of a negative number")
	}

	return 42, nil
}

func sqrtWithFmtErrorF(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math: square root of a negative number: %v", f)
	}

	return 42, nil
}

//build your own error
type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func sqrtWithOwnError(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math: square root of a negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}

	return 42, nil
}
