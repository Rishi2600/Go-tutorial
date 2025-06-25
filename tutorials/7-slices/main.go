package main

import "fmt"

func main() {

	var users []string = make([]string, 3)
	println(len(users))
	println(cap(users))

	fmt.Println(users[0] == "")
	fmt.Println(users)

}
