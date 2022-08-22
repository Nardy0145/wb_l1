package main

import "fmt"

// В  чем разница make и new?

// answer:
// new - Аллоцирует память под заданный тип. Возвращает указатель на него.
// make - Аллоцирует память и создаетиспользуется только для выделения и инициализации объектов типов срезов, карт и каналов.
// Все три типа являются структурами. Возвращаемое значение - это тип, а не указатель.

func main() {
	id := new(int)
	name := new(string)
	flag := new(bool)
	fmt.Printf("id type: %T  value: %v\n", id, *id)
	fmt.Printf("name type: %T  value: %v\n", name, *name)
	fmt.Printf("flag type: %T  value: %v\n", flag, *flag)
}
