package main

import "fmt"

func main() {
	list := []int{1, 1, 2}
	val := 1
	newList, result := removeDuplicates(list, val)
	fmt.Println(newList, result)
}

func removeDuplicates(list []int, val int) ([]int, int) {
	j := 1
	for i := 1; i < len(list); i++ {
		if list[i] != list[i-1] {
			list[j] = list[i]
			j++
		}
	}
	return list, j
}
