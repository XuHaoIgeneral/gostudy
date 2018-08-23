package main

import "fmt"

func main() {
	lists := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	test:=merge_sort(lists)
	fmt.Println(test)
}

//类似二分的拆解
func merge_sort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	mid := len(list) / 2
	left := merge_sort(list[:mid])
	right := merge_sort(list[mid:])
	new_seq:=merge_sort_list(left,right)
	return new_seq
}

//比较 并且进行合并
func merge_sort_list(seq_a, seq_b []int) []int {
	count_a, count_b := len(seq_a), len(seq_b)

	a,b:=0,0
	new_seq:=make([]int,0)


	//公共小标值的合并
	for a<count_a && b<count_b{
		if seq_a[a]>=seq_b[b]{
			new_seq=append(new_seq,seq_b[b])
			b++
		}else {
			new_seq=append(new_seq,seq_a[a])
			a++
		}
	}

	//非公共下表值的合并
	for a<count_a {
		new_seq=append(new_seq,seq_a[a])
		a++
	}

	for b<count_b  {
		new_seq=append(new_seq,seq_b[b])
		b++
	}

	return new_seq


}
