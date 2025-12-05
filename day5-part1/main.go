package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type idRange struct {
	from, to int
}

func main() {
	rangeChan, idChan := startInputStream()
	resultChan := startHandler(rangeChan, idChan)

	result := <-resultChan

	fmt.Println(result)
}

func startInputStream() (<-chan *idRange, <-chan int) {
	ranges := make(chan *idRange)
	ids := make(chan int)
	go func(r chan<- *idRange, i chan<- int) {
		defer close(i)
		idMode := false
		for _, line := range strings.Split(input, "\n") {
			if line == "" {
				idMode = true
				close(r)
				continue
			}
			if idMode {
				id, _ := strconv.Atoi(line)
				i <- id
				continue
			}
			arr := strings.Split(line, "-")
			from, _ := strconv.Atoi(arr[0])
			to, _ := strconv.Atoi(arr[1])
			r <- &idRange{from, to}
		}

	}(ranges, ids)

	return ranges, ids
}

func startHandler(rangeChan <-chan *idRange, idChan <-chan int) <-chan int {
	resultChan := make(chan int)
	go func(rc <-chan *idRange, ic <-chan int, res chan<- int) {
		freshCount := 0
		defer func() {
			res <- freshCount
		}()
		var ranges []*idRange
		for {
			r, ok := <-rc
			if !ok {
				break
			}
			ranges = append(ranges, r)
		}

		for {
			id, ok := <-idChan
			if !ok {
				break
			}
			for _, rng := range ranges {
				if id >= rng.from && id <= rng.to {
					freshCount++
					break
				}
			}
		}

	}(rangeChan, idChan, resultChan)

	return resultChan
}
