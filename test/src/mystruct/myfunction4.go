package main

import (
	"math"
	"fmt"
)

type Point struct {
	x,y float64
}

type NamePoint struct {
	Point
	name string
}

func (self *Point) Abs() float64{
	return math.Sqrt(self.x*self.x+self.y*self.y)
}

func main()  {
	n:=&NamePoint{Point{3,4},"Pythagpras"}
	fmt.Println(n.Abs())
}