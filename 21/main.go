package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

// Приведу в пример использование Micro-Jack наушников на Iphone 8, в котором такой разъём отсутствует.
// Есть только Lightning.

type Client struct{}
type Adapter interface {
	insertIntoLightning()
}
type MicroJack struct{}

func (m MicroJack) insertIntoLightning() {
	fmt.Println("Inserted into Lightning")
}
func (c *Client) insertMicroJackIntoLightning(ada Adapter) {
	ada.insertIntoLightning()
}

func main() {
	client := &Client{}
	microjack := &MicroJack{}
	client.insertMicroJackIntoLightning(microjack)
}
