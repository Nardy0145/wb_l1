package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

var wg sync.WaitGroup

func main() {
	chanClose()
	wg.Wait()
	timeout()
	wg.Wait()
	interrupt()
}

// Остановка по прошествию времени
func timeout() {
	wg.Add(1)
	var duration time.Duration
	duration = 3
	endDuration := time.After(duration * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("\n\nThird goroutine will be closed after 3 seconds")
		for {
			select {
			case <-endDuration:
				fmt.Println("Time is out. Closing")
				return
			default:
				fmt.Print("Another second of time...\n")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

// Остановка с помощью Interrupt
func interrupt() {
	wg.Add(1)
	fmt.Println("\n\nWaiting for an interrupt.")
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	// Здесь создаём сигнал и присваеваем его каналу signalChan
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		defer wg.Add(1)
		for range signalChan {
			fmt.Printf("Received an interrupt, stopping.")
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

// Остановка с помощью закрытия канала
func chanClose() {
	fmt.Println("Channel closing")
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch:
				fmt.Println("Stopping")
				return
			default:
				fmt.Println("One second later...")
				time.Sleep(1 * time.Second)
				// ...какие-то действия
			}
		}
	}()
	// ..что-то выполняемое...
	// ...............
	// Остановка
	time.Sleep(5 * time.Second)
	close(ch)
}
