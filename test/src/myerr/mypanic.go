package main

import (
	"fmt"
	"os"
)
var user = os.Getenv("USER")

func main() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}

func check() {
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
}