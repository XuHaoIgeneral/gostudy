package main

import (
		"fmt"
	"time"
)

func main() {
	c := make(chan int, 50)

	go func() {
		x := <-c
		fmt.Println("received", x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	time.Sleep(1e9)
	fmt.Println("sent", 10)
	time.Sleep(1e9)
}
