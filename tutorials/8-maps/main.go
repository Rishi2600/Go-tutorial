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
		fmt.Printf("value of %v is: %v \n", key, value)
	} else {
		fmt.Println("\n value not found, hence deleted")
	}

	age["tanjiro"] = 18
	fmt.Println(age)

	var elem = age["rishi"]
	fmt.Println(elem)
}

/*The second variable
(whether you name it ok, exists, found, or something else) will be a bool that indicates whether the key was found in the map.*/
