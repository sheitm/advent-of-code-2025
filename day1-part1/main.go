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

type operation struct {
	operationType operationType
	argument      int
}

//go:embed input.txt
var input string

func (o operation) String() string {
	op := ""
	switch o.operationType {
	case operationTypeSet:
		op = "set"
	case operationTypeRight:
		op = "right"
	case operationTypeLeft:
		op = "left"
	}

	return fmt.Sprintf("%s %d", op, o.argument)
}

func main() {
	inputStream := startInputStream(50)
	zeroCountChan := startStater(inputStream)
	resultChan := startOutputter(zeroCountChan)

	result := <-resultChan

	fmt.Println(result)
}

func startStater(input <-chan operation) (zeroChan <-chan struct{}) {
	zeroChannel := make(chan struct{})
	go func(ch <-chan operation, zChan chan<- struct{}) {
		state := 0
		for {
			op, ok := <-ch
			if !ok {
				close(zeroChannel)
				return
			}
			switch op.operationType {
			case operationTypeSet:
				state = op.argument % 100
			case operationTypeRight:
				state = (state + op.argument) % 100
			case operationTypeLeft:
				state = (state - op.argument%100 + 100) % 100
			}

			if state == 0 {
				zChan <- struct{}{}
			}
		}
	}(input, zeroChannel)

	return zeroChannel
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

func startOutputter(zeroStream <-chan struct{}) (resultChan <-chan int) {
	rc := make(chan int)
	count := 0
	go func(stream <-chan struct{}) {
		defer func() {
			rc <- count
		}()
		for {
			_, ok := <-stream
			if !ok {
				return
			}
			count++
		}
	}(zeroStream)

	return rc
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
