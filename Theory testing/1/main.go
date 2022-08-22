package main

import (
	"fmt"
	"strings"
)

// Какой самый эффективный способ конкатенации строк?

// ответ:
//string.Builder()

func main() {
	a := "12345"
	b := "67890"
	joined := join(a, b)
	fmt.Print(joined)
}

func join(strs ...string) string {
	var sb strings.Builder
	for _, v := range strs {
		sb.WriteString(v)
	}
	return sb.String()
}
