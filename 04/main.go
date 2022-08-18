package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	// Запрос кол-ва воркеров.
	var workers int
	fmt.Println("how much workers?")
	_, err := fmt.Scan(&workers)
	if err != nil {
		log.Fatal("Workers input error: ", err)
		return
	}

	ch := make(chan string, workers)
	var wg sync.WaitGroup

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	for i := 1; i < workers+1; i++ {
		num := strings.Repeat("~", i)
		ch <- num
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				default:
					out := <-ch
					fmt.Println(out)
					ch <- out

				case signalSigterm := <-signalChan:
					signalChan <- signalSigterm
					return
				}
			}
		}()
	}

	wg.Wait()
}
