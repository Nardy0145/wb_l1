package main

import (
	"fmt"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2
После чего данные из второго канала должны выводиться в stdout.
*/
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inpCh := make(chan int)
	outCh := make(chan int)
	stop := make(chan bool)

	// Запись
	go func(nums []int) {
		for _, value := range nums {
			inpCh <- value
		}
		close(inpCh)
	}(nums)
	// Чтение с последующей записью после операций
	go func() {
		for value := range inpCh {
			outCh <- value * value
		}
		close(outCh)
	}()
	// Вывод результата, закрытие главного канала
	go func() {
		for result := range outCh {
			fmt.Println(result)
		}
		stop <- true
	}()
	<-stop
}
