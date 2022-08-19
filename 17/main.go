package main

import (
	"fmt"
	"sort"
)

// Реализовать бинарный поиск встроенными методами языка.

// Тут я решил написать его сам
func main() {
	arr := []int{1, 4, 13, 100, 23, 65, 77, 79, 16}
	sort.Ints(arr)
	x := binFind(arr, 65)
	fmt.Println(x)
}

func binFind(arr []int, num int) bool {
	// finding := 65
	first := 0
	last := len(arr) - 1
	for first <= last {
		middle := (first + last) / 2
		if arr[middle] < num {
			first = middle + 1
		} else {
			last = middle - 1
		}
	}
	if first == len(arr) || arr[first] != num {
		return false
	}
	return true
}
