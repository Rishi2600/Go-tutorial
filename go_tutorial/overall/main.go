package main

import (
	"fmt"
)

func main() {

	sum := 0

	func() {
		for i := range 3 {
			sum += i
			fmt.Println(sum)
		}
		fmt.Println(sum)
	}()
	func() {
		for i := range 3 {
			sum += i
			fmt.Println(sum)
		}
		fmt.Println(sum)
	}()

	fmt.Println("done")
}
