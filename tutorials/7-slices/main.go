package main

import "fmt"

func main() {

	var array [5]int = [5]int{1, 2, 3, 4, 5}

	var slice = array[1:4]

	fmt.Println(array)
	fmt.Println(slice)

}
