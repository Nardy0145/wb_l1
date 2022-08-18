package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись в map.

func main() {
	someMap := make(map[int]bool)
	mutex := sync.RWMutex{}
	ch := make(chan int)
	for i := 1; i < 15; i++ {
		go func(i int) {
			// Блокируем чтение с мапы
			mutex.Lock()
			switch i % 2 {
			case 1:
				someMap[i] = false
			case 0:
				someMap[i] = true
			}
			// Разблокируем, записываем в канал, после чего в основном потоке это моментально считывается
			mutex.Unlock()
			ch <- i
		}(i)
	}
	for i := 1; i < 15; i++ {
		x := <-ch
		fmt.Println(x, " ", someMap[x])
	}
}
