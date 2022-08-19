package main

import (
	"fmt"
	"github.com/imdario/mergo"
)

// Реализовать пересечение двух множеств.
func main() {
	mp1 := map[string]interface{}{
		"one": 1,
		"two": 2,
	}

	mp2 := map[string]interface{}{
		"three": 3,
		"four":  4,
	}
	err := mergo.Merge(&mp1, mp2)
	if err != nil {
		return
	}
	fmt.Println(mp1)
}
