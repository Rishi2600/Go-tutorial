package main

import "fmt"

func main() {
	funcValue := multiply(5)
	finalValue := funcValue(5)

	fmt.Println(finalValue)

}

func multiply(n int) func(num int) int {
	return func(num int) int {
		return num * n
	}
}
