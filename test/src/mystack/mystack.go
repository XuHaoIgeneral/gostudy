package mystack

import "fmt"

type myStack struct {
	index int
	stackList [4]int
}

func MyStack() *myStack{
	StackClass:=new(myStack)
	StackClass.index=-1
	return StackClass
}
func (self *myStack) Push(i int){
	if self.index<3 {
		self.index++
		self.stackList[self.index]=i
		fmt.Printf("添加栈成功 值=%d 索引=%d\n",i,self.index)
	}else {
		fmt.Println(self)
		fmt.Println("添加栈失败")
	}
}

func (self *myStack) Pop() int{
	if self.index>-1{
		i:=self.stackList[self.index]
		self.stackList[self.index]=0
		self.index--
		fmt.Println("POP出栈正常")
		return i
	}else {
		fmt.Println("已经到了栈底")
		return 0
	}
}



