package main

import (
	"fmt"
)

func main() {
	init := []rune{'.'}
	square = append(square, init)
	test := Initdata()
	if !test {
		fmt.Println("ERROR")
		return
	}
	Result()
	AffichageCarre()
}
