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
	g := newGraph(input)
	pathCount := g.traverse()
	fmt.Println("Path count:", pathCount)
}

func newGraph(s string) *graph {
	lines := strings.Split(s, "\n")
	m := map[string]*node{}
	for _, line := range lines {
		a := strings.Split(line, ":")
		n := node{label: a[0]}
		m[a[0]] = &n
	}
	m["out"] = &node{label: "out"}

	for _, line := range lines {
		a := strings.Split(line, ":")
		var children []*node
		for _, k := range strings.Split(strings.Trim(a[1], " "), " ") {
			children = append(children, m[k])
		}
		m[a[0]].children = children
	}

	return &graph{
		nodes: m,
	}
}

type graph struct {
	nodes map[string]*node
}

func (g *graph) traverse() (pathCount int) {
	pathCount = 0
	paths := []*path{
		&path{
			nodes: []*node{g.nodes["you"]},
		},
	}
	var donePaths []*path

	for {
		var nextPaths []*path
		for _, p := range paths {
			for _, child := range p.nodes[len(p.nodes)-1].children {
				next, done, circularity := p.expand(child)
				if circularity {
					continue
				}
				if done {
					donePaths = append(donePaths, next)
					continue
				}
				nextPaths = append(nextPaths, next)
			}
		}
		if len(nextPaths) == 0 {
			break
		}
		paths = nextPaths
	}

	return len(donePaths)
}

type node struct {
	label    string
	children []*node
}

type path struct {
	nodes []*node
}

func (p *path) expand(n *node) (next *path, done, circularity bool) {
	next = nil
	circularity = false
	done = n.label == "out"
	if done {
		return
	}

	for _, prev := range p.nodes {
		if prev.label == n.label {
			circularity = true
			return
		}
	}

	var nodes []*node
	nodes = append(nodes, p.nodes...)
	nodes = append(nodes, n)
	next = &path{nodes: nodes}
	return
}
