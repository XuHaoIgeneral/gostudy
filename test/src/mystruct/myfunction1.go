package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func MyTwoInts(a,b int) *TwoInts{
	return &TwoInts{a,b}
}
func main() {
	two1 := MyTwoInts(12,11)
	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
	fmt.Println(two2.AddToParam(10))
}


func (self *TwoInts) AddThem() int {
	return self.a + self.b
}


func (self *TwoInts) AddToParam(param int) int {
	return self.a + self.b + param
}
