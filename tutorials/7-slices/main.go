package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}

	var value = 0

	for i := range array {
		value += array[i]
	}

	fmt.Println(value)
	fmt.Println(array)

	var maze [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(maze)
}

/*Size of arrays in Go has a predefined size and very strictly typed type whereas in JS, the size of array is dynamic.*/
