package main

import "fmt"

func main() {
	var result = add(7, 9)
	fmt.Println(result)

	fmt.Println(concat("Hello", " world"))
	fmt.Println(concat("Hey", " there!"))
	fmt.Println(concat("whats", " up!"))
}

func add(a int, b int) int {
	return a + b
}

func concat(first string, second string) string {
	return first + second
}
