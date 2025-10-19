package main

import "fmt"

func main() {
	const (
		East = iota
		West
		North
		South
	)

	fmt.Println(East)
	fmt.Println(West)
	fmt.Println(North)
	fmt.Println(South)

}
