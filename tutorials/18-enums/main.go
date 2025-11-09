package main

import "fmt"

func main() {
	const (
		East
		West = iota
		North 
		South
	)

	fmt.Println(East)
	fmt.Println(West)
	fmt.Println(North)
	fmt.Println(South)

}
