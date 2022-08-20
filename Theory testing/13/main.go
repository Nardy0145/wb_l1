package main

import "fmt"

func someAction(v *[]int, b int) {
	*v = append(*v, b)
}

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = [5]int{1, 2, 3, 4, 5}
	c := b[:]
	c = append(c, 6)
	fmt.Println(a)
	fmt.Println(b)
	b[0] = 100
	fmt.Println(c)
	someAction(&c, 7)
	fmt.Println(c)
}

// Выведет 100 2 3 4 5. Д
