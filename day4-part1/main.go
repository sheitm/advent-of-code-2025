package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	b := loadBoard()

	// 1218 is too low!
	fmt.Println(b.countAccessibleRolls(4))
}

func loadBoard() *board {
	var cells [][]bool
	for _, line := range strings.Split(input, "\n") {
		row := make([]bool, 0, len(line))
		for _, c := range line {
			row = append(row, c == '@')
		}
		cells = append(cells, row)
	}

	return &board{cells: cells}
}

type board struct {
	cells [][]bool
}

func (b *board) countAccessibleRolls(lim int) int {
	counterF := func(x, y int) int {
		count := 0
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				xx := x + dx
				if xx < 0 || xx >= len(b.cells[0]) {
					continue
				}
				yy := y + dy
				if yy < 0 || yy >= len(b.cells) {
					continue
				}
				if b.cells[xx][yy] {
					count++
				}
			}
		}
		return count
	}

	count := 0
	for y := 0; y < len(b.cells); y++ {
		for x := 0; x < len(b.cells[0]); x++ {
			c := counterF(x, y)
			if c < lim {
				count++
			}
		}
	}

	return count
}
