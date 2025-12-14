package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed testInput.txt
var testInput string

func main() {
	points := load(testInput)
	mx := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			a := area(points[i], points[j])
			if a > mx {
				mx = a
			}
		}
	}
	fmt.Println(mx)
}

func area(p, q point) int {
	dx := q.x - p.x
	if p.x > q.x {
		dx = p.x - q.x
	}
	dx++
	dy := q.y - p.y
	if p.y > q.y {
		dy = p.y - q.y
	}
	dy++
	return dx * dy
}

func load(input string) []point {
	lines := strings.Split(input, "\n")
	var points []point
	for _, line := range lines {
		a := strings.Split(line, ",")
		x, _ := strconv.Atoi(a[0])
		y, _ := strconv.Atoi(a[1])
		points = append(points, point{x, y})
	}
	return points
}

type point struct {
	x, y int
}
