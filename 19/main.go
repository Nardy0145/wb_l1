package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func main() {
	fmt.Println(reverse("привет, как дела?"))
	fmt.Println(reverse("Ångström"))
	fmt.Println(reverse("\\xff\\xfe\\xfd"))
}

func reverse(word string) string {
	runeWord := []rune(word)
	var rev []rune
	for i := len(runeWord) - 1; i >= 0; i-- {
		rev = append(rev, runeWord[i])
	}
	return string(rev)
}
