package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 50}
	sum := 0

	for _, value := range array {
		sum += value
	}
	fmt.Println(sum)
}
