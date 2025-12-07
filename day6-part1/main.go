package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type problem struct {
	input     []int
	operation string
}

func (p *problem) compute() int {
	if p.operation == "+" {
		sum := 0
		for _, i := range p.input {
			sum += i
		}
		return sum
	}

	tot := 1
	for _, i := range p.input {
		tot *= i
	}
	return tot
}

func main() {
	problems := parseInput()
	total := 0
	for _, p := range problems {
		total += p.compute()
	}

	fmt.Println(total)
}

func parseInput() []*problem {
	lines := strings.Split(input, "\n")
	fields := strings.Fields(lines[0])
	problems := make([]*problem, len(fields))
	for i, field := range fields {
		d, _ := strconv.Atoi(field)
		p := problem{input: []int{d}}
		problems[i] = &p
	}

	for i := 1; i < len(lines)-1; i++ {
		fields = strings.Fields(lines[i])
		for j := 0; j < len(fields); j++ {
			d, _ := strconv.Atoi(fields[j])
			problems[j].input = append(problems[j].input, d)
		}
	}

	fields = strings.Fields(lines[len(lines)-1])
	for i := 0; i < len(fields); i++ {
		problems[i].operation = fields[i]
	}

	return problems
}
