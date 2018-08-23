package main

import (
	"fmt"
	)

type action interface {
	Run() bool
	Write()
	Read()
}

type student struct {
	name string
}

type teacher struct{
	name string
}

func (this *student) Run() bool{
	if this.name==""{
		return false
	}
	fmt.Println(this.name,"is run")
	return true
}

func (this *teacher) Run() bool{
	if this.name==""{
		return false
	}
	fmt.Println(this.name,"is run")
	return true
}

func (this *student) Write()  {
	fmt.Println("Student is writing")
}

func (this *teacher) Write()  {
	fmt.Println("teacher is writing")
}

func (this *student) Read()  {
	fmt.Println("Student is reading")
}

func (this *teacher) Read()  {
	fmt.Println("teacher is reading")
}

func ScAction(ac action)  {
	if i:=ac.Run();i==true{
		ac.Write()
		ac.Run()
	}
}

func main()  {

	st:=new(student)
	st.name="Student"
	ScAction(st)
	te:=&teacher{name:"Teacher"}
	te.Run()
}

