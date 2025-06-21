package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
	perimeter() float32
}

type rect struct {
	height float32
	width  float32
}

type circle struct {
	radius float32
}

func main() {
	var rect rect = rect{
		height: 55,
		width:  4,
	}
	fmt.Println(rect)

	var circle circle = circle{
		radius: 5,
	}
	fmt.Println(circle)
}

func (r rect) area() float32 {
	return r.width * r.height
}

func (r rect) perimeter() float32 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}
