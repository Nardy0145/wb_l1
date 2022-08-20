package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	nums := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(removeNum(nums, 6))
}

func removeNum(sl []int, i int) []int {
	var tempSlice []int
	// Итерируем и сверяем текущее значение на данное
	for _, v := range sl {
		if v != sl[i-1] {
			tempSlice = append(tempSlice, v)
		}
	}
	return tempSlice
}
