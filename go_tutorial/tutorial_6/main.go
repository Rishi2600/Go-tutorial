package main

import "fmt"

type object struct {
	name 
	age 
}

func (e object) details() string {
	fmt.Print(e.name, e.age)
	return e.name
}

func main() {
	var user object = object{name: "rishi", age: 23}
	user.name = "cristiano"
	fmt.Printf("user name: %v & user age: %v \n", user.name, user.age)

	user.details()
}

/*type struct and interface in go is very similar to that of typescript*/
/*16 line -- {name: "rishi", age: 23} can also be written directly as {"rishi", 23}*/
