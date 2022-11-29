package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	a, b float64
}

type Circle struct {
	r float64
}

func (r Rectangle) Area() float64 {
	return r.a * r.b
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.a + r.b)
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func InterfaceFunc(s Shape) {
	fmt.Println(s)
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

func main() {
	r1 := Rectangle{a: 5, b: 10}
	c1 := Circle{5}
	/*fmt.Println(r1.Area())
	fmt.Println(r1.Perimeter())*/
	InterfaceFunc(r1)
	fmt.Println("********************")
	InterfaceFunc(c1)
}
