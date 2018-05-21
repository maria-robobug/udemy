package main

import "fmt"

type shape interface {
	printArea() float64
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tr := triangle{height: 5.5, base: 10.5}
	sq := square{sideLength: 8.5}

	tr.printArea()
	sq.printArea()
}

func (t triangle) printArea() {
	fmt.Println("The triangle area is: ", t.getArea())
}

func (s square) printArea() {
	fmt.Println("The square area is: ", s.getArea())
}

// 0.5 * base * height
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

// sideLength * sideLength
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
