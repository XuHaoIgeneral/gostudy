package main

import "fmt"

func main() {
	functions := make([]func(), 3)
	for i := 0; i < 3; i++ {
		functions[i] = func() {
			fmt.Println(fmt.Sprintf("iterator value: %d", i))
		}
	}

	functions[0]()
	functions[1]()
	functions[2]()
}