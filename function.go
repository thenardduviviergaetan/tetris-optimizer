package main

import (
	"os"
	"os/exec"
)

var (
	tab    []Tetrominos
	square [][]rune
)

func GrowSquare() {
	Addcol()
	Addligne()
}

func Addcol() {
	if len(square) == 0 {
		square = append(square, []rune{'.'})
		return
	}
	for index := range square {
		square[index] = append(square[index], '.')
	}
}

func Addligne() {
	col := len(square[0])
	var ligne []rune
	for index := 0; index < col; index++ {
		ligne = append(ligne, '.')
	}
	square = append(square, ligne)
}

func Remove(co [2]int, index int) {
	form := tab[index]
	for _, test := range form.co {
		square[co[0]+test[0]][co[1]+test[1]] = '.'
	}
	tab[index].placer = false
}

func Verifformat(form Tetrominos) bool {
	var test [4]int
	compt := 0
	for i, co := range form.co {
		for _, cotest := range form.co {
			if co[0]+1 == cotest[0] && co[1] == cotest[1] {
				compt++
			} else if co[0]-1 == cotest[0] && co[1] == cotest[1] {
				compt++
			} else if co[0] == cotest[0] && co[1]+1 == cotest[1] {
				compt++
			} else if co[0] == cotest[0] && co[1]-1 == cotest[1] {
				compt++
			}
		}
		if compt == 0 || compt >= 4 {
			return false
		} else {
			test[i] = compt
			compt = 0
		}
		trois := 0
		deux := 0
		un := 0
		for _, c := range test {
			switch c {
			case 1:
				un++
			case 2:
				deux++
			case 3:
				trois++
			}
		}
		if un > 2 && trois == 0 {
			return false
		}
	}
	return true
}

func Minsquare() {
	intcote := 2
	for intcote*intcote < len(tab)*4 {
		intcote++
	}
	for i := 1; i < intcote; i++ {
		GrowSquare()
	}
}

func Clear() {
	// time.Sleep(time.Millisecond * 100)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Place(co [2]int, index int) bool {
	form := tab[index]
	for _, test := range form.co {
		if test[0]+co[0] > len(square)-1 {
			return false
		} else if test[1]+co[1] > len(square[0])-1 {
			return false
		} else if test[0]+co[0] >= 0 && test[1]+co[1] >= 0 {
			if square[test[0]+co[0]][test[1]+co[1]] != '.' {
				return false
			}
		} else if test[0]+co[0] < 0 || test[1]+co[1] < 0 {
			return false
		}
	}
	for _, test := range form.co {
		square[co[0]+test[0]][co[1]+test[1]] = form.id
	}
	tab[index].placer = true
	return true
}
