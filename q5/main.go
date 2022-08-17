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
	for {
		select {
		case <-endDuration:
			fmt.Println("Timed out!")
			return
		default:
			fmt.Println("One second has been passed.")
			time.Sleep(1 * time.Second)
		}
	}
}
