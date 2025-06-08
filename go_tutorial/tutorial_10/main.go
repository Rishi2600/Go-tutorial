package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(isEmpty(intSlice))

	var floatSlice = []float32{1, 2, 3}
	fmt.Println(isEmpty(floatSlice))

}

func isEmpty[T any](slice []T) bool {

	return len(slice) == 0
}
