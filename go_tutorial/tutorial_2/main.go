package main

import "fmt"

func main() {
	var props string = "this is a parameter passed in between the functions: from main to printMe"
	greet(props)

	printMe()

	var num int = 50
	var den int = 5
	var value int = intDivision(num, den)
	fmt.Println(value)

}

func printMe() {
	fmt.Println("")
}

func greet(props string) {
	fmt.Println(props)
}

func intDivision(num int, den int) int {
	var result int = num / den
	return result
}
