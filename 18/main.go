package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

func main() {
	type Counter struct {
		Counts int
	}

	var wg sync.WaitGroup
	ctr := Counter{}

	for i := 1; i < 11; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ctr.Counts += 1
			fmt.Printf("[%d]++\n", id)
		}(i)
	}
	wg.Wait()

	fmt.Println(ctr.Counts)
}
