package main

import "fmt"

func badCall()  {
	panic("bad end")
}

func test(){
	defer func() {
		if e:=recover();e!=nil {
			fmt.Println("Panicing %s\r\n",e)
		}
	}()
	badCall()
	fmt.Println("After bad call\r\n")
}

func main()  {
	fmt.Println("Calling test")
	test()
	fmt.Println("Test completed\r\n")
}
