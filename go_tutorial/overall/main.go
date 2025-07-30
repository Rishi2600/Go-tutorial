package main

import "fmt"

func main() {
	gender := "ttrans"

	switch gender {
	case "male":
		fmt.Println("you are male")
	case "female":
		fmt.Println("you are female")
	default:
		fmt.Println("neither male nor female")
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
