package main

import "fmt"

func main() {
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

func getUser() (string, error) {
	var result = "Hello"
	var err error = nil
	return result, err
}
