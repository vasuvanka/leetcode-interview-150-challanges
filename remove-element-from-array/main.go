package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 4, 5}
	val := 4
	list1, result := removeElement(list, val)
	fmt.Println(list1, result)
}

func removeElement(list []int, val int) (int, []int) {
	j := 0
	for i := 0; i < len(list); i++ {
		if list[i] != val {
			list[j] = list[i]
			j++
		}
	}
	return j, list
}
