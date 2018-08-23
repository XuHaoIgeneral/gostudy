package main

import (
	"fmt"
	"time"
)

func writeto(c chan string)  {
	c<-"go go go"
}

func readto(c chan string)  {
	fmt.Println(<-c)
}

func main()  {
	ch:=make(chan string)
	go readto(ch)
	go writeto(ch)
	time.Sleep(1e9)
}