package main

import "fmt"

func main() {
	result1 := calculator(89, 24, sum)
	result2 := calculator(25, 69, sub)

	fmt.Println(result1, result2)

	double := multiplier(5)
	triple := multiplier(3)

	fmt.Println(double(3))
	fmt.Println(triple(3))

	numbers := []int{1, 2, 3, 4}
	doubled := mapInts(double, numbers)

	fmt.Println(doubled)
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

func multiplier(number int) func(int) int {
	return func(a int) int {
		return a * number
	}
}

func mapInts(f func(int) int, nums []int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = f(v)
	}
	return result
}

func double(x int) int {
	return x * 2
}
