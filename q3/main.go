package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	var wg sync.WaitGroup
	nums := []int{2, 4, 6, 8, 10}
	var x []int
	var result int
	// Создадим цикл, конкурентно высчитывающий квадраты и записывающий результат в х
	for _, value := range nums {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			i := value * value
			x = append(x, i)
		}(value)
		wg.Wait()
	}
	// Посчитаем сумму чисел всего х.
	for _, value := range x {
		result += value
	}
	// Выведем результат
	fmt.Println(result)
}
