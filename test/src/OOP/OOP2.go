package main

import "fmt"

type Animals struct {
	leg int
	action string
}

type Tiger struct {
	Animals
	tail string
}

type Goat struct {
	Animals
	horn string
}

type Anaction interface {
	Run()
	Eat()
}

func (this *Tiger) Run()  {
	fmt.Printf("老虎run，有%v条腿, %v，%v\n",this.leg,this.action,this.tail)
}

func (this *Tiger) Eat()  {
	fmt.Printf("老虎 eat %v\n",this.action)
}

func (this *Goat) Run()  {
	fmt.Printf("山羊run，有%v条腿, %v，%v\n",this.leg,this.action,this.horn)
}

func (this *Goat) Eat()  {
	fmt.Printf("山羊 eat %v\n",this.action)
}

func main()  {

	var act Anaction
	ti:=new(Tiger)
	ti.tail="老虎tail"
	ti.action="吃肉的"
	act=ti
	act.Run()
	act.Eat()

	goa:=&Goat{Animals{4,"吃草"},"山羊角"}
	act=goa
	act.Run()
	act.Eat()
}