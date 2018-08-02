package main

import "fmt"

type employee struct {
	salary float64
}

func Myemployee(salary float64) *employee{
	return &employee{salary}
}

func (self *employee) giveRaise(val float64) float64{
	self.salary=self.salary+self.salary*val
	return self.salary
}

func main(){

	test1:=Myemployee(1024.5)

	re:=test1.giveRaise(0.2)

	fmt.Println(re)
}

