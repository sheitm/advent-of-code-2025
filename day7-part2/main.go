package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed testInput.txt
var testInput string

var (
	microInput1 = `.......S.......
...............
.......^.......`
)

func main() {
	start, levels := parseInput(testInput)
	result := compute(start, 0, 1, levels)
	fmt.Println(result)
}

func compute(beam, currentLevel, timelineCount int, levels [][]bool) int {
	if currentLevel >= len(levels) {
		return timelineCount
	}

	if levels[currentLevel][beam] {
		timelineCount++
		return compute(beam-1, currentLevel+1, timelineCount, levels) //+ compute(beam+1, currentLevel+1, timelineCount, levels)
	}
	return compute(beam, currentLevel+1, timelineCount, levels)
}

func parseInput(inp string) (int, [][]bool) {
	lines := strings.Split(inp, "\n")
	start := 0
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == 'S' {
			start = i
			break
		}
	}
	var levels [][]bool
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		var splitters []bool
		for _, c := range line {
			splitters = append(splitters, c == '^')
		}
		levels = append(levels, splitters)
	}

	return start, levels
}
