package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name string
	Age  int
	Sex  string
}

// GoWalk ...Создаем метод пойти на прогулку, ссылаясь на тип Human
func (h Human) GoWalk() {
	fmt.Println(h.Name, ": I'm walking now!")
}

type Action struct {
	Human
}

func (a Action) GoDrink() {
	fmt.Println(a.Name, ": I'm drinking now!")
}

func main() {
	// Вариант 1 (как в ТЗ)
	jenna := Action{Human{
		Name: "Jenna",
		Age:  24,
		Sex:  "Female",
	}}
	jenna.GoDrink()
	// Вариань 2 (альтернативно)
	steve := Human{
		Name: "Steve",
		Age:  30,
		Sex:  "Male",
	}
	walk := Action{steve}
	walk.GoWalk()
}
