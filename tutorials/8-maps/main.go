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

	key := "gojo"
	value, exists := age[key]

	if exists {
		fmt.Printf("value of %v is: %v", key, value)
	} else {
		fmt.Println("\n value not found, hence deleted")
	}
}
