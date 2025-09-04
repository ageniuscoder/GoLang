package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	fmt.Println("Circle Area:", CalculateArea(circle))
	fmt.Println("Rectangle Area:", CalculateArea(rectangle))

}
