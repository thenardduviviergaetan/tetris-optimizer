package main

type Tetrominos struct {
	id     rune
	forme  [4][4]rune
	co     [][2]int
	placer bool
}
