package main

import (
	"fmt"
	)

func sendData(ch chan string){
	ch<-"a"
	ch<-"b"
	ch<-"c"
	ch<-"d"
	ch<-"e"
	ch<-"f"
}

func sendData2(ch chan string){
	ch<-"a1"
	ch<-"b1"
	ch<-"c1"
	ch<-"d1"
	ch<-"e1"
	ch<-"f1"
}

func getDate(ch chan string){
	var input string

	for true {
		input=<-ch
		fmt.Println(input)
	}
}

func main()  {
	ch:=make(chan string)
	go getDate(ch)
	go sendData(ch)
	go sendData2(ch)
	//time.Sleep(1e9)
}