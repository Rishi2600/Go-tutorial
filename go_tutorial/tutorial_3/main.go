package main

import "fmt"

func main() {
	x := 7

	switchCase(x)

	ifElseCase(x)

}

func switchCase(value int) {
	switch value {
	case 1:
		fmt.Println("value is 1")
	case 2:
		fmt.Println("value is 2")
	case 3:
		fmt.Println("value is 3")
	case 4:
		fmt.Println("value is 4")
	case 5:
		fmt.Println("value is 5")
	default:
		fmt.Println("value is not 1, 2, 3, 4, or 5")
	}
}

func ifElseCase(value int) {
	if value == 1 {
		fmt.Println("value is 1")
	} else if value == 2 {
		fmt.Println("value is 2")
	} else if value == 3 {
		fmt.Println("value is 3")
	} else if value == 4 {
		fmt.Println("value is 4")
	} else if value == 5 {
		fmt.Println("value is 5")
	} else {
		fmt.Println("value is not 1, 2, 3, 4, or 5")
	}
}
