package main

import (
	"fmt"
	)

func main()  {
	str:="测测试test strinng"
	str1,str2:=strintTostr(str,2)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println("=============")
	str3:=changestr(str)
	fmt.Println(str3)
	fmt.Println("=============")
	str4:=setstr(str)
	fmt.Println(str4)
	fmt.Println("=============")
	nums:=[]int{1,2,3,4,5,6,7,8,9,0}
	num1:=mapfunc(maps,nums)
	fmt.Println(num1)
}


func strintTostr(str string,i int) (string,string){
	strings:=[]rune(str)
	str1:=string(strings[:i])
	str2:=string(strings[i:])
	return str1,str2
}


func changestr(str string) string {
	strings:=[]rune(str)
	count:=len(strings)
	for i:=0;i<int(count/2); i++{
		strings[i],strings[count-1-i]=strings[count-1-i],strings[i]
	}
	res:=string(strings)
	return res
}


func setstr(str string) string{

	strings:=[]rune(str)
	var lastby rune=strings[0]
	str1:=make([]rune,0,len(strings))
	str1=append(str1,lastby)
	for i:=1;i<len(strings);i++ {
		if lastby==strings[i] {
			continue
		}
		lastby=strings[i]
		str1=append(str1,strings[i])
	}
	return string(str1)
}


func mapfunc(f func(int) int,nums []int) []int{
	list:=make([]int,len(nums))
	for key,val:=range nums{
		list[key]=f(val)
	}
	return list
}


func maps(i int) int{
	return i*10
}