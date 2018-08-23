package main

import "fmt"

type action interface {
	live() bool
	run()
}

type father struct {
	name string
	sex  string
}

type son struct {
	father
}

func (self *father) live() bool {
	return true
}

func (self *son) live() bool {
	return true
}

func (self *father) run() {
	fmt.Println("对象为", self.name, "正在执行")
}

func (self *son) run() {
	fmt.Println("对象为", self.name, "正在执行")
}

func main() {
	var inter action
	fa := new(father)
	fa.name="father"
	fa.sex="man"
	son:=new(son)
	son.name="son"
	son.sex="man"
	inter=fa
	fmt.Println(inter.live())
	inter.run()

	var inter1 action
	inter1=son
	inter1.run()
}
