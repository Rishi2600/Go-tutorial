package main

import "fmt"

func main() {
	var user, err = getUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func getUser() string {
	var result = "Hello"
	return result
}
