package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed testInput.txt
var testInput string

func main() {
	points := parseInput(testInput)
	compute(points, 10)
}

func compute(points []point, pairCount int) {
	var distances []distance
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d := distance{
				from:     points[i],
				to:       points[j],
				distance: computeDist(points[i], points[j]),
			}
			distances = append(distances, d)
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	var circuits []*circuit
	for i := 0; i < pairCount; i++ {
		d := distances[i]
		handled := false
		for _, c := range circuits {
			if c.has(d.from.index, d.to.index) {
				c.add(d.from.index, d.to.index)
				handled = true
				break
			}
		}
		if !handled {
			c := circuit{
				junctions: map[int]int{},
			}
			c.add(d.from.index, d.to.index)
			circuits = append(circuits, &c)
		}
	}

	l := len(circuits)
	_ = l
}

type point struct {
	index, x, y, z int
}

type distance struct {
	from, to point
	distance float64
}

func (d distance) String() string {
	return fmt.Sprintf("(%d, %d, %f)", d.from.index, d.to.index, d.distance)
}

type circuit struct {
	junctions map[int]int
}

func (c *circuit) has(i, j int) bool {
	if _, ok := c.junctions[i]; ok {
		return true
	}
	if _, ok := c.junctions[j]; ok {
		return true
	}
	return false
}

func (c *circuit) add(p, q int) {
	c.junctions[p] = p
	c.junctions[q] = q
}

func parseInput(s string) []point {
	lines := strings.Split(s, "\n")
	points := make([]point, 0, len(lines))
	for i, line := range lines {
		arr := strings.Split(line, ",")
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])
		z, _ := strconv.Atoi(arr[2])
		p := point{index: i, x: x, y: y, z: z}
		points = append(points, p)
	}
	return points
}

func computeDist(p, q point) float64 {
	xDiff := p.x - q.x
	yDiff := p.y - q.y
	zDiff := p.z - q.z

	return math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff))
}
