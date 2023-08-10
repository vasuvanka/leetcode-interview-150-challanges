package main

import "fmt"

func main() {
	list := []int{1, 1, 1, 2, 2, 3}
	newList, result := removeDuplicates(list)
	fmt.Println(newList, result)
}

func removeDuplicates(list []int) ([]int, int) {
	position := 0
	for i := 0; i < len(list); i++ {
		if position == 0 || position == 1 || list[i] != list[position-2] {
			list[position] = list[i]
			position++
		}
	}
	return list, position
}
