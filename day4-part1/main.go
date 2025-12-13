package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed testInput.txt
var testInput string

func main() {
	b := loadBoard(input)
	c := b.countFewerThan(4)
	fmt.Println(c)

	////c := b.countNeighbors(2, 0)
	////fmt.Println(c)
	//
	//counts := b.counts()
	//for y := 0; y < len(counts); y++ {
	//	for x := 0; x < len(counts[y]); x++ {
	//		fmt.Print(counts[y][x])
	//	}
	//	fmt.Println()
	//}
	//
	//// 1218 is too low!
	//fmt.Println(b.countAccessibleRolls(4))
}

func loadBoard(inp string) *board {
	var cells [][]bool
	for _, line := range strings.Split(inp, "\n") {
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

func (b *board) countFewerThan(lim int) int {
	res := 0
	for y := 0; y < len(b.cells); y++ {
		for x := 0; x < len(b.cells[y]); x++ {
			if b.cells[y][x] {
				c := b.countNeighbors(x, y)
				if c < lim {
					res++
				}
			}
		}
	}
	return res
}

func (b *board) counts() [][]int {
	var results [][]int
	for y := 0; y < len(b.cells); y++ {
		var row []int
		for x := 0; x < len(b.cells[y]); x++ {
			if !b.cells[y][x] {
				row = append(row, -1)
				continue
			}
			row = append(row, b.countNeighbors(x, y))
		}
		results = append(results, row)
	}
	return results
}

func (b *board) countNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy < 2; dy++ {
		yy := y + dy
		if yy < 0 || yy >= len(b.cells) {
			continue
		}
		for dx := -1; dx < 2; dx++ {
			xx := x + dx
			if xx < 0 || xx >= len(b.cells[yy]) {
				continue
			}
			if xx == x && yy == y {
				continue
			}
			//fmt.Printf("checking (%d %d) - %v\n", xx, yy, b.cells[yy][xx])
			if b.cells[yy][xx] {
				count++
			}
		}
	}
	return count
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
