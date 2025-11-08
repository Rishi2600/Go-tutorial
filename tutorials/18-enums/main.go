package main

import "fmt"

func main() {
	const (
		East 
		West 
		North 
		South = iota
	)

	fmt.Println(East)
	fmt.Println(West)
	fmt.Println(North)
	fmt.Println(South)

}
