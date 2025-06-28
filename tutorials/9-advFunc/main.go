package main

import "fmt"

func main() {
	result1 := calculator(89, 24, sum)
	result2 := calculator(25, 69, sub)

	fmt.Println(result1, result2)
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func calculator(a int, b int, fn func(int, int) int) int {
	return fn(a, b)
}
