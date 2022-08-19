package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}

func main() {
	var someInterface interface{}
	someInterface = 1
	fmt.Println(reflect.TypeOf(someInterface))
	someInterface = "hello world!"
	fmt.Println(reflect.TypeOf(someInterface))
	someInterface = true
	fmt.Println(reflect.TypeOf(someInterface))
	someInterface = struct {
		me int
	}{}
	fmt.Println(reflect.TypeOf(someInterface))
	someInterface = make(chan string)
	fmt.Println(reflect.TypeOf(someInterface))
}
