package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	var wg sync.WaitGroup
	x := []int{2, 4, 6, 8, 10}
	for _, value := range x {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			fmt.Println(value * value)
		}(value)
		wg.Wait()
	}
}
