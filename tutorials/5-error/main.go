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
	var err error = fmt.Errorf("there is an error")
	return result, err
}
