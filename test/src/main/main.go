/*
	test  md
 */

package main

import  (
	he "hello"
	"fmt"
	"trans"
	"mystack"
)

var twoPi = 2 * trans.Pi

func main() {
	fmt.Println("=========TEST=======")
	he.Hello("Go!")
	he.Hello1("Go!")
	he.Hello2()
	fmt.Printf("2*Pi = %g\n", twoPi)
	test:=mystack.MyStack()
	test.Push(1)
	test.Pop()
}
