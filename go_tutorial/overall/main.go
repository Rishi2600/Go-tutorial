package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 50}
	var slice []int = make([]int, 5)

	slice = append(slice, 56)

	fmt.Println(array, slice)

}
