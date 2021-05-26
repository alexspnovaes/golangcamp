package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	w, h float64
}

type Circle struct {
	r float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.w * 2) + (r.h * 2)
}

func (r Rectangle) Area() float64 {
	return (r.w * r.h)
}

func (r Circle) Perimeter() float64 {
	return math.Pi * 2 * r.r
}

func (r Circle) Area() float64 {
	return math.Pi * math.Pow(r.r, 3)
}

func Calculate(s Shape) {
	var r Shape = Rectangle{
		w: 10,
		h: 20,
	}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())

	var c Shape = Circle{r: 10}
	fmt.Println(c.Area())
	fmt.Println(c.Perimeter())
}

func main() {
	var s Shape
	Calculate(s)
}
