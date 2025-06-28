package main

import "fmt"

func main() {
	var result = add(7, 9)
	fmt.Println(result)

	fmt.Println(concat("Hello", " world"))
	fmt.Println(concat("Hey", " there!"))
	fmt.Println(concat("whats", " up!"))

	add, sub := calculator(1, 2)
	fmt.Printf("add: %v & sub: %v \n", add, sub)
}

func add(a int, b int) int {
	return a + b
}

func concat(first string, second string) string {
	return first + second
}

func calculator(a, b int) (int, int) {
	return a + b, a - b
}
