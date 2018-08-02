package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	wheelCount int
	Engine
}

func (self *Car) Start(){
	fmt.Println("汽车启动")
}

func (self *Car) Stop(){
	fmt.Println("汽车熄火")
}


func (self *Car) numberOfWheels()  {
	fmt.Printf("汽车轮子数目为 %d\n",self.wheelCount)
}


type Mercedes struct {
	Car
}

func (self *Mercedes)_sayHiToMerkel()  {
	fmt.Println("梅赛德斯奔驰:Hi")
}


func main()  {


	testcar:=new(Mercedes)
	testcar.wheelCount=4
	testcar.Start()
	testcar._sayHiToMerkel()
	testcar.numberOfWheels()
	testcar.Stop()

}