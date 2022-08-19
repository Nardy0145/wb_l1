package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	sentence := "SNOW DOG SUN"
	fmt.Println(reverseSentence(sentence))
}

func reverseSentence(sentence string) string {
	arrSentence := strings.Split(sentence, " ")
	var revSentence string
	for i := len(arrSentence) - 1; i >= 0; i-- {
		revSentence += arrSentence[i]
		revSentence += " "
	}
	return revSentence
}
