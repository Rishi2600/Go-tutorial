package main

import "fmt"

func main() {
	var sum1 int = 0
	for i := 0; i < 10; i++ {
		sum1 += i //sum = sum + i
	}

	var sum2 int = 0
	for j := range 10 {
		sum2 += j
	}

	fmt.Println(sum1, sum2)
}
