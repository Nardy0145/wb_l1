package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	fmt.Println("Start")
	Sleep(3 * time.Second)
	fmt.Println("3 seconds passed")
}

// Для написания своего слип я буду использовать time.After().

func Sleep(s time.Duration) {
	select {
	case <-time.After(s):
	}
}
