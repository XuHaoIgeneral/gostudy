package main

import "fmt"

type Rectangle struct {
	long int
	weight int
}

func main()  {
	test:=new(Rectangle)
	test.long=10
	test.weight=5
	area:=Area(&test.long,&test.weight)
	per:=Perimeter(&test.long,&test.weight)
	fmt.Println(area,per)
}

func Area(x,y *int) int{
	x1:=*x
	y1:=*y
	return x1*y1
}

func Perimeter(x,y *int) int{
	return (*x+*y)*2
}
