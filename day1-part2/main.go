package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type operationType int

const (
	operationTypeSet operationType = iota
	operationTypeRight
	operationTypeLeft
)

//go:embed input.txt
var input string

type operation struct {
	operationType operationType
	argument      int
}

func main() {
	inputStream := startInputStream(50)
	wrapsCountChan := startStater(inputStream)
	resultChan := startOutputter(wrapsCountChan)

	result := <-resultChan

	fmt.Println(result)

	//5856
	//5847
}

func startInputStream(startPoint int) <-chan operation {
	output := make(chan operation)

	lines := strings.Split(input, "\n")

	go func(ch chan<- operation) {
		defer close(ch)
		ch <- operation{operationType: operationTypeSet, argument: startPoint}

		for _, line := range lines {
			op, err := splitLine(line)
			if err != nil {
				panic(err)
			}
			ch <- op
		}

	}(output)

	return output
}

func startStater(input <-chan operation) (zeroChan <-chan int) {
	zeroChannel := make(chan int)
	go func(ch <-chan operation, zChan chan<- int) {
		state := 0
		var wraps int
		for {
			op, ok := <-ch
			if !ok {
				close(zeroChannel)
				return
			}
			switch op.operationType {
			case operationTypeSet:
				state = op.argument % 100
				continue
			case operationTypeRight:
				state, wraps = incrementWithWrapCount(state, op.argument, 100)
			case operationTypeLeft:
				state, wraps = decrementWithWrapCount(state, op.argument, 100)
			}
			zChan <- wraps
		}
	}(input, zeroChannel)

	return zeroChannel
}

func startOutputter(zeroStream <-chan int) (resultChan <-chan int) {
	rc := make(chan int)
	sum := 0
	go func(stream <-chan int) {
		defer func() {
			rc <- sum
		}()
		for {
			c, ok := <-stream
			if !ok {
				return
			}
			sum += c
		}
	}(zeroStream)

	return rc
}

func incrementWithWrapCount(state, increment, base int) (newState, wraps int) {
	newState = (state + increment) % base
	wraps = (state + increment) / base
	return newState, wraps
}

func decrementWithWrapCount(state, decrement, base int) (newState, wraps int) {
	newState = (state - decrement%base + base) % base
	if state == 0 {
		wraps = decrement / base
	} else if decrement >= state {
		wraps = 1 + (decrement-state)/base
	} else {
		wraps = 0
	}
	return newState, wraps
}

func splitLine(line string) (operation, error) {
	var ot operationType
	switch line[0] {
	case 'R', 'r':
		ot = operationTypeRight
	case 'L', 'l':
		ot = operationTypeLeft
	default:
		return operation{}, errors.New("invalid operation type")
	}

	arg, err := strconv.Atoi(line[1:])
	if err != nil {
		return operation{}, err
	}

	return operation{operationType: ot, argument: arg}, nil
}
