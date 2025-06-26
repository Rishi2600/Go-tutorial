package main

import "fmt"

func main() {
	var age map[string]int = map[string]int{
		"rishi":     23,
		"cristiano": 40,
		"grandma":   56,
		"gojo":      35,
		"naruto":    27,
	}
	fmt.Println(age)

	delete(age, "grandma")
	fmt.Println(age)
}
