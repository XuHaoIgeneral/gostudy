package main

import "fmt"

type Simpler interface {
	Set(i int)
	Get() int
}

type Simple struct {
	i int
}

func (self *Simple) Set(i int){
	self.i=i*10
}

func (self *Simple) Get() int{
	return self.i
}

func root(r Simpler,i int) int{
	r.Set(i)
	a:=r.Get()
	return a
}

func main()  {
	var sinter Simpler=new(Simple)
	re:=root(sinter,5)
	fmt.Println(re)
}