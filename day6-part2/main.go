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

type problem struct {
	accumulators map[int]string
	op           string
}

func main() {
	total := process(input)
	fmt.Println(total)
}

func process(inp string) int {
	if inp[len(inp)-1:] == "\n" {
		inp = inp[:len(inp)-1]
	}
	ll := strings.Split(inp, "\n")
	lines := make([]string, 0, len(ll))
	for _, l := range ll {
		if len(l) == 3770 {
			lines = append(lines, l)
			continue
		}
		lines = append(lines, l+"  ")
	}

	for i, line := range lines {
		fmt.Printf("%d:%d\n", i, len(line))
	}

	ops := strings.Fields(lines[len(lines)-1])

	opCount := 0
	total := 0
	accums := map[int]string{}
	for i := 0; i < len(lines)-1; i++ {
		accums[i] = ""
	}
	cursor := 0
	done := false
	for {
		hasDigit := false
		for i := 0; i < len(lines)-1; i++ {
			if cursor >= len(lines[i]) {
				for k, v := range accums {
					accums[k] = v + "*"
				}
				done = true
				break
			}
			c := string(lines[i][cursor])
			if c != " " {
				hasDigit = true
			} else {
				c = "*"
			}
			accums[i] = accums[i] + c
		}

		if !hasDigit {
			subTotal := compute(accums, ops[opCount])
			//fmt.Printf("subTotal: %d\n", subTotal)
			total += subTotal
			accums = map[int]string{}
			opCount++
			cursor++
			//fmt.Println(cursor)
			if done {
				break
			}
			continue
		}

		cursor++
	}

	return total
}

func compute(m map[int]string, op string) int {
	var digits []int
	lim := len(m[0]) - 1
	for i := 0; i < lim; i++ {
		vertical := ""
		for j := 0; j < len(m); j++ {
			d := string(m[j][i])
			if d == "*" {
				d = ""
			}
			vertical += d
		}

		digit, _ := strconv.Atoi(vertical)
		digits = append(digits, digit)
	}

	if op == "+" {
		sum := 0
		for _, d := range digits {
			sum += d
		}
		return sum
	}

	product := 1
	for _, d := range digits {
		product *= d
	}

	return product
}
