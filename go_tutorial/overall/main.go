package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 50}
	fmt.Println(array)
	array[0] = 596
	fmt.Println(array)
}
