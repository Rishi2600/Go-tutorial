package main

import "fmt"

type messageTosend struct {
	phoneNumber int64
	message     string
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
}

func message(m messageTosend) {
	fmt.Printf("message sent to phone numbe: %v, & it is %v \n", m.phoneNumber, m.message)
}
