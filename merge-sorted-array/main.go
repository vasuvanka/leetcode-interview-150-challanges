package main

import "fmt"

func main() {
	list1 := []int{1, 2, 3, 0, 0, 0}
	list2 := []int{2, 5, 6}
	m := 3
	n := 3
	result := merge(list1, m, list2, n)
	fmt.Println(result)
}

func merge(list1 []int, m int, list2 []int, n int) []int {
	cursorM := m - 1
	cursorN := n - 1
	cursor := m + n - 1
	for cursorN >= 0 {
		if cursorM >= 0 && list1[cursorM] > list2[cursorN] {
			list1[cursor] = list1[cursorM]
			cursorM -= 1
		} else {
			list1[cursor] = list2[cursorN]
			cursorN -= 1
		}
		cursor--
	}
	return list1
}
