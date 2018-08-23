package main

import "fmt"

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	
	midPivot := list[0]
	var lessBefor []int
	var biggerAfter []int

	for i := 1; i < len(list); i++ {
		if list[i] > midPivot {
			biggerAfter = append(biggerAfter, list[i])
		} else {
			lessBefor = append(lessBefor, list[i])
		}
	}

	finallyList := append(quicksort(lessBefor), midPivot)
	finallyList=append(finallyList,quicksort(biggerAfter)...)
	return finallyList
}

func main() {
	lists := []int{10,9,8,7,6,5,4,3,2,1}
	l := quicksort(lists)
	for _, v := range l {
		fmt.Println(v)
	}
}
