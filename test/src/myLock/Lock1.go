package main

import (
	"sync"

	"fmt"
	"time"
)

func main()  {
	var mutex sync.Mutex
	fmt.Println("准备在main主程上加锁")
	mutex.Lock()
	fmt.Println("加锁成功 在main主程上")

	for i:=1;i<4 ;i++  {
		go func(i int) {
			fmt.Printf("进入Go%d程,并加锁\n",i)
			mutex.Lock()
			fmt.Printf("加锁成功 在Go%d程上\n",i)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (G0)")
	//解锁mutex
	mutex.Unlock()

	fmt.Println("The lock is unlocked. (G0)")
	//休息一会,等待打印结果
	time.Sleep(time.Second)

}
