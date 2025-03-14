package main

import "fmt"


type Shape interface {
	area() float64
}
type Square struct {
	side float64
}

type Triangle struct {
	height float64
	base float64
}

func main() {
	s := Square{side: 10}
	t := Triangle{height: 12, base: 6}

	printArea(s)
	printArea(t)
}

func (s Square) area() float64 {
	return s.side * s.side
}	

func (t Triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s Shape) {
	fmt.Println(s.area())
}