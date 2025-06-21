package main

import "fmt"

type shape interface {
	area() float32
	perimeter() float32
}

type rect struct {
	width  float32
	height float32
}

func (r rect) area() float32 {
	return r.width * r.height
}

func (r rect) perimeter() float32 {
	return 2 * (r.width + r.height)
}

type square struct {
	side float32
}

func (s square) area() float32 {
	return s.side * s.side
}

func (s square) perimeter() float32 {
	return 4 * s.side
}

func print(shape shape) {
	fmt.Printf("area is %v \n", shape.area())
	fmt.Printf("perimeter is %v \n", shape.perimeter())

}

func main() {
	var s square = square{side: 12}
	var r rect = rect{height: 5, width: 6}

	print(s)
	print(r)
}
