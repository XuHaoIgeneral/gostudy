package main

import "fmt"

func insertSort(list []int) []int{
	count:=len(list)
	for i:=1;i<count;i++ {
		key:=list[i]
		j:=i-1
		for j>=0 {
			if list[j]>key {
				list[j+1]=list[j]
				list[j]=key
			}
			j--
		}
	}
	return list
}

func main()  {
	list:=[]int{6, 5, 8, 9, 7, 2, 3, 4, 5, 6}
	l:=insertSort(list)
	for _,v:=range list{
		fmt.Println(l[v])
	}
}