package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()

		}(&wg, i)
	}
	wg.Wait()

	fmt.Println("exit")
}

// Даная программа выведет DEADLOCK. Потому, что не передан указатель на уже существующую сущность WG. Вместо этого, создаётся новая
// И происходит ошибка. Фикс - либо не перавать WG вообще, либо передавать с указателем и ссылкой.
