package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

//go:embed testInput.txt
var testInput string

//go:embed input.txt
var input string

func main() {
	machines := load(input)

	//m := machines[59]
	//////m := machines[9]
	//////m := machines[61]
	////fmt.Println(m.goalState)
	//s := m.solve()
	//fmt.Println(s)

	//fmt.Printf("we have %d machines\n", len(machines))
	//
	iterativeSolve(machines)

}

func iterativeSolve(machines []*machine) {
	//longRunning := map[int]bool{
	//	59: true,
	//	78: true,
	//}

	solution := 0
	for i, m := range machines {
		//if _, ok := longRunning[i]; ok {
		//	fmt.Printf("skipping index %d\n", i)
		//	continue
		//}
		s := m.solve()
		solution += s
		fmt.Printf("machine %d: %d - %d\n", i, s, solution)
	}
}

func channelSolve(machines []*machine) {
	solutionChan := make(chan []int)
	wg := sync.WaitGroup{}

	for i, mc := range machines {
		wg.Add(1)
		go func(sc chan<- []int, index int, m *machine) {
			sc <- []int{index, m.solve()}
		}(solutionChan, i, mc)
	}

	solution := 0
	remaining := len(machines)
	go func(ch <-chan []int) {
		for {
			s := <-ch
			remaining--
			solution += s[1]
			fmt.Printf("machine %d finished with answer: %d sum is: %d remaining: %d\n", s[0], s[1], solution, remaining)
			wg.Done()
		}

	}(solutionChan)

	wg.Wait()

	fmt.Println(solution)
}

func load(s string) []*machine {
	lines := strings.Split(s, "\n")
	m := make([]*machine, len(lines))
	for i, line := range lines {
		m[i] = newMachine(line)
	}
	return m
}

func newMachine(line string) *machine {
	elements := strings.Split(line, " ")
	m := &machine{}
	m.setStates(elements[0])

	var buttons []*button
	for i := 1; i < len(elements)-1; i++ {
		buttons = append(buttons, newButton(i, elements[i]))
	}
	m.buttons = buttons

	return m
}

type machine struct {
	goalState []bool
	buttons   []*button
}

func (m *machine) solve() int {
	var currentState []bool
	for range m.goalState {
		currentState = append(currentState, false)
	}

	var paths []*path
	for range m.buttons {
		p := &path{
			stack: []buttonClickState{
				buttonClickState{
					b:     nil,
					state: currentState,
				},
			},
		}
		paths = append(paths, p)
	}

	circularityCount := 0
	var nextPaths []*path
	//level := 0
	for {
		//level++
		nextPaths = nil
		for _, p := range paths {
			for _, b := range m.buttons {
				nextP, circularity := p.click(b)
				if circularity {
					//fmt.Println("circularity detected")
					circularityCount++
					continue
				}
				s, c := nextP.solved(m.goalState)
				if s {
					return c
				}
				nextPaths = append(nextPaths, nextP)
			}
		}
		paths = nextPaths
		//fmt.Printf("level: %d sol: %d paths: %d circularities: %d\n", level, len(paths[0].stack), len(paths), circularityCount)
	}
}

func equalStates(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (m *machine) setStates(s string) {
	g := s[1 : len(s)-1]
	var goalStates []bool
	for _, s := range g {
		goalStates = append(goalStates, s == '#')
	}
	m.goalState = goalStates
}

type button struct {
	name    int
	indexes []int
}

func (b *button) apply(input []bool) []bool {
	var result []bool
	for i := 0; i < len(input); i++ {
		if b.hasIndex(i) {
			result = append(result, !input[i])
			continue
		}
		result = append(result, input[i])
	}
	return result
}

func (b *button) hasIndex(in int) bool {
	for _, index := range b.indexes {
		if index == in {
			return true
		}
	}
	return false
}

func newButton(name int, s string) *button {
	g := s[1 : len(s)-1]
	a := strings.Split(g, ",")
	var indexes []int
	for _, i := range a {
		i, _ := strconv.Atoi(i)
		indexes = append(indexes, i)
	}
	return &button{indexes: indexes, name: name}
}

type buttonClickState struct {
	b     *button
	state []bool
}

func stateString(state []bool) string {
	s := "["
	for i := 0; i < len(state); i++ {
		if state[i] {
			s += "#"
			continue
		}
		s += "."
	}
	s += "]"
	return s
}

type path struct {
	stack []buttonClickState
}

func (p *path) click(b *button) (nextPath *path, circularity bool) {
	newState := b.apply(p.stack[len(p.stack)-1].state)

	for i := 0; i < len(p.stack); i++ {
		if equalStates(newState, p.stack[i].state) {
			return nil, true
		}
	}

	stack := append(p.stack, buttonClickState{b: b, state: newState})
	return &path{stack: stack}, false
}

func (p *path) solved(goalState []bool) (solved bool, level int) {
	if equalStates(goalState, p.stack[len(p.stack)-1].state) {
		return true, len(p.stack) - 1
	}

	return false, 0
}
