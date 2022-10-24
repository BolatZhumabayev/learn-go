package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func pringArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	tShape := triangle{
		base:   2,
		height: 4,
	}
	sShape := square{
		side: 9,
	}

	pringArea(tShape)
	pringArea(sShape)
}
