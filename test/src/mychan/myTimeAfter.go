package main

import (
	"fmt"
	"time"
)

func goOut(c chan int)  {
	for i:=0;i<5 ;i++  {
		time.Sleep(time.Second*1)
		c<-i
	}
}
func main(){
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go goOut(ch1)
	go goOut(ch2)
	for {
		select {
		case e1 := <-ch1:
			//如果ch1通道成功读取数据，则执行该case处理语句
			fmt.Printf("1th case is selected. e1=%v\n", e1)
		case e2 := <-ch2:
			//如果ch2通道成功读取数据，则执行该case处理语句
			fmt.Printf("2th case is selected. e2=%v\n", e2)
		case <-time.After(2 * time.Second):
			fmt.Println("Timed out")
		}
	}
}
