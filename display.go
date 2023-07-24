package main

import "fmt"

func AffichageCarre() {
	for _, ligne := range square {
		str := ""
		for _, caractère := range ligne {
			str += string(caractère) + " "
		}
		fmt.Println(str)
	}
}

func AffichageTetrominos(form Tetrominos) {
	fmt.Println("Id du Tetrominos: " + string(form.id))
	for _, ligne := range form.forme {
		str := ""
		for _, letter := range ligne {
			if letter == '#' {
				str += string(form.id) + " "
			} else {
				str += string(letter) + " "
			}
		}
		fmt.Println(str)
	}
	fmt.Println(form.co)
}
