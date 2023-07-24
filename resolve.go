package main

func Result() {
	Minsquare()
	var backtracking func(int) bool
	backtracking = func(index int) bool {
		for line := range square {
			for col := range square[line] {
				co := [2]int{line, col}
				if !Place(co, index) {
					continue
				}
				if index+1 == len(tab) {
					return true
				}
				check := backtracking(index + 1)
				if check {
					return true
				}
				Remove(co, index)
			}
		}
		return false
	}
	find := false
	for !find {
		find = backtracking(0)
		if !find {
			GrowSquare()
		}
	}
}
