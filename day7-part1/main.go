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
	start, levels := parseInput(input)
	_, hits, _ := compute(levels, []int{start}, 0, 0)
	fmt.Println(hits)
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

func compute(levels [][]bool, beams []int, hits, currentLevel int) (int, int, []int) {
	if currentLevel >= len(levels) {
		return 0, hits, beams
	}

	mp := map[int]bool{}
	for _, beam := range beams {
		if levels[currentLevel][beam] {
			mp[beam+1] = true
			mp[beam-1] = true
			hits++
			continue
		}
		mp[beam] = true
	}

	var nextBeams []int
	for k := range mp {
		nextBeams = append(nextBeams, k)
	}

	return compute(levels, nextBeams, hits, currentLevel+1)
}
