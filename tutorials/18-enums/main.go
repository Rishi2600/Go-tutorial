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

}
