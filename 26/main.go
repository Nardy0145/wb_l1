package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.
*/

func main() {
	fmt.Println(checkUnique("Dabcd"))
	fmt.Println(checkUnique("abcd"))
}

func checkUnique(s string) bool {
	s = strings.ToLower(s)
	// Запускаем цикл, итерирующий по каждому символу
	for i := 0; i < len(s); i++ {
		// Запускаем в итерации по символу еще одну, в которой итерируемый символ сверяется со всеми остальными (шаги учетны)
		for x := i + 1; x < len(s); x++ {
			if s[i] == s[x] {
				return false
			}
		}
	}
	return true
}
