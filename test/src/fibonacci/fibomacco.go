package main

import (
	"time"
	"fmt"
)

const LVM=40

var fibs [LVM]int

func main()  {
	startTime:=time.Now()
	for i:=0;i<LVM ;i++  {
		re:=fibomacci(i)
		fmt.Println(re)
	}
	endTime:=time.Now()
	times:=endTime.Sub(startTime)
	fmt.Printf("总共用时：%s",times)
}

func fibomacci(i int) (res int){
	if fibs[i]!=0{
		res=fibs[i]
		return
	}
	if i<=1 {
		res=1
		return
	}else {
		res=fibomacci(i-2)+fibomacci(i-1)
	}
	fibs[i]=res
	return
}