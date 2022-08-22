package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(sliceN *[]string) {
		*sliceN = append(*sliceN, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(&slice)
	fmt.Print(slice)
}

// [b b a] [a a] Проблема в том, что не передан указатель на ячейку памяти. Также идет затенение переменной
