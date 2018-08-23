package main

import (
	"time"
	"fmt"
)

func producer(ch chan int)  {
	for i:=0;;i++ {
		ch<-i
	}
}

func consumers(ch chan int)  {
	for true {
		fmt.Println(<-ch)
	}
}

func main()  {
	c:=make(chan int)
	go producer(c)
	go consumers(c)
	time.Sleep(2e9)
}
