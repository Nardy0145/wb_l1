package main

// Что такое интерфейсы, как они применяются в Go?

// ответ:
// 1) Для проверки чего-то на наличие метода
// 2) Как универсальный контейнер для переменной любого типа
// 3) Как мост при реализации ооп

func main() {
	me := Human{
		Head:       1,
		Arm:        2,
		Leg:        3,
		UrineValue: 100,
	}
	me.goPiss()
}

type Human struct {
	Head       int
	Arm        int
	Leg        int
	UrineValue int
}

type Animal struct {
	UrineValue int
}

func (h *Human) goPiss() {
	h.UrineValue = 0
}

func (a *Animal) goPiss() {
	a.UrineValue = 0
}

type Piss interface {
	goPiss()
}
