package main

import "fmt"

func main() {
	sum := 0

	for i := range 10 {
		sum = sum + i
	}
	fmt.Println(sum)
}
