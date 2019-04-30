package main

import "fmt"
import "math"

type square struct {
	length int
}

type circle struct {
	radius int
}

type shape interface {
	area() float64
}

func main() {
	c1 := circle{
		radius: 5,
	}

	s1 := square{
		length: 5,
	}

	info(c1)
	info(s1)
}

func (s square) area() float64 {
	return float64(s.length * s.length)
}

func (c circle) area() float64 {
	return math.Pi * float64(c.radius*c.radius)
}

func info(s shape) {
	fmt.Println(s.area())
}
