package main

import (
	"time"
	"fmt"
)

func main() {
	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
	v,k:=<-c
	fmt.Println(v,k)
	fmt.Println("Finished")
}