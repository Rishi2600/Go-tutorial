package main

import "fmt"

func main() {
	var intSliceEmpty = []int{1, 2, 3}
	fmt.Println(isEmpty(intSliceEmpty))

	var floatSliceEmpty = []float{1, 2, 3}
	fmt.Println(isEmpty(floatSliceEmpty))

	var intSliceNotEmpty = []int{1, 2, 3}
	fmt.Println(isNotEmpty(intSliceNotEmpty))

	var floatSliceNotEmpty = []float32{1, 2, 3}
	fmt.Println(isNotEmpty(floatSliceNotEmpty))

}

func isEmpty[T any](slice []T) bool {

	return len(slice) == 0
}

func isNotEmpty[T any](slice []T) bool {
	return len(slice) != 0
}
