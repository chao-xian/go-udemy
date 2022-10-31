package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Square struct {
	sideLength float64
}

type Triangle struct {
	height float64
	base   float64
}

func (sq Square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tri Triangle) getArea() float64 {
	return 0.5 * tri.base * tri.height
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := Square{sideLength: 4.0}
	tri := Triangle{height: 7.0, base: 3}

	printArea(sq)
	printArea(tri)
}
