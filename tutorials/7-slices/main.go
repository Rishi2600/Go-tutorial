package main

import "fmt"

func main() {

	var array [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	var slice = array[1:4]
	fmt.Println(slice)

	slice = []int{5, 6, 7, 8}
	fmt.Println(slice)

}
