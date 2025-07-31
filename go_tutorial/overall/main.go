package main

import "fmt"

func main() {
	var sum = func(a, b int) int {
		return a + b
	}(2, 3)

	fmt.Println(sum)
}
