package main

import (
	"fmt"
	"os"
)

func Initdata() bool {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("nombre d'arguments invalide")
		return false
	}
	data, err := os.ReadFile("datatest/" + string(args[0])) // lire le fichier text.txt
	if err != nil {
		fmt.Println(err)
	}
	sautligne := 0
	ligne := 0
	col := 0
	var form Tetrominos
	id := 'A'
	compt := 0
	for _, letter := range string(data) {
		form.id = id
		if letter == '.' || letter == '#' {
			sautligne = 0
			if letter == '#' {
				form.co = append(form.co, [2]int{ligne, col})
				compt++
			}
			form.forme[ligne][col] = letter
			col++
		} else if letter != '\n' || sautligne >= 2 && letter == '\n' {
			return false
		} else {
			sautligne++
		}
		if col == 4 {
			col = 0
			ligne++
		}
		if ligne == 4 && compt == 4 {
			compt = 0
			ligney := 0
			colx := 0
			for index, co := range form.co {
				if index == 0 {
					ligney = co[0]
					colx = co[1]
				}
				form.co[index] = [2]int{co[0] - ligney, co[1] - colx}
			}
			if !Verifformat(form) {
				return false
			}
			ligne = 0
			form.placer = false
			tab = append(tab, form)
			var newform Tetrominos
			form = newform
			id++
		} else if ligne == 4 && (compt == 0 || compt > 4) {
			return false
		}
	}
	return true
}
