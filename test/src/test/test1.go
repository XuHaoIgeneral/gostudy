package main

import "fmt"

type EmptyInterface interface {

}

type WithFuncInterface interface {
	Func()
}

type TestStruct struct {
	Member int
}

func (t *TestStruct) Func(){
	fmt.Println("HHH")
}

func TestEmptyInterfce(i EmptyInterface)  {

}