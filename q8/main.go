package main

import (
	"fmt"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	bitNum := replaceAtIndex(1510, 1, 1)
	fmt.Println(bitNum)
	intNum, _ := strconv.ParseInt(bitNum, 2, 64)
	fmt.Println(intNum)
}

func replaceAtIndex(num int64, ind int, to int) string {
	bitNum := strconv.FormatInt(num, 2)
	if len(bitNum) < ind {
		fmt.Println("Too big index!")
		return ""
	}
	runeNum := []rune(bitNum)
	switch to {
	case 1:
		runeNum[ind] = 49
	case 0:
		runeNum[ind] = 48
	}
	strNum := string(runeNum)
	return strNum
}
