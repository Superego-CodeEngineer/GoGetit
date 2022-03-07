package main

import "fmt"

var result [8]int

func cal8queens(row int) {
	if row == 8 {
		printQuees(result)
		return
	}

	for col := 0; col < 8; col++ {
		if isOK(row, col) {
			result[row] = col
			cal8queens(row + 1)
		}
	}
}

func isOK(row, col int) bool {
	leftup, rightup := col-1, col+1
	for i := row - 1; i >= 0; i-- {
		if result[i] == col {
			return false
		}
		if leftup >= 0 && result[i] == leftup {
			return false
		}
		if rightup < 8 && result[i] == rightup {
			return false
		}
		leftup--
		rightup++
	}
	return true
}

func printQuees(result [8]int) {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if result[row] == col {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
