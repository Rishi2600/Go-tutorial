package main

import "fmt"

type messageTosend struct {
	phoneNumber int
	message     string
	number      struct {
		area string
		code int8
	}
}

func main() {
	message(messageTosend{
		phoneNumber: 789456132,
		message:     "Hello bro",
	})

	message(messageTosend{
		phoneNumber: 123456,
		message:     "you're the guy with phone number",
	})

	area(&rect{
		width:  200,
		height: 6,
	})
}

func message(m messageTosend) {
	fmt.Printf("message sent to phone numbe: %v, & it is %v \n", m.phoneNumber, m.message)
}

type rect struct {
	width  int32
	height int32
}

func area(r *rect) int32 {
	return r.height * r.width
}
