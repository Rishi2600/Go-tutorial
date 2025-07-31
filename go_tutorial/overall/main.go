package main

import "fmt"

type rect struct {
	height int
	width  int
}

func main() {
	value := rect{height: 5, width: 5}
	fmt.Println(value.area())
}

func (r rect) area() int {
	return r.height * r.width
}
