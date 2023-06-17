package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Bagus")
	fmt.Println(message)
}
