package main

import "fmt"

func main() {
	var props string = "this is a parameter passed in between the functions: from main to printMe"
	greet(props)

	printMe()
}

func printMe() {
	fmt.Println("Hello world")
}

func greet(props string) {
	fmt.Println(props)
}
