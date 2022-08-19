package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	m := make(map[string]struct{})
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, v := range strings {
		m[v] = struct{}{}
	}
	fmt.Println(m)
}
