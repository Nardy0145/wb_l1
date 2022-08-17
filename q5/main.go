package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	var duration time.Duration
	fmt.Println("Enter duration")
	_, err := fmt.Scan(&duration)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	endDuration := time.After(duration * time.Second)

	// Функция, в которой будет происходить чтение/запись в этот канал, а также остановка

	var x []int
	ch := make(chan int, 1)
	for {
		select {
		case <-endDuration:
			fmt.Println("Timed out!")
			fmt.Println(x)
			return
		default:
			ch <- 1
			x = append(x, <-ch)
			time.Sleep(500 * time.Millisecond)
		}

	}
}
