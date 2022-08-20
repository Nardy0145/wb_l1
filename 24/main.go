package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	X float64
	Y float64
}

func main() {
	first := Constructor(5, 10)
	second := Constructor(15, 2)
	difference := Calculate(first, second)
	fmt.Println(difference)
}

func Constructor(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// Для вычисления воспользуемся формулой из геометрии.
// d = sqrt( (x2 - x 1)^2 + (y2 - y 1)^2 )

func Calculate(first, second *Point) float64 {
	math.Pow(second.X-first.X, 2)
	return math.Sqrt(math.Pow(second.X-first.X, 2) + (math.Pow(second.Y-first.Y, 2)))
}
