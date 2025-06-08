package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumIntSlice(intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumFloat32Slice(float32Slice))
}

func sumIntSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sumFloat32Slice(slice []float32) float32 {
	var sum float32
	for _, v := range slice {
		sum += v
	}
	return sum
}
