package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20

func main() {
	a := new(big.Int)
	a.SetString("48000000000000000000", 10)
	b := new(big.Int)
	b.SetString("24000000000000000000", 10)
	r := new(big.Int)
	fmt.Println(r.Add(a, b)) // +
	fmt.Println(r.Sub(a, b)) // -
	fmt.Println(r.Mul(a, b)) // *
	fmt.Println(r.Div(a, b)) // /
}
