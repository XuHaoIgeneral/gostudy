package main

import "fmt"

func main() {
	s:="hello在"
	s1:=[]rune(s)
	s1[0]='好'
	fmt.Println(s1)
	//c:=[]byte(s)
	//fmt.Println(c)
	//c[0]='c'
	s2:=string(s1)
	fmt.Println(s2)
}