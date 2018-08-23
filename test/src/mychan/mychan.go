package main

import (
	"fmt"
	)

func fibonacci(n int, c chan int) {

	fmt.Println("aa")
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		fmt.Println("a")
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	//c1:=make(chan int,10)
	go fibonacci(cap(c), c)
	//go fibonacci(cap(c1),c1)
	fmt.Printf("chan c Type is %v",c)
	for i := range c {
		fmt.Println(i)
	}
}
