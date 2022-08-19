package main

import (
	"fmt"
	"sort"
)

func quickSort(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func main() {
	array := []int{3, 4, 1, 15, 2, 7, 5, 6, 8, 9, 13, 12, 10, 11, 14}
	x := quickSort(array)
	fmt.Println(x)
}
